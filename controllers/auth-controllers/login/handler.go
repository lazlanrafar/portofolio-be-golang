package loginAuth

import "github.com/gin-gonic/gin"

type handler struct {
	service Service
}

func NewHandlerLogin(service Service) *handler {
	return &handler{service}
}

func (h *handler) LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := h.service.LoginRepository(username, password)
	if err != "success" {
		c.JSON(500, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}