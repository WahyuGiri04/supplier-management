package model

import (
	baseModel "supplier-management/model/base"

	"gorm.io/gorm"
)

type Supplier struct {
	baseModel.BaseModel
	Code     string `json:"code" gorm:"size:255"`
	Name     string `json:"name" gorm:"size:255"`
	Nickname string `json:"nickname" gorm:"size:255"`
	Logo     string `json:"logo" gorm:"size:500"`

	SupplierGroups []SupplierGroup `json:"supplierGroup" gorm:"foreignKey:SupplierID"`
	Addresses      []Address       `json:"address" gorm:"foreignKey:SupplierID"`
	Contacts       []Contact       `json:"contact" gorm:"foreignKey:SupplierID"`
}

func (Supplier) TableName() string {
	return "supplier_management.supplier"
}

func (Supplier) PreloadFields(db *gorm.DB) *gorm.DB {
	return db.
		Preload("SupplierGroups").
		Preload("Addresses").
		Preload("Contacts")
}
