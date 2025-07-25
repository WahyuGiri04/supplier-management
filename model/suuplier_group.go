package model

import baseModel "supplier-management/model/base"

type SupplierGroup struct {
	baseModel.BaseModel
	SupplierID uint   `json:"supplierId"`
	GroupName  string `json:"groupName" gorm:"size:255"`
	// Value      string `json:"value" gorm:"size:255"`
}

func (SupplierGroup) TableName() string {
	return "supplier_management.supplier_group"
}
