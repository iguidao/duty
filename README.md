## 介绍
这是一个用golang编写的自动排班的值班系统

### 开发环境
- golang v1.12.4
- github.com/fsnotify/fsnotify v1.4.9
- github.com/gin-gonic/gin v1.5.0
- github.com/jinzhu/gorm v1.9.12
- github.com/robfig/cron v1.2.0
- github.com/spf13/viper v1.6.2

### 功能列表
1. 自动排班 --- 定时按照每周进行排班，自动跳过节假日
2. 用户的添加删除 --- 针对用户的一些信息简单接口
3. 值班信息的查询  --- 查询值班信息，变更值班人员

### 启动方式
- 进入docker/mysql目录下，执行mysql.txt文件中命令，创建本地数据库，注意密码没有填写
- 修改yaml下的配置文件mysql地址，目前没有redis
- 执行 go run migrate.go -m=true 进行数据库初始化，会创建默认数据库
- 执行 go run main.go 运行代码

### API列表

--- | --- | --- 
GET |   /api/v1/health | 健康检查
GET |   /api/v1/cookie | 获取cookie
GET |   /duty/v1/info | 获取时间内的值班信息
PUT |   /duty/v1/shift/:id | 更改值班人员
GET |   /staff/v1/user | 获取所有人员信息
POST |  /staff/v1/user | 新增人员信息
DELETE | /staff/v1/user/:id | 删除人员信息

### 联系方式
mail: xiaohui920@sina.cn
