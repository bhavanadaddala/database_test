package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"golang.org/x/crypto/bcrypt"
)

type Product struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	os.Remove("test.db")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})
	user := username()
	hash := hashvalue()

	// Create
	db.Create(&Product{Username: user, Password: hash})

	// Read
	var product Product
	// db.First(&product, 1) // find product with id 1
	db.First(&product, "username = ?", user) // find product with code l1212
	fmt.Println(product)
	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&product)

}

//Ask user to enter username and password
func username() string {
	var temp string
	fmt.Println("enter username")
	fmt.Scanln(&temp)

	return temp

}

func hashvalue() string {
	var pwd []byte
	fmt.Println("enter password")
	fmt.Scanln(&pwd)
	fmt.Println(string(pwd))
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	fmt.Printf("%x\n", hash)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
