### orm方法
```go
users:=[]models.User{}
query.Query().Model(&models.User{}).Where("name LIKE ?", "%小%").FindAll(&users)

user:=models.User{}
query.Query().Model(&models.User{}).Where("name LIKE ?", "%小%").FindOne(&user)
```

### table方法
```go
db.Find("depts")
db.WhereMap("users",map[string]interface{}{"name":"市场部-经理"})
```

### cache方法
```go

```