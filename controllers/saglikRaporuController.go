package controllers

import (
	"isg_API/db"
	"isg_API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SaglikRaporuController struct {
	db *gorm.DB
}

func NewSaglikRaporuController() *SaglikRaporuController {
	return &SaglikRaporuController{
		db: db.GetDB(),
	}
}

type SaglikRaporuInput struct {
	PersonelID      uint      `json:"personel_id" binding:"required"`
	BaslangicTarihi time.Time `json:"baslangic_tarihi" binding:"required"`
	BitisTarihi     time.Time `json:"bitis_tarihi" binding:"required"`
}

type SaglikRaporuUpdateInput struct {
	BaslangicTarihi time.Time `json:"baslangic_tarihi"`
	BitisTarihi     time.Time `json:"bitis_tarihi"`
}

func (src *SaglikRaporuController) CreateSaglikRaporu(c *gin.Context) {
	var payload SaglikRaporuInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}

	var personel models.Personel
	if err := src.db.First(&personel, payload.PersonelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Personel not found"})
		return
	}

	saglikRaporu := models.SaglikRaporu{
		PersonelID:      payload.PersonelID,
		BaslangicTarihi: payload.BaslangicTarihi,
		BitisTarihi:     payload.BitisTarihi,
	}

	if err := src.db.Create(&saglikRaporu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Saglik Raporu"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Saglik Raporu successfully created", "saglik_raporu": saglikRaporu})
}

func (src *SaglikRaporuController) UpdateSaglikRaporu(c *gin.Context) {
	var saglikRaporu models.SaglikRaporu
	id := c.Param("id")

	if err := src.db.First(&saglikRaporu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Saglik Raporu not found"})
		return
	}

	var payload SaglikRaporuUpdateInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	if !payload.BaslangicTarihi.IsZero() {
		saglikRaporu.BaslangicTarihi = payload.BaslangicTarihi
	}
	if !payload.BitisTarihi.IsZero() {
		saglikRaporu.BitisTarihi = payload.BitisTarihi
	}

	if err := src.db.Save(&saglikRaporu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Saglik Raporu"})
		return
	}

	c.JSON(http.StatusOK, saglikRaporu)
}

func (src *SaglikRaporuController) DeleteSaglikRaporu(c *gin.Context) {
	id := c.Param("id")

	var saglikRaporu models.SaglikRaporu
	if err := src.db.First(&saglikRaporu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Saglik Raporu not found"})
		return
	}

	if err := src.db.Delete(&saglikRaporu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Saglik Raporu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Saglik Raporu deleted successfully"})
}

func (src *SaglikRaporuController) GetSaglikRaporuByPersonelID(c *gin.Context) {
	personelID := c.Param("personel_id")

	var saglikRaporlari []models.SaglikRaporu
	if err := src.db.Where("personel_id = ?", personelID).Find(&saglikRaporlari).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Saglik Raporlari"})
		return
	}

	c.JSON(http.StatusOK, saglikRaporlari)
} 