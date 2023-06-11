package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/middleware"
)

func SetupRoutes(app *App) *gin.Engine {
	r := gin.Default()
	authMiddleware := middleware.AuthMiddleware(app.config)

	authRouter := r.Group("/api")
	authRouter.POST("/login", app.authController.Login)
	authRouter.GET("/me", authMiddleware, app.authController.GetCurrentUser)

	userRouter := r.Group("/api/users")
	userRouter.POST("", app.userController.Create)
	userRouter.GET("/:id", app.userController.GetById)

	postsRouter := r.Group("/api/posts")
	postsRouter.GET("", app.postController.GetAll)
	postsRouter.GET("/:id", app.postController.GetById)
	postsRouter.POST("", authMiddleware, app.postController.Create)
	postsRouter.PUT("/:id", authMiddleware, app.postController.Update)
	postsRouter.DELETE("/:id", authMiddleware, app.postController.Delete)

	return r
}
