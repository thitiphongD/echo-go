package usecase

import "github.com/thitiphongD/echo-go/src/domain"

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
