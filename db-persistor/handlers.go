package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manoj-JU/web-page-scrapper/models"
	"github.com/manoj-JU/web-page-scrapper/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//returns mysql URL
func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.userName, dbConf.password, dbConf.host, dbConf.port, dbConf.dbName)
}

// Add new amazon product or update existing amazon product
func CreateAmazonProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var amazonProduct AmazonProduct
	err := json.NewDecoder(r.Body).Decode(&amazonProduct)
	if err != nil {
		utils.HandleError(err, http.StatusInternalServerError, w)
		return
	}

	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.HandleError(err, http.StatusInternalServerError, w)
		return
	}

	amazonProductModel := R2MAmazonProduct(&amazonProduct)

	result := db.First(&amazonProductModel, "title = ?", amazonProductModel.Title)
	if result.Error != nil {
		//create new record
		result = db.Create(&amazonProductModel)
		if result.Error != nil {
			utils.HandleError(result.Error, http.StatusInternalServerError, w)
			return
		}

	} else {
		//update existing record
		result = db.Save(&amazonProductModel)
		if result.Error != nil {
			utils.HandleError(result.Error, http.StatusInternalServerError, w)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&amazonProductModel)
}

// Get all amazon products
func GetAmazonProductsHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if there is an error opening the connection, handle it
	if err != nil {
		utils.HandleError(err, http.StatusInternalServerError, w)
		return
	}

	amazonProducts := []AmazonProduct{}
	// Get all records
	result := db.Find(&amazonProducts)
	if result.Error != nil {
		utils.HandleError(result.Error, http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(amazonProducts)
}

// R2MAmazonProduct converts the amazonProduct from request to amazonProduct db model
func R2MAmazonProduct(req *AmazonProduct) *models.AmazonProduct {
	return &models.AmazonProduct{
		Title:        req.Title,
		ImageURL:     req.ImageURL,
		Description:  req.Description,
		Price:        req.Price,
		TotalReviews: req.TotalReviews,
	}
}
