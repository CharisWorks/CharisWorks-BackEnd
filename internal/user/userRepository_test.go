package user

import (
	"log"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/stripe/stripe-go/v76"
)

func Test_UserDB_Create_User(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")

	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name   string
		userId string
		want   User
	}{
		{
			name:   "正常",
			userId: "aaa",
			want: User{
				UserId: "aaa",
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDB.CreateUser(tt.userId, 1)
			if err != nil {
				t.Errorf("error")
			}

			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

			log.Print(&User)
			if User.UserId != tt.want.UserId {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserId, tt.want.UserId)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
func Test_UserDB_Update_Profile(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")

	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name    string
		userId  string
		payload map[string]interface{}
		want    User
	}{
		{
			name:   "正常",
			userId: "aaa",
			payload: map[string]interface{}{
				"display_name": "display_name",
				"description":  "description",
			},
			want: User{
				UserId: "aaa",
				UserProfile: UserProfile{
					DisplayName: "display_name",
					Description: "description",
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDB.CreateUser(tt.userId, 1)
			if err != nil {
				t.Errorf("error")
			}
			err = UserDB.UpdateProfile(tt.userId, tt.payload)
			log.Print(err)
			if err != nil {
				t.Errorf("error")
			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

			log.Print(&User)
			if User.UserId != tt.want.UserId {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserId, tt.want.UserId)
			}
			if User.UserProfile.DisplayName != tt.want.UserProfile.DisplayName {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserProfile.DisplayName, tt.want.UserProfile.DisplayName)
			}
			if User.UserProfile.Description != tt.want.UserProfile.Description {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserProfile.Description, tt.want.UserProfile.Description)
			}

			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
func Test_UserDB_Register_Address(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")

	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name      string
		userId    string
		shippings UserAddressRegisterPayload
		want      User
	}{
		{
			name:   "正常",
			userId: "aaa",
			shippings: UserAddressRegisterPayload{
				ZipCode:       "000-0000",
				Address1:      "abc",
				Address2:      "def",
				Address3:      stripe.String("ghi"),
				PhoneNumber:   "000-0000-0000",
				FirstName:     "適当",
				FirstNameKana: "テキトウ",
				LastName:      "太郎",
				LastNameKana:  "タロウ",
			},
			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      stripe.String("ghi"),
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDB.CreateUser(tt.userId, 1)
			if err != nil {
				t.Errorf("error")
			}
			err = UserDB.RegisterAddress(tt.userId, tt.shippings)
			log.Print(err)
			if err != nil {
				t.Errorf("error")

			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

			log.Print(&User)
			if User.UserId != tt.want.UserId {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserId, tt.want.UserId)
			}
			if User.UserAddress.Address1 != tt.want.UserAddress.Address1 {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.Address1, tt.want.UserAddress.Address1)
			}
			if User.UserAddress.Address2 != tt.want.UserAddress.Address2 {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.Address2, tt.want.UserAddress.Address2)
			}
			if *User.UserAddress.Address3 != *tt.want.UserAddress.Address3 {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.Address3, tt.want.UserAddress.Address3)
			}
			if User.UserAddress.FirstName != tt.want.UserAddress.FirstName {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.FirstName, tt.want.UserAddress.FirstName)
			}
			if User.UserAddress.FirstNameKana != tt.want.UserAddress.FirstNameKana {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.FirstNameKana, tt.want.UserAddress.FirstNameKana)
			}
			if User.UserAddress.LastName != tt.want.UserAddress.LastName {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.LastName, tt.want.UserAddress.LastName)
			}
			if User.UserAddress.LastNameKana != tt.want.UserAddress.LastNameKana {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.LastNameKana, tt.want.UserAddress.LastNameKana)
			}
			if User.UserAddress.PhoneNumber != tt.want.UserAddress.PhoneNumber {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.PhoneNumber, tt.want.UserAddress.PhoneNumber)
			}
			if User.UserAddress.ZipCode != tt.want.UserAddress.ZipCode {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.ZipCode, tt.want.UserAddress.ZipCode)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
func Test_UserDB_Update_Address(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name          string
		userId        string
		shippings     UserAddressRegisterPayload
		updatePayload map[string]interface{}
		want          User
	}{
		{
			name:   "正常",
			userId: "aaa",
			shippings: UserAddressRegisterPayload{
				ZipCode:       "000-0000",
				Address1:      "abc",
				Address2:      "def",
				Address3:      stripe.String("ghi"),
				PhoneNumber:   "000-0000-0000",
				FirstName:     "適当",
				FirstNameKana: "テキトウ",
				LastName:      "太郎",
				LastNameKana:  "タロウ",
			},
			updatePayload: map[string]interface{}{
				"address_1": "123",
			},
			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "123",
					Address2:      "def",
					Address3:      stripe.String("ghi"),
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
		},
		{
			name:   "address3の挙動のテスト(アドレスが新たに登録された)",
			userId: "aaa",
			shippings: UserAddressRegisterPayload{
				ZipCode:       "000-0000",
				Address1:      "abc",
				Address2:      "def",
				Address3:      nil,
				PhoneNumber:   "000-0000-0000",
				FirstName:     "適当",
				FirstNameKana: "テキトウ",
				LastName:      "太郎",
				LastNameKana:  "タロウ",
			},
			updatePayload: map[string]interface{}{
				"address_3": "123",
			},
			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      stripe.String("123"),
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDB.CreateUser(tt.userId, 1)
			if err != nil {
				t.Errorf("error")
			}
			err = UserDB.RegisterAddress(tt.userId, tt.shippings)
			log.Print(err)
			if err != nil {
				t.Errorf("error")

			}
			err = UserDB.UpdateAddress(tt.userId, tt.updatePayload)
			log.Print(err)
			if err != nil {
				t.Errorf("error")

			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

			log.Print(&User)
			if User.UserId != tt.want.UserId {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserId, tt.want.UserId)
			}
			if User.UserAddress.Address1 != tt.want.UserAddress.Address1 {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.Address1, tt.want.UserAddress.Address1)
			}
			if User.UserAddress.Address2 != tt.want.UserAddress.Address2 {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.Address2, tt.want.UserAddress.Address2)
			}
			if *User.UserAddress.Address3 != *tt.want.UserAddress.Address3 {
				t.Errorf("%v,got,%v,want%v", tt.name, *User.UserAddress.Address3, *tt.want.UserAddress.Address3)
			}
			if User.UserAddress.FirstName != tt.want.UserAddress.FirstName {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.FirstName, tt.want.UserAddress.FirstName)
			}
			if User.UserAddress.FirstNameKana != tt.want.UserAddress.FirstNameKana {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.FirstNameKana, tt.want.UserAddress.FirstNameKana)
			}
			if User.UserAddress.LastName != tt.want.UserAddress.LastName {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.LastName, tt.want.UserAddress.LastName)
			}
			if User.UserAddress.LastNameKana != tt.want.UserAddress.LastNameKana {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.LastNameKana, tt.want.UserAddress.LastNameKana)
			}
			if User.UserAddress.PhoneNumber != tt.want.UserAddress.PhoneNumber {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.PhoneNumber, tt.want.UserAddress.PhoneNumber)
			}
			if User.UserAddress.ZipCode != tt.want.UserAddress.ZipCode {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserAddress.ZipCode, tt.want.UserAddress.ZipCode)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
