package main

import (
	"log"
	"net/http"

	"github.com/wandermaia/pos-golang-apis/configs"
	"github.com/wandermaia/pos-golang-apis/internal/entity"
	"github.com/wandermaia/pos-golang-apis/internal/infra/database"
	"github.com/wandermaia/pos-golang-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)

	log.Println("Servidor iniciado!")
	http.ListenAndServe(":8000", nil)

}

/*

sqlite3 cmd/server/test.db
select * from products;
.exit


*/
