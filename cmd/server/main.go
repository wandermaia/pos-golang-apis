package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/wandermaia/pos-golang-apis/configs"
	_ "github.com/wandermaia/pos-golang-apis/docs"
	"github.com/wandermaia/pos-golang-apis/internal/entity"
	"github.com/wandermaia/pos-golang-apis/internal/infra/database"
	"github.com/wandermaia/pos-golang-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wander Maia
// @contact.url    http://www.wms.com.br
// @contact.email  wandermaia@yahoo.com.br

// @license.name   GNU GENERAL PUBLIC LICENSE
// @license.url    https://www.gnu.org/licenses/gpl-3.0.html

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	//userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExperesIn)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))

	//r.Use(LogRequest) // Midleware criado de exemplo abaixo.

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json"))) //http://localhost:8000/docs/index.html

	log.Println("Servidor iniciado!")
	http.ListenAndServe(":8000", r)

}

// Exemplo de um midleware para criação de logs, semelhante ao do chi, que foi comentado no código acima.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

/*

sqlite3 cmd/server/test.db
select * from products;
.exit


*/
