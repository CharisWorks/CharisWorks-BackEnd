package users

/*
import (

	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"

)

	func ExampleUser(UserId string) User {
		// ユーザー型の例を生成
		user := User{
			UserId: UserId,
			UserProfile: &UserProfile{
				DisplayName: "John Doe",
				Description: "Example user",
				CreatedAt:   time.Now(),
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
			Manufacturer: Manufacturer{
				StripeAccountId: stripe.String("acct_1Okj9YPFjznovTf3"),
			},
		}
		return user

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
			Manufacturer: Manufacturer{
				StripeAccountId: stripe.String("acct_1OkjHjPKEl3posmB"),
			},
		}
		return user
	}

	func ExampleUser3(UserId string) User {
		// ユーザー型の例を生成
		user := User{
			UserId: UserId,
			UserProfile: &UserProfile{
				DisplayName: "John Doe",
				Description: "Example user",
				CreatedAt:   time.Now(),
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
			Manufacturer: Manufacturer{
				StripeAccountId: new(string),
			},
		}
		return user
	}

	func ExampleUser4(UserId string) User {
		// ユーザー型の例を生成
		user := User{
			UserId: UserId,
			UserProfile: &UserProfile{
				DisplayName: "John Doe",
				Description: "Example user",
				CreatedAt:   time.Now(),
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
			Manufacturer: Manufacturer{
				StripeAccountId: nil,
			},
		}
		return user
	}

type ExampleUserRequests struct {
}

	func (u ExampleUserRequests) Get(UserId string, ctx *gin.Context, UserDB IUserRepository) (*User, error) {
		log.Println("UserId: ", UserId)
		if UserId == "zbGHo5B74BNcnzEYFDIqmmtdtF82" {
			log.Print("hoge")
			user := ExampleUser(UserId)
			return &user, nil
		}
		if UserId == "zr4Fx8qwumep5nlTW4faAGQ1vHp1" {
			log.Print("fuga")
			user := ExampleUser2(UserId)
			return &user, nil
		}
		if UserId == "Ep5mTE4BGxPWnxoFmj9tjzqlHoA3" {
			log.Print("foo")
			user := ExampleUser3(UserId)
			return &user, nil
		}
		if UserId == "7EQAcIwLSYYuN9cZvCoTdC0H11v1" {
			log.Print("bar")
			user := ExampleUser4(UserId)
			return &user, nil
		}
		return nil, nil
	}

	func (u ExampleUserRequests) Delete(UserId string, ctx *gin.Context) error {
		log.Println("UserId: ", UserId)
		return nil
	}

	func (u ExampleUserRequests) UserProfileRegister(p UserProfileRegisterPayload, ctx *gin.Context) error {
		log.Println("UserProfileRegisterPayload: ", p)
		return nil
	}

	func (u ExampleUserRequests) ProfileUpdate(p UserProfile, ctx *gin.Context) error {
		log.Println("UserProfile: ", p)
		return nil
	}

	func (u ExampleUserRequests) AddressRegister(p AddressRegisterPayload, ctx *gin.Context) error {
		log.Println("AddressRegisterPayload: ", p)
		return nil
	}

	func (u ExampleUserRequests) AddressUpdate(p UserAddress, ctx *gin.Context) error {
		log.Println("UserAddress: ", p)
		return nil
	}

	func (u ExampleUserRequests) UserCreate(UserId string, ctx *gin.Context) error {
		log.Println("Create User for DB")
		return nil
	}
*/
/* type ExampleUserUtils struct {
}

func (r ExampleUserUtils) InspectAddressRegisterPayload(p AddressRegisterPayload) error {
	return nil
}
func (r ExampleUserUtils) InspectAddressUpdatePayload(p UserAddress) map[string]interface{} {
	return nil
}
func (r ExampleUserUtils) InspectProfileUpdatePayload(p UserProfile) map[string]interface{} {
	return nil
}
*/
