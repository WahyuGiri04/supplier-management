package model

import baseModel "supplier-management/model/base"

type Contact struct {
	baseModel.BaseModel
	SupplierID        uint   `json:"supplierId"`
	Name              string `json:"name" gorm:"size:255"`
	JobPosition       string `json:"jobPosition" gorm:"size:255"`
	Email             string `json:"email" gorm:"size:255"`
	PhoneNumber       string `json:"phoneNumber" gorm:"size:255"`
	MobilePhoneNumber string `json:"mobilePhoneNumber" gorm:"size:255"`
	IsMain            bool   `json:"isMain" gorm:"default:false"`
}

func (Contact) TableName() string {
	return "supplier_management.contacts"
}
