package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	dsn := "root:yb123456@tcp(192.168.88.130:3306)/golang2_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}
type User struct {
	gorm.Model
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []Language `gorm:"many2many:user_languages;"`
}
type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}
type Address struct {
	gorm.Model
	Address1 string
	UserID   uint
}

func create30() {
	db.AutoMigrate(&User{}, &Email{}, &Address{}, &Language{})

}
func test1() {
	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails: []Email{
			{Email: "jinzhu@example.com"},
			{Email: "jinzhu-2@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	db.Create(&user)
}
func test2() {
	var user User
	db.First(&user)
	var languages []Language
	db.Model(&user).Association("Languages").Find(&languages)
	fmt.Printf("Languages: %v\n", languages)
}
func test3() {
	var user User
	db.First(&user)
	i := db.Model(&user).Association("Languages").Count()
	fmt.Printf("Languages: %v\n", i)
}
func main() {
	test3()
}
