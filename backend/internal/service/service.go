package service

import (
	"github.com/loloneme/pr12/backend/internal/entities"
	"github.com/loloneme/pr12/backend/internal/repository"
)

type Service struct {
	Drink
	Cart
}

type Drink interface {
	CreateDrink(drink *entities.Drink) (int64, error)

	GetDrinkByID(drinkID int64) (*entities.Drink, error)
	GetDrinks(userID string) ([]*entities.Drink, error)

	UpdateDrink(drink *entities.Drink) error
	ToggleFavorite(drinkID int64, userID string) error

	DeleteDrink(drinkID int64) error
}

type Cart interface {
	AddToCart(cartItem *entities.CartItem) (int64, error)

	GetCart(userID string) ([]*entities.CartItem, error)
	UpdateCartItem(cartItemID int64, quantity int64) error

	DeleteCartItem(cartItemID int64, userID string) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Drink: NewDrinkService(repos.Drink),
		Cart:  NewCartService(repos.Cart),
	}
}
