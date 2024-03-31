package entity

import (
	"github.com/wandermaia/pos-golang-apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"` // Foi criado o pacote importado acima em pkg/entity/id.go
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` //O password nunca vai ser exibido para usuário final.
}

func NewUser(name, email, password string) (*User, error) {
	// gerando um hash da senha informada para salvar no banco de dados.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewId(), // Utilizando a função criada em pkg/entity/id.go
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// Função para validar se o password informado está correto
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

/*
As entidades ajudam a manter a consistência da aplicação.
As regras de negócios ficam nas entidades.
As senhas serão criptografadas utilizando o pacote golang.org/x/crypto/bcrypt

*/
