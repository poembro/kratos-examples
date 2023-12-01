package data

import (
	"fmt"
	"time"

	"github.com/go-kratos/examples/blog/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	// init mysql driver
	"github.com/go-kratos/examples/blog/internal/data/model"
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

type OrmLog struct {
	log *log.Helper
}

func (l *OrmLog) Printf(format string, v ...interface{}) {
	/*
		body := fmt.Sprintf(format, v...)
		sql2 := strings.Replace(body, "\n", "", -1)
		sql := strings.Replace(sql2, "\t", "", -1)
	*/
	l.log.Info("sql", fmt.Sprintf(format, v...))
}

func NewOrmLog(log *log.Helper) *OrmLog {
	return &OrmLog{
		log: log,
	}
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	newLogger := gormLogger.New(
		NewOrmLog(log), // 自定义输出至zap
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		gormLogger.Config{
			SlowThreshold:             500 * time.Millisecond, // 1 * time.Second, // 慢 SQL 阈值
			LogLevel:                  gormLogger.Info,        // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                  // 彩色打印
		},
	)

	client, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		SkipDefaultTransaction: true,  // 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行
		PrepareStmt:            false, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //严格匹配表名 默认是复数形式
		},
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   newLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Error("失败的打开mysql连接", err)
		panic(err)
	}

	for _, t := range model.Migrations {
		defTableOpts := "ENGINE=InnoDB"

		if t, ok := t.(model.TableOptions); ok {
			defTableOpts = t.TableOptions()
		}

		if err := client.Set("gorm:table_options", defTableOpts).AutoMigrate(t); err != nil {
			log.Error("运行迁移失败", err)
			panic(err)
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		sqlDB, _ := d.db.DB()
		if err := sqlDB.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
