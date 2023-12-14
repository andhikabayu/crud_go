package controllers

import (
	"crud_go/helpers"
	"crud_go/middleware"
	"crud_go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	var registerRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		helpers.ErrorJSON(c, err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	// check if the user already exist or not
	db.Where("username = ?", registerRequest.Username).First(&user)
	if user.ID != 0 {
		helpers.ErrorJSON(c, "Username already taken.")
		// c.JSON(http.StatusConflict, gin.H{"error": "Username already taken."})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorJSON(c, "Failed to hash password.")
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password."})
		return
	}
	user.Username = registerRequest.Username
	user.Password = string(hashedPassword)

	// save user to database
	db.Create(&user)
	helpers.ErrorJSON(c, "User registered successfully.")
	// c.JSON(http.StatusOK, gin.H{"message": "User registered successfully."})
}

func LoginHandler(c *gin.Context) {
	var user models.User

	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		helpers.ErrorJSON(c, err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retrieve the user from the database
	db := c.MustGet("db").(*gorm.DB)
	db.Where("username = ?", loginRequest.Username).First(&user)
	if user.ID == 0 {
		helpers.ErrorJSON(c, "Invalid Credential.")
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials."})
		return
	}

	// compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		helpers.ErrorJSON(c, "Invalid credentials.")
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// generate token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		helpers.ErrorJSON(c, "Error creating token")
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	helpers.SuccessJSON(c, "login succesfull", token)
	// c.JSON(http.StatusOK, gin.H{"token": token})
}
