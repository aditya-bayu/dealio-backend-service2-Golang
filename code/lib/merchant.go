package lib

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "./database"
	"github.com/jinzhu/gorm"
)

// Responses to client
type Responses struct {
	Message string     `json:"message"`
	Data    []Merchant `json:"data"`
}

// Response to client
type Response struct {
	Message string   `json:"message"`
	Data    Merchant `json:"data"`
}

// Merchant table
type Merchant struct {
	ID    uint    `gorm:"primary_key" json:"id"`
	Name  *string `gorm:"size:255" json:"name"`
	Image *string `gorm:"size:255" json:"image"`
}

// TableName for custom tablename of merchant
func (Merchant) TableName() string {
	return "merchant"
}

//ReadMerchants from database
func ReadMerchants(db *gorm.DB) []Merchant {
	var items []Merchant
	db.Find(&items)
	return items
}

// ReadMerchant for reading from database
func ReadMerchant(db *gorm.DB, idx int) Merchant {
	var item Merchant
	db.First(&item, idx)
	return item
}

// GetMerchants for API
func GetMerchants(w http.ResponseWriter, r *http.Request) {
	var db *gorm.DB
	db = database.DB()
	defer db.Close()

	mer := ReadMerchants(db)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Responses{Message: "OK", Data: mer})
}

// GetMerchant for API
func GetMerchant(w http.ResponseWriter, r *http.Request) {
	// var db *gorm.DB
	db := database.DB()
	defer db.Close()

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	item := ReadMerchant(db, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Data: item})
}
