package api

import "github.com/gin-gonic/gin"

func Status(context *gin.Context) {
	context.JSON(200, &struct {
		OK string
	}{
		OK: "OK",
	})
}

// CreateRouter creates a gin router group that can then have a relative path
// easily applied to all routes.
// Routes registerd
//   - /status
func CreateRouter() gin.RouterGroup {
	group := gin.RouterGroup{}

	group.GET("/status", Status)

	return group
}
