# tinker
- install文件已经下载到本地的情况下，因为网络原因需要这样，如下是一个makefile文件内容
install-yaegi:
@bash install.sh -b $(GOPATH)/bin v0.9.0
- 运行go generate
go generate ./...

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
