package controllers

import (
	"isg_API/db"
	"isg_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectController struct {
	db *gorm.DB
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		db: db.GetDB(),
	}
}

type ProjectRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Description string `json:"description"`
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		UserID:      req.UserID,
		Description: req.Description,
	}

	if err := pc.db.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}


	var createdProject models.Project
	if err := pc.db.Preload("Personel").First(&createdProject, project.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created project"})
		return
	}

	c.JSON(http.StatusCreated, createdProject)
}

func (pc *ProjectController) GetProjects(c *gin.Context) {
	var projects []models.Project
	if err := pc.db.Preload("Personel.Isg_Egitim").Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (pc *ProjectController) GetProjectByID(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := pc.db.Preload("Personel.Isg_Egitim").First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var project models.Project
	if err := pc.db.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	project.UserID = req.UserID
	project.Description = req.Description

	if err := pc.db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	
	var updatedProject models.Project
	if err := pc.db.Preload("Personel").First(&updatedProject, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated project"})
		return
	}

	c.JSON(http.StatusOK, updatedProject)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := pc.db.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	if err := pc.db.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

func (pc *ProjectController) GetProjectByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	var projects []models.Project

	if err := pc.db.Preload("Personel").Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
} 