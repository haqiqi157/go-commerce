package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Role     string    `json:"role"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Auditable
	Transactions []Transaction `json:"transactions"`
}

func NewUser(email, password, role, address, phone string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Password:  password,
		Role:      role,
		Address:   address,
		Phone:     phone,
		Auditable: NewAuditable(),
	}
}

func UpdateUser(id uuid.UUID, email, password, role, address, phone string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Password:  password,
		Role:      role,
		Address:   address,
		Phone:     phone,
		Auditable: UpdateAuditable(),
	}
}
