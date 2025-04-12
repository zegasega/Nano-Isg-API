package controllers

import (
	"isg_API/config"
	"isg_API/db"
	"isg_API/models"
	"isg_API/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{
		DB: db.GetDB(),
	}
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"oneof=user admin"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (uc *UserController) Login(c *gin.Context) {
	var payload LoginInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Login Payload"})
		return
	}

	var user models.User
	if err := uc.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password or email"})
		return
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password or email"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour * 365)
	claims := &models.Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.LoadConfig().JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (uc *UserController) Register(c *gin.Context) {

	var payload RegisterInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existUser models.User
	if err := uc.DB.Where("email", payload.Email).First(&existUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email Already Exist!"})
		return
	}

	if payload.Role == "" {
		payload.Role = "user"
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Occued Whiles hashing Password"})
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     payload.Role,
	}

	if err := uc.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Registerde Succesfuly",
		"user":    newUser})
}

func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}


func (uc *UserController) GetProjectByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	var projects []models.Project
	if err := uc.DB.Preload("Personel").Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	if len(projects) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No projects found for this user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"projects": projects,
	})
}
