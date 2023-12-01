package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/examples/blog/internal/biz"
	"github.com/go-kratos/examples/blog/internal/data/model"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(db *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: db,
		log:  log.NewHelper(logger),
	}
}

func (d *articleRepo) ListArticle(ctx context.Context, typ string, page, limit int) (items []model.Article, count int64, err error) {
	conn := d.data.db.Model(&model.Article{})
	if typ != "" {
		conn.Where("type = ? ", typ)
	}

	conn.Count(&count)
	if err = conn.Order("sort_index ASC").Offset((page - 1) * limit).Limit(limit).Find(&items).Error; err != nil && err != gorm.ErrRecordNotFound {

		return items, count, err
	}
	return items, count, nil
}

func (d *articleRepo) GetArticle(ctx context.Context, where ...interface{}) (*model.Article, error) {
	var item = model.Article{}
	err := d.data.db.First(&item, where...).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &item, nil
}

func (d *articleRepo) CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	item := model.Article{}

	err := d.data.db.Model(&item).Omit("id").Create(article).Error
	if err != nil {
		return &item, err
	}
	return &item, nil
}

func (d *articleRepo) UpdateArticle(ctx context.Context, id int64, article *model.Article) error {
	item := model.Article{}
	err := d.data.db.Model(&item).Omit("id").Where("id = ?", id).Updates(article).Error
	return err
}

func (d *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return d.data.db.Where("id = ?", id).Delete(&model.Article{}).Error
}

func likeKey(id int64) string {
	return fmt.Sprintf("like:%d", id)
}

func (d *articleRepo) GetArticleLike(ctx context.Context, id int64) (rv int64, err error) {
	get := d.data.rdb.Get(ctx, likeKey(id))
	rv, err = get.Int64()

	return
}

func (d *articleRepo) IncArticleLike(ctx context.Context, id int64) error {
	_, err := d.data.rdb.Incr(ctx, likeKey(id)).Result()
	return err
}
