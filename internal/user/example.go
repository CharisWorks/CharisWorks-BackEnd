package user

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ExampleUser(UserId string) User {
	e := new(User)
	e.UserId = UserId

	return *e

}

func ExampleUser2(UserId string) User {
	// ユーザー型の例を生成
	user := User{
		UserId: UserId,
		UserProfile: &UserProfile{
			DisplayName:    "John Doe",
			Description:    "Example user",
			CreatedAt:      time.Now(),
			IsManufacturer: true,
		},
		UserAddress: &UserAddress{
			LastName:      "適当",
			LastNameKana:  "テキトウ",
			FirstName:     "太郎",
			FirstNameKana: "タロウ",
			ZipCode:       "3050821",
			Address1:      "茨城県つくば市春日",
			Address2:      "天王台1-1-1",
			Address3:      nil,
			PhoneNumber:   "+81 80 12345678",
		},
		Manufacturer: &Manufacturer{
			StripeAccountId: "acct_1Okj9YPFjznovTf3",
		},
	}
	return user
}

type ExampleUserRequests struct {
}

func (u ExampleUserRequests) UserGet(UserId string, ctx *gin.Context) (*User, error) {
	log.Println("UserId: ", UserId)
	user := ExampleUser2(UserId)
	return &user, nil
}
func (u ExampleUserRequests) UserDelete(UserId string, ctx *gin.Context) error {
	log.Println("UserId: ", UserId)
	return nil
}
func (u ExampleUserRequests) UserProfileRegister(p UserProfileRegisterPayload, ctx *gin.Context) error {
	log.Println("UserProfileRegisterPayload: ", p)
	return nil
}
func (u ExampleUserRequests) UserProfileUpdate(p UserProfile, ctx *gin.Context) error {
	log.Println("UserProfile: ", p)
	return nil
}
func (u ExampleUserRequests) UserAddressRegister(p UserAddressRegisterPayload, ctx *gin.Context) error {
	log.Println("UserAddressRegisterPayload: ", p)
	return nil
}
func (u ExampleUserRequests) UserAddressUpdate(p UserAddress, ctx *gin.Context) error {
	log.Println("UserAddress: ", p)
	return nil
}
