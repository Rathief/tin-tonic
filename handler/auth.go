package handler

import (
	"net/http"
	"tin-tonic/entity"
	"tin-tonic/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (ah AuthHandler) Register(c *gin.Context) {
	// get input
	var acc entity.Store
	if err := c.ShouldBindJSON(&acc); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "Cannot read input properly.",
		// 	"details": err,
		// })
		utils.ErrorMessage(c, utils.ErrReadData)
		return
	}
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password.",
			"details": err,
		})
		return
	}
	acc.Password = string(hashed)
	// exec sql
	result := ah.DB.Select("StoreEmail", "password").Create(&acc)
	if result.Error != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"message": "Something went wrong upon executing query.",
		// 	"details": result.Error,
		// })
		utils.ErrorMessage(c, utils.ErrInsertData)
	}
}
func (ah AuthHandler) Login(c *gin.Context) {
	var acc entity.Store
	var storedAcc entity.Store
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot read input properly.",
			"details": err,
		})
		return
	}
	if err := ah.DB.Where("StoreEmail = ?", acc.StoreEmail).First(&storedAcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot read stored data properly.",
			"details": err,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedAcc.Password), []byte(acc.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong Credentials.",
			"details": err,
		})
		return
	}
}
