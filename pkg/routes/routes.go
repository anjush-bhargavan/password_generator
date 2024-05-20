package routes

import ("github.com/gin-gonic/gin"
"github.com/anjush-bhargavan/password_generator/pkg/handler"
)

func InitRoutes(handler *handler.PasswordHandler, r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/generate-password", handler.GeneratePassword)

}
