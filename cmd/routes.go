package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/middleware"
)

func SetupRoutes(app *App) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	authMiddleware := middleware.AuthMiddleware(app.config)

	authRouter := router.Group("/api")
	authRouter.POST("/login", app.authController.Login)
	authRouter.GET("/me", authMiddleware, app.authController.GetCurrentUser)

	userRouter := router.Group("/api/users")
	userRouter.POST("", app.userController.Create)
	userRouter.GET("/:id", app.userController.GetById)

	postsRouter := router.Group("/api/posts")
	postsRouter.GET("", app.postController.GetPosts)
	postsRouter.GET("/:id", app.postController.GetById)
	postsRouter.POST("", authMiddleware, app.postController.Create)
	postsRouter.PUT("/:id", authMiddleware, app.postController.Update)
	postsRouter.DELETE("/:id", authMiddleware, app.postController.Delete)

	return router
}
