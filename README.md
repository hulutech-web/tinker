# 环境
## 项目的go.mod文件中，必须为 module goravel 否则会出现错误
``go.mod`` 文件内容如下：
```go
module goravel
````


# tinker 安装
- 克隆该扩展包，然后放入packages/目录下
- 注册：
```go
import 	"goravel/packages/tinker"
"providers": []foundation.ServiceProvider{
	    ...
        &tinker.ServiceProvider{},
}
```
- 系统命令：``go run . artisan:tinker` 进入tinker环境，根据提示完成操作
- install文件已经下载到本地的情况下，因为网络原因需要这样，如下是一个makefile文件内容
install-yaegi:
@bash install.sh -b $(GOPATH)/bin v0.9.0
- yaegi预加载
    进入到tinker目录下，命令行执行 ``make`` ，将自动安装所需依赖
- 运行go generate
go generate ./...
- ``go mod tidy``

#### 数据库方法
```go
db.Find("depts")
db.WhereMap("users",map[string]interface{}{"name":"市场部-经理"})
db.WhereNot("users",map[string]interface{}{"name":"市场部-经理"})
db.WhereNotEqual("users","name","市场部-经理")
db.WhereIn("users","name", []string{"市场部-经理", "市场部-主管"})
db.WhereLike("users","name", "市场部")
db.WhereBetween("users","created_at", `2024-08-22 00:00:00`, `2024-08-23 23:59:59`)
```
#### orm方法，请参考gorm的用法，在shebang下，需要引入模型，再进行操作,其中方式来源于gorm官方，可查看gorm官方文档
- 示例
```go
users:=[]models.User{}
query.Model(&models.User{}).Where("name LIKE ?", "%市场%").Find(&users)

user:=models.User{}
query.Query().Model(&models.User{}).Where("name LIKE ?", "%小%").FindOne(&user)
```
#### 文件扫描，artisan tinker下，扩展会观察app/models下的模型文件，当有变化时，会自动执行makefile文件，这个过程需要一定时间，请耐心等待

#### 示例：可查看example.md文件

#### 预览：
<p align="center">
  <img src="https://github.com/hulutech-web/tinker/blob/master/images/dashboard.png?raw=true" width="750" />
</p>

