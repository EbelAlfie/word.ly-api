package wordly

import (
	"github.com/gin-gonic/gin"

	"wordly/user"
)

func main() {
	server := gin.Default()

	server.Group("/wordly")
	{
		server.Group("/user", user.UserRoute)
	}

	server.Run("6969")
}
