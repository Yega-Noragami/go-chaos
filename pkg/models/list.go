package models

import (
	"github.com/Yega-Noragami/go-chaos/pkg/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type List struct {
	gorm.Model
	TempKey   string `json:"Tempkey"`
	TempValue string `json:"Tempvalue"`
}

func SETUP_DB() {
	db = database.Setup()
}

func (l *List) CreateList() *List {
	db.NewRecord(l)
	db.Create(&l)
	return l
}

func GetAllList() []List {
	var Lists []List
	db.Order("created_at desc").Find(&Lists)
	return Lists
}

func GetListByKey(TempKey string) (*List, *gorm.DB) {
	var getList List
	db := db.Where("Tempkey=?", TempKey).Find(&getList)
	return &getList, db
}

func DeleteList(TempKey string) List {
	var list List
	db.Where("TempKey=?", TempKey).Delete(list)
	return list
}
