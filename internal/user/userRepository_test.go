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
		t.Errorf(err.Error())

	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name        string
		userId      string
		payload     map[string]interface{}
		want        User
		wantUpdated User
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
			},
			wantUpdated: User{
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
				t.Errorf(err.Error())
			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}

			log.Print("got: ", *User, "want: ", tt.want)
			if CompareUser(*User, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.want)
			}
			err = UserDB.UpdateProfile(tt.userId, tt.payload)
			log.Print(err)
			if err != nil {
				t.Errorf(err.Error())
			}
			User, err = UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}

			log.Print("got: ", *User, "want: ", tt.wantUpdated)

			if CompareUser(*User, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.wantUpdated)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}

		})
	}
}

func Test_UserDB_Register_Update_Address(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name          string
		userId        string
		shippings     UserAddressRegisterPayload
		updatePayload map[string]interface{}
		want          User
		wantUpdated   User
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
					Address3:      "ghi",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
			updatePayload: map[string]interface{}{
				"address_1": "123",
			},
			wantUpdated: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "123",
					Address2:      "def",
					Address3:      "ghi",
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
			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
			updatePayload: map[string]interface{}{
				"address_3": "123",
			},
			wantUpdated: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      "123",
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

			want: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      "ghi",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "適当",
					FirstNameKana: "テキトウ",
					LastName:      "太郎",
					LastNameKana:  "タロウ",
				},
			},
			updatePayload: map[string]interface{}{
				"address_3": nil,
			},
			wantUpdated: User{
				UserId: "aaa",
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "abc",
					Address2:      "def",
					Address3:      "",
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
				t.Errorf(err.Error())
			}
			err = UserDB.RegisterAddress(tt.userId, tt.shippings)
			if err != nil {
				t.Errorf(err.Error())

			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}
			if CompareUser(*User, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.want)
			}
			err = UserDB.UpdateAddress(tt.userId, tt.updatePayload)
			if err != nil {
				t.Errorf(err.Error())

			}
			User, err = UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}
			if CompareUser(*User, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *User, tt.wantUpdated)
			}
			err = UserDB.DeleteUser(tt.userId)
			if err != nil {
				t.Errorf(err.Error())
			}

		})
	}
}

func CompareUser(got User, want User) bool {
	gotUser := User{
		UserId: got.UserId,
		UserProfile: UserProfile{
			DisplayName:     got.UserProfile.DisplayName,
			Description:     got.UserProfile.Description,
			StripeAccountId: got.UserProfile.Description,
		},
		UserAddress: got.UserAddress,
	}
	wantUser := User{
		UserId: want.UserId,
		UserProfile: UserProfile{
			DisplayName:     want.UserProfile.DisplayName,
			Description:     want.UserProfile.Description,
			StripeAccountId: want.UserProfile.Description,
		},
		UserAddress: want.UserAddress,
	}

	return !reflect.DeepEqual(gotUser, wantUser)
}
