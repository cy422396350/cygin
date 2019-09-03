package route

import "github.com/gin-gonic/gin"
import . "github.com/cy422396350/cygin/api"

func InitRoute() *gin.Engine {
	r := gin.Default()
	//可恶的小图标pass
	r.StaticFile("/favicon.ico", "./favicon.ico")
	// IndexApi为一个handler
	r.GET("/", IndexApi)
	return r
}
