package route

import "github.com/gin-gonic/gin"
import . "github.com/cy422396350/cygin/api"

func InitRoute() *gin.Engine {
	r := gin.Default()
	//可恶的小图标pass
	r.StaticFile("/favicon.ico", "./favicon.ico")
	// 从redis集群里面取账号密码
	name, password := GetAuthUserInfo()
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		name: password,
	}))
	authorized.POST("/index", IndexApi)
	return r
}
