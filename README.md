## 介绍
这是一个用golang编写的自动排班的值班系统

### 系统环境

- Mac os
- golang 
- gin
- gorm

### 列表
1. 自动排班 --- 定时按照每周进行排班，自动跳过节假日
2. 用户的添加删除 --- 针对用户的一些信息简单接口
3. 值班信息的查询  --- 查询值班信息，变更值班人员

### 启动方式
- 修改yaml下的配置文件mysql地址，目前没有redis
- 执行 go run migrate.go -m=true 进行数据库初始化
- 执行 go run main.go 即可

### 联系方式
mail: xiaohui920@sina.cn
