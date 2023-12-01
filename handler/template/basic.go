package template

import (
	"fmt"
	"microservice/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Basic struct {
	BasicRepo *repository.BasicRepo
}

func (o *Basic) Create(c *gin.Context) {
	fmt.Println("Create method")
}
func (o *Basic) List(c *gin.Context) {
	fmt.Println("Basic")
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
