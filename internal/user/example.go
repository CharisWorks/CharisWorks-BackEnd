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
			DisplayName: "John Doe",
			Description: "Example user",
			CreatedAt:   time.Now(),
		},
		UserAddress: &UserAddress{
			RealName:    "John Doe",
			ZipCode:     "12345",
			Address1:    "123 Main St",
			Address2:    "Apt 4B",
			Address3:    nil,
			PhoneNumber: "555-1234",
		},
	}
	return user
}

type ExampleUserRequests struct {
}

func (u ExampleUserRequests) UserGet(UserId string, ctx *gin.Context) *User {
	log.Println("UserId: ", UserId)
	user := ExampleUser(UserId)
	return &user
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
