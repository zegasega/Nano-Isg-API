package routes

import (
	"isg_API/controllers"
	"isg_API/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Logger middleware'ini ekle
	r.Use(middleware.LoggerMiddleware())

	userController := controllers.NewUserController()
	projectController := controllers.NewProjectController()
	personelController := controllers.NewPersonelController()
	isgController := controllers.NewIsgController()
	saglikRaporuController := controllers.NewSaglikRaporuController()

	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user/:id", userController.GetProfile)

		// Project routes
		protected.POST("/project", projectController.CreateProject)
		protected.GET("/projects", projectController.GetProjects)
		protected.GET("/project/:id", projectController.GetProjectByID)
		protected.PUT("/project/:id", projectController.UpdateProject)
		protected.DELETE("/project/:id", projectController.DeleteProject)
		protected.GET("/project/user/:user_id", projectController.GetProjectByUserID)

		// Personnel routes
		protected.POST("/personnel", personelController.CreatePersonel)
		protected.PUT("/personnel/:id", personelController.UpdatePersonel)
		protected.DELETE("/personnel/:id", personelController.DeletePersonel)

		// ISG Egitim routes
		protected.POST("/isg", isgController.CreateIsg)
		protected.GET("/isg", isgController.GetAllIsg)
		protected.PUT("/isg/:id", isgController.UpdateIsg)
		protected.DELETE("/isg/:id", isgController.DeleteIsg)

		// Saglik Raporu routes
		protected.POST("/saglik-raporu", saglikRaporuController.CreateSaglikRaporu)
		protected.PUT("/saglik-raporu/:id", saglikRaporuController.UpdateSaglikRaporu)
		protected.DELETE("/saglik-raporu/:id", saglikRaporuController.DeleteSaglikRaporu)
		protected.GET("/saglik-raporu/personel/:personel_id", saglikRaporuController.GetSaglikRaporuByPersonelID)
	}

	return r
}
