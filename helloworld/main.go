package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)

type Car struct {
	ID int 					`gorm:"column:ID;primary_key"`
	LPC string				`gorm:"column:LPC"`
	IDowner int				`gorm:"column:IDowner"`
	Color string			`gorm:"column:Color"`
	Lable string			`gorm:"column:Lable"`
	KindOf string			`gorm:"column:KindOf"` //all need colmum name
	CurrentKM int			`gorm:"column:CurrentKM"`
	Price int				`gorm:"column:Price"`
	CarStatus string 		`gorm:"column:CarStatus;DEFAULT:'available'"`
}

// TableName set name table for struct
func (ca Car) TableName() string {
	return "cars"
}

func main(){
	db, err := gorm.Open("mysql", "root:1@(localhost:3306)/rentcar?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic("Open database error")
	} else {
		fmt.Println("Connected to DB")
	}

	minhCar := Car{
//		ID: 1,
		LPC: "98 M3 2468",
		IDowner: 1,
		Color: "Red",
		Lable: "Dream",
		KindOf: "car",
		CurrentKM:1536,
		Price:500,
	//	CarStatus:"available",
	}

	//insert a record
	fmt.Println(minhCar.TableName())
	if db.NewRecord(minhCar) {
		fmt.Println("primari Key is blank")
		db.Create(&minhCar)
	} else {
		fmt.Println("primari Key existed")
	}

	// query car 
	var allCar []Car
	if db.Find(&allCar).Error != nil {
		fmt.Println("[LHM] query false")
	} else {
		fmt.Println("[LHM] query success, all car ====>", allCar)
	}
}