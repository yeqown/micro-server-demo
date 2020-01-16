package model

import "github.com/jinzhu/gorm"

// FooModel is an go struct to related to SQL-DB table
type FooModel struct {
	gorm.Model
	Bar string
}

// TableName .
func (m FooModel) TableName() string {
	return "foo"
}
