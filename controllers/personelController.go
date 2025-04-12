package controllers

import (
	"isg_API/db"
	"isg_API/models"
	"net/http"


	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonelController struct {
	db *gorm.DB
}

func NewPersonelController() *PersonelController {
	return &PersonelController{
		db: db.GetDB(),
	}
}

type PersonelRequest struct {
	ProjectID   uint   `json:"project_id" binding:"required"`
	NameSurname string `json:"name_surname" binding:"required"`
	TcNo        string `json:"tc_no" binding:"required"`
	PhoneNo     string `json:"phone_no" binding:"required"`
	Profession  string `json:"profession"`
	BloodType   string `json:"blood_type"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
}

func (pc *PersonelController) CreatePersonel(c *gin.Context) {
	var req PersonelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var project models.Project
	if err := pc.db.First(&project, req.ProjectID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found"})
		return
	}

	personel := models.Personel{
		ProjectID:   req.ProjectID,
		NameSurname: req.NameSurname,
		TcNo:        req.TcNo,
		PhoneNo:     req.PhoneNo,
		Profession:  req.Profession,
		BloodType:   req.BloodType,
		IsActive:    true,
		Description: req.Description,
	}

	if err := pc.db.Create(&personel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create personnel"})
		return
	}

	c.JSON(http.StatusCreated, personel)
}

func (pc *PersonelController) UpdatePersonel(c *gin.Context) {
	id := c.Param("id")
	var req PersonelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var project models.Project
	if err := pc.db.First(&project, req.ProjectID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found"})
		return
	}

	var personel models.Personel
	if err := pc.db.First(&personel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Personnel not found"})
		return
	}

	personel.ProjectID = req.ProjectID
	personel.NameSurname = req.NameSurname
	personel.TcNo = req.TcNo
	personel.PhoneNo = req.PhoneNo
	personel.Profession = req.Profession
	personel.BloodType = req.BloodType
	personel.Description = req.Description

	if err := pc.db.Save(&personel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update personnel"})
		return
	}

	c.JSON(http.StatusOK, personel)
}

func (pc *PersonelController) DeletePersonel(c *gin.Context) {
	id := c.Param("id")

	var personel models.Personel
	if err := pc.db.First(&personel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Personnel not found"})
		return
	}

	if err := pc.db.Delete(&personel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete personnel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Personnel deleted successfully"})
}