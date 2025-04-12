package controllers

import (
	"isg_API/db"
	"isg_API/logger"
	"isg_API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IsgController struct {
	db *gorm.DB
}

func NewIsgController() *IsgController {
	return &IsgController{
		db: db.GetDB(),
	}
}

type ISGEgitimInput struct {
	PersonelID      uint      `json:"personel_id" binding:"required"`
	BaslangicTarihi time.Time `json:"baslangic_tarihi" binding:"required"`
	BitisTarihi     time.Time `json:"bitis_tarihi" binding:"required"`
}

type ISGEgitimUpdateInput struct {
	BaslangicTarihi time.Time `json:"baslangic_tarihi"`
	BitisTarihi     time.Time `json:"bitis_tarihi"`
}

func (ic *IsgController) CreateIsg(c *gin.Context) {
	var payload ISGEgitimInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.LogErrorWithContext("CreateIsg", "Invalid payload", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}

	var personel models.Personel
	if err := ic.db.First(&personel, payload.PersonelID).Error; err != nil {
		logger.LogErrorWithContext("CreateIsg", "Personel not found", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Personel not found"})
		return
	}

	isg := models.Isg_Egitim{
		PersonelID:      payload.PersonelID,
		BaslangicTarihi: payload.BaslangicTarihi,
		BitisTarihi:     payload.BitisTarihi,
	}

	if err := ic.db.Create(&isg).Error; err != nil {
		logger.LogErrorWithContext("CreateIsg", "Failed to create ISG Egitim", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ISG Egitim"})
		return
	}

	logger.LogInfo("ISG Egitim created successfully for personel ID: %d", payload.PersonelID)
	c.JSON(http.StatusCreated, gin.H{"message": "Isg Egitim Succesfully created", "isg_egitim": isg})
}

func (ic *IsgController) UpdateIsg(c *gin.Context) {
	var isg models.Isg_Egitim
	id := c.Param("id")

	if err := ic.db.First(&isg, id).Error; err != nil {
		logger.LogErrorWithContext("UpdateIsg", "ISG Egitim not found", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "ISG Egitim not found"})
		return
	}

	var payload ISGEgitimUpdateInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.LogErrorWithContext("UpdateIsg", "Invalid payload", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	if !payload.BaslangicTarihi.IsZero() {
		isg.BaslangicTarihi = payload.BaslangicTarihi
	}
	if !payload.BitisTarihi.IsZero() {
		isg.BitisTarihi = payload.BitisTarihi
	}

	if err := ic.db.Save(&isg).Error; err != nil {
		logger.LogErrorWithContext("UpdateIsg", "Failed to update ISG Egitim", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ISG Egitim"})
		return
	}

	logger.LogInfo("ISG Egitim updated successfully for ID: %s", id)
	c.JSON(http.StatusOK, isg)
}

func (ic *IsgController) DeleteIsg(c *gin.Context) {
	id := c.Param("id")

	var isg models.Isg_Egitim
	if err := ic.db.First(&isg, id).Error; err != nil {
		logger.LogErrorWithContext("DeleteIsg", "ISG Egitim not found", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "ISG Egitim not found"})
		return
	}

	if err := ic.db.Delete(&isg).Error; err != nil {
		logger.LogErrorWithContext("DeleteIsg", "Failed to delete ISG Egitim", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ISG Egitim"})
		return
	}

	logger.LogInfo("ISG Egitim deleted successfully for ID: %s", id)
	c.JSON(http.StatusOK, gin.H{"message": "ISG Egitim deleted successfully"})
}

func (ic *IsgController) GetAllIsg(c *gin.Context) {
	var isgEgitimleri []models.Isg_Egitim
	if err := ic.db.Find(&isgEgitimleri).Error; err != nil {
		logger.LogErrorWithContext("GetAllIsg", "Failed to fetch ISG Egitimleri", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ISG Egitimleri"})
		return
	}

	logger.LogInfo("Retrieved %d ISG Egitimleri", len(isgEgitimleri))
	c.JSON(http.StatusOK, isgEgitimleri)
}
