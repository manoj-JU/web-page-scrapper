package models

import "gorm.io/gorm"

// mysql table name for amazon products table
const (
	AmazonProductCName = "amazon_products"
)

// AmazonProduct defines the db structure of amazon product
type AmazonProduct struct {
	gorm.Model
	Title        string `gorm:"primaryKey;index"`
	ImageURL     string ``
	Description  string ``
	Price        string ``
	TotalReviews string ``
}

// CName returns collection name or table name for amazon product model
func (p *AmazonProduct) CName() string {
	return AmazonProductCName
}
