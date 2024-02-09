package user

import "time"

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
