package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Projects  []Project `gorm:"foreignKey:UserID" json:"projects"`
}

type Personel struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	ProjectID     uint           `gorm:"index" json:"project_id"`
	NameSurname   string         `gorm:"not null" json:"name_surname"`
	TcNo          string         `gorm:"unique;not null" json:"tc_no"`
	PhoneNo       string         `gorm:"not null" json:"phone_no"`
	Profession    string         `json:"profession"`
	BloodType     string         `json:"blood_type"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	Description   string         `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Isg_Egitim    []Isg_Egitim   `gorm:"foreignKey:PersonelID" json:"isg_egitim"`
	Saglik_Raporu []SaglikRaporu `gorm:"foreignKey:PersonelID" json:"saglik_raporu"`
}

type SaglikRaporu struct {
	ID              uint      `json:"id"`
	PersonelID      uint      `json:"personel_id"`
	BaslangicTarihi time.Time `json:"baslangic_tarihi"`
	BitisTarihi     time.Time `json:"bitis_tarihi"`
}

type Isg_Egitim struct {
	ID              uint      `json:"id"`
	PersonelID      uint      `json:"personel_id"`
	BaslangicTarihi time.Time `json:"baslangic_tarihi"`
	BitisTarihi     time.Time `json:"bitis_tarihi"`
}
type Project struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserID      uint       `json:"user_id"`
	Description string     `json:"description"`
	Personel    []Personel `gorm:"foreignKey:ProjectID" json:"personel"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
