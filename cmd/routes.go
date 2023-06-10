package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/middleware"
)

func SetupRoutes(app *App) *gin.Engine {
	r := gin.Default()
	authMiddleware := middleware.AuthMiddleware(app.config)

	authRouter := r.Group("/api")
	authRouter.POST("/login", app.authController.SignInHandler)

	postsRouter := r.Group("/api/posts")
	postsRouter.GET("", app.postController.GetAll)
	postsRouter.GET("/:id", app.postController.GetById)
	postsRouter.POST("", authMiddleware, app.postController.Create)
	postsRouter.PUT("/:id", authMiddleware, app.postController.Update)
	postsRouter.DELETE("/:id", authMiddleware, app.postController.Delete)

	return r
}
