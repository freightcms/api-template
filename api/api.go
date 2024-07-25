package api

import "github.com/gin-gonic/gin"

// Status returns an OK string for the root of the application to call out that the service is up and functing
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
func Register(group *gin.RouterGroup) {
	group.GET("/status", Status)
}
