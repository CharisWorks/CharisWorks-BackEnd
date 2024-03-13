package users

import (
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestUserCRUD(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserRequests := Requests{
		UserRepository: UserRepository{DB: db},
		UserUtils:      UserUtils{},
	}
	Cases := []struct {
		name                  string
		userId                string
		want                  User
		updateProfile         UserProfile
		wantProfileUpdated    User
		registerAddress       AddressRegisterPayload
		wantAddressRegistered User
		updateAddress         UserAddress
		wantAddressUpdated    User
	}{
		{
			name:   "正常",
			userId: "test",
			want: User{
				UserId: "test",
			},
			updateProfile: UserProfile{
				DisplayName: "test",
				Description: "test",
			},
			wantProfileUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
			},
			registerAddress: AddressRegisterPayload{
				ZipCode:       "000-0000",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "000-0000-0000",
				FirstName:     "test",
				LastName:      "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
			wantAddressRegistered: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "test",
					Address2:      "test",
					Address3:      "test",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "test",
					LastName:      "test",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
			updateAddress: UserAddress{
				ZipCode:       "000-0000",
				Address1:      "updated",
				Address2:      "updated",
				Address3:      "updated",
				PhoneNumber:   "000-0000-0000",
				FirstName:     "updated",
				LastName:      "updated",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
			wantAddressUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "updated",
					Address2:      "updated",
					Address3:      "updated",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "updated",
					LastName:      "updated",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
		},
		{
			name:   "正常",
			userId: "test",
			want: User{
				UserId: "test",
			},
			updateProfile: UserProfile{
				DisplayName: "test",
			},
			wantProfileUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
				},
			},
			registerAddress: AddressRegisterPayload{
				ZipCode:       "000-0000",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "000-0000-0000",
				FirstName:     "test",
				LastName:      "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
			wantAddressRegistered: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "test",
					Address2:      "test",
					Address3:      "test",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "test",
					LastName:      "test",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
			updateAddress: UserAddress{
				ZipCode:  "000-0000",
				Address1: "updated",
			},
			wantAddressUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "updated",
					Address2:      "test",
					Address3:      "test",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "test",
					LastName:      "test",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
		},
		{
			name:   "電話番号などが自動変換されるか",
			userId: "test",
			want: User{
				UserId: "test",
			},
			updateProfile: UserProfile{
				DisplayName: "test",
				Description: "test",
			},
			wantProfileUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
			},
			registerAddress: AddressRegisterPayload{
				ZipCode:       "0000000",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "00000000000",
				FirstName:     "test",
				LastName:      "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
			wantAddressRegistered: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "test",
					Address2:      "test",
					Address3:      "test",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "test",
					LastName:      "test",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
			updateAddress: UserAddress{
				ZipCode:       "000-0000",
				Address1:      "updated",
				Address2:      "updated",
				Address3:      "updated",
				PhoneNumber:   "000-0000-0000",
				FirstName:     "updated",
				LastName:      "updated",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
			wantAddressUpdated: User{
				UserId: "test",
				UserProfile: UserProfile{
					DisplayName: "test",
					Description: "test",
				},
				UserAddress: UserAddress{
					ZipCode:       "000-0000",
					Address1:      "updated",
					Address2:      "updated",
					Address3:      "updated",
					PhoneNumber:   "000-0000-0000",
					FirstName:     "updated",
					LastName:      "updated",
					FirstNameKana: "テスト",
					LastNameKana:  "テスト",
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			// create
			err := UserRequests.Create(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			// read
			user, err := UserRequests.Get(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(user, tt.want) {
				t.Errorf("got: %v, want: %v", user, tt.want)
			}
			// update
			err = UserRequests.ProfileUpdate(tt.userId, tt.updateProfile)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.Get(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(user, tt.wantProfileUpdated) {
				t.Errorf("got: %v, want: %v", user, tt.wantProfileUpdated)
			}
			err = UserRequests.AddressRegister(tt.userId, tt.registerAddress)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.Get(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(user, tt.wantAddressRegistered) {
				t.Errorf("got: %v, want: %v", user, tt.wantAddressRegistered)
			}
			err = UserRequests.AddressUpdate(tt.userId, tt.updateAddress)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.Get(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(user, tt.wantAddressUpdated) {
				t.Errorf("got: %v, want: %v", user, tt.wantAddressUpdated)
			}

			// delete
			err = UserRequests.Delete(tt.userId)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
