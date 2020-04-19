package lib

import (
	"encoding/json"
	"net/http"

	Database "./database"
	"github.com/jinzhu/gorm"
)

// Deal table
type Deal struct {
	ID          *int    `gorm:"primary_key" json:"id"`
	Name        *string `gorm:"size:255" json:"name"`
	AudienceID  *int    `json:"audience_id"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Action      *string `json:"action"`
	ActionLink  *string `json:"action_link"`
	MerchantID  *int    `json:"merchant_id"`
	Banner      *string `json:"banner"`
	Type        *string `gorm:"column:type" json:"type"`
	CategoryID  *int    `json:"category_id"`
	HotDeals    *int    `json:"hot_deals"`
}

// ReadDeals reading all deal from database
func ReadDeals(db *gorm.DB) []Deal {
	var items []Deal
	db.Find(&items)
	return items
}

// GetDeals for API
func GetDeals(w http.ResponseWriter, r *http.Request) {
	var db *gorm.DB
	db = Database.DB()

	items := ReadDeals(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
