package controllers

import (
	"net/http"
	"time"

	"api-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ageRatingCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetAllAgeRatingCategory godoc
// @Summary Get all AgeRatingCategory.
// @Description Get a list of AgeRatingCategory.
// @Tags AgeRatingCategory
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.AgeRatingCategory
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateAgeRatingCategory godoc
// @Summary Create New AgeRatingCategory.
// @Description Creating a new AgeRatingCategory.
// @Tags AgeRatingCategory
// @Param Body body ageRatingCategoryInput true "the body to create a new AgeRatingCategory"
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func CreateRating(c *gin.Context) {
	// Validate input
	var input ageRatingCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	rating := models.AgeRatingCategory{Name: input.Name, Description: input.Description}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetAgeRatingCategoryById godoc
// @Summary Get AgeRatingCategory.
// @Description Get an AgeRatingCategory by id.
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "AgeRatingCategory id"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [get]
func GetRatingById(c *gin.Context) { // Get model if exist
	var rating models.AgeRatingCategory

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetMovieByAgeRatingCategoryId godoc
// @Summary Get Movies.
// @Description Get all Movies by AgeRatingCategoryId.
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "AgeRatingCategory id"
// @Success 200 {object} []models.Movie
// @Router /age-rating-categories/{id}/movies [get]
func GetMoviesByRatingId(c *gin.Context) { // Get model if exist
	var movies []models.Movie

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("age_rating_category_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// UpdateAgeRatingCategory godoc
// @Summary Update AgeRatingCategory.
// @Description Update AgeRatingCategory by id.
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "AgeRatingCategory id"
// @Param Body body ageRatingCategoryInput true "the body to update age rating category"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [patch]
func UpdateRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ageRatingCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.AgeRatingCategory
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// DeleteAgeRatingCategory godoc
// @Summary Delete one AgeRatingCategory.
// @Description Delete a AgeRatingCategory by id.
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "AgeRatingCategory id"
// @Success 200 {object} map[string]boolean
// @Router /age-rating-categories/{id} [delete]
func DeleteRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
