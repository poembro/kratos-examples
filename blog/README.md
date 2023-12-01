# How to run this blog example server
1. You should ensure that your mysql server is running.
2. Ensure that the database named `testdb` has been created, 
   otherwise you should execute the following database script:
```mysql
create database testdb;
```
3. Modify the `configs/config.yaml` file and add your mysql information in the data source:
```yaml
data:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:3306)/testdb?parseTime=True
```


4. Run your blog server:
```bash 
 1141  make init     
 1141  make proto 
 1141  make proto-chw
 1142  ./bin/blog --conf=configs/config.yaml


  go generate 将扫描当前目录和子目录中的所有Go源文件，并查找以//go:generate开头的注释。找到注释后，它将执行指定的命令或脚本。 比如会执行   kratos proto client .  这个命令

1143  wire gen ./    
在执行完wire命令之后就会生成wire_gen.go，它的内容就是构建 依赖关系的结构对象。 需要连同wire_gen.go一起提交到版本控制系统里


https://blog.csdn.net/weixin_50071922/article/details/133278161

```



## 启动
```
root@gostart-dev:/workspace/kratos-examples/blog# ./bin/blog --conf=configs/config.yaml
INFO msg=config loaded: config.yaml format: yaml
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:228
[0.579ms] [rows:-] SELECT DATABASE()
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:231
[2.082ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1
INFO msg=sql/workspace/kratos-examples/blog/internal/data/data.go:90
[1.555ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test' AND table_name = 't_article' AND table_type = 'BASE TABLE'
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:228
[0.537ms] [rows:-] SELECT DATABASE()
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:231
[1.231ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:148
[0.799ms] [rows:-] SELECT * FROM `t_article` LIMIT 1
INFO msg=sql/workspace/pkg/mod/gorm.io/driver/mysql@v1.3.2/migrator.go:166
[4.007ms] [rows:-] SELECT column_name, column_default, is_nullable = 'YES', data_type, character_maximum_length, column_type, column_key, extra, column_comment, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'test' AND table_name = 't_article' ORDER BY ORDINAL_POSITION
INFO msg=[gRPC] server listening on: [::]:9000
INFO msg=[HTTP] server listening on: [::]:8000
```