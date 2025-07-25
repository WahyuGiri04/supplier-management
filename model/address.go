package model

import baseModel "supplier-management/model/base"

type Address struct {
	baseModel.BaseModel
	SupplierID  uint   `json:"supplierId"`
	Name        string `json:"name" gorm:"size:255"`
	FullAddress string `json:"fullAddress" gorm:"size:255"`
	IsMain      bool   `json:"isMain" gorm:"default:false"`
}

func (Address) TableName() string {
	return "supplier_management.address"
}
