package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
)

type IStripeRequests interface {
	GetClientSecret(userId string, CartRequests cart.IRequests, cartRepository cart.IRepository, CartUtils cart.IUtils) (url *string, err error)
	GetRegisterLink(email string, user users.User, UserDB users.IRepository) (url *string, err error)
	GetStripeMypageLink(stripeAccountId string) (url *string, err error)
}
