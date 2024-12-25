<p align="center">
  <img src="https://github.com/hulutech-web/tinker/blob/master/images/logo.png?raw=true" width="300" />
</p>
### 环境
项目的go.mod文件中，必须为 module goravel 否则会出现错误
```go
module goravel
```
### 简介
artisan 是一个适配 goravel 的go的运行环境，可以在该环境下执行go语言的语法，与浏览器的控制台运行js代码同理，如输入1+1回车会得到结果2，扩展实现了goravel的内置功能，如数据库db操作，orm操作，cache缓存，秘钥key生成,jwt秘钥等。

### 原理
artisan命令行使用了三方库[yaegi](https://github.com/traefik/yaegi)作为基础，该库是一个go语言的解释器，可以执行go语言代码，该扩展包在yaegi的基础上，扩展了goravel的内置功能，如数据库db操作，orm操作，cache缓存，秘钥key生成,jwt秘钥等。

### 安装
- 克隆该扩展包，然后放入packages/目录下
- 注册服务提供者，config.app中：
```go
import "goravel/packages/tinker"

"providers": []foundation.ServiceProvider{ 
	... 
	&tinker.ServiceProvider{}, 
}
```
- 系统命令：``go run . artisan tinker` 进入tinker环境，根据提示完成操作
- 命令行执行：``go generate go generate ./...`` 生成代码
- 命令行执行：``go mod tidy`` 清理依赖

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
#### orm方法，请参考gorm的用法
连续输入问题：使用方式为，进入orm下，命令行输入shebang回车，进入shebang多行命令行模式，编写完成后，输入命令:save保存回车，查看命令执行结果，默认引入了goravel/app/models模型。
- 函数调用参考gorm官方实现方式，可查看gorm官方文档
- 示例：orm命令行模式下，控制台参考命令如下，
  - ``bash go run . artisan tinker``
  - 选择orm菜单，进入orm tinker环境 // 2-上下箭头选择选项，回车进入
  - 输入命令：``shebang`` // 进入shebang多行命令行模式
  - 进入多行语法环境模式 // 3-进入go 语言环境
    ```go
users := []models.User{} //4-编写完成后回车
query.Model(&models.User{}).Where("name LIKE ?", "%市场%").Find(&users) //5-编写完成后回车
:save //6-输入:save 后回车保存查看执行结果

user:=models.User{}  //1-编写完成后回车
query.Model(&models.User{}).Where("name LIKE ?", "%小%").FindOne(&user) //2-编写完成后回车
:save //3-输入:save 后回车保存查看执行结果
    ``` 
#### 文件扫描，tinker 环境下，扩展会观察app/models下的模型文件，当有变化时，会自动执行makefile文件，这个过程需要一定时间，请耐心等待