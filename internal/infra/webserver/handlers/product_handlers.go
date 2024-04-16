package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wandermaia/pos-golang-apis/internal/dto"
	"github.com/wandermaia/pos-golang-apis/internal/entity"
	"github.com/wandermaia/pos-golang-apis/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Erro para converter o json: %s", err)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Erro ao criar o produto: %s", err)
		fmt.Print(err)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Erro para gravar o produto no banco de dados: %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
