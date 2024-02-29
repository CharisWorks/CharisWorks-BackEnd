package user

import (
	"log"
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/stripe/stripe-go/v76"
)

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
			if reflect.DeepEqual(*User, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.want)
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
			if reflect.DeepEqual(*User, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.want)
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
		{
			name:   "address3の挙動のテスト(アドレスが削除された)",
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
				"address_3": nil,
			},
			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      stripe.String(""),
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
			if reflect.DeepEqual(*User, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.want)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
