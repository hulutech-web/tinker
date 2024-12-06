```go
import "goravel/app/models"

users:=[]models.User{}
query.Query().Model(&models.User{}).Where("name LIKE ?", "%小%").FindAll(&users)

user:=models.User{}
query.Query().Model(&models.User{}).Where("name LIKE ?", "%小%").FindOne(&user)
```


