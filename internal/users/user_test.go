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
	UserDB := UserDB{DB: db}
	UserUtils := UserUtils{}
	UserRequests := UserRequests{}
	Cases := []struct {
		name                  string
		userId                string
		want                  User
		updateProfile         UserProfile
		wantProfileUpdated    User
		registerAddress       UserAddressRegisterPayload
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
			registerAddress: UserAddressRegisterPayload{
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
			registerAddress: UserAddressRegisterPayload{
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
			registerAddress: UserAddressRegisterPayload{
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
			err := UserRequests.UserCreate(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}
			// read
			user, err := UserRequests.UserGet(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(*user, tt.want) {
				t.Errorf("got: %v, want: %v", user, tt.want)
			}
			// update
			err = UserRequests.UserProfileUpdate(tt.userId, tt.updateProfile, UserDB, UserUtils)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.UserGet(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(*user, tt.wantProfileUpdated) {
				t.Errorf("got: %v, want: %v", user, tt.wantProfileUpdated)
			}
			err = UserRequests.UserAddressRegister(tt.userId, tt.registerAddress, UserDB, UserUtils)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.UserGet(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(*user, tt.wantAddressRegistered) {
				t.Errorf("got: %v, want: %v", user, tt.wantAddressRegistered)
			}
			err = UserRequests.UserAddressUpdate(tt.userId, tt.updateAddress, UserDB, UserUtils)
			if err != nil {
				t.Errorf("error")
			}
			user, err = UserRequests.UserGet(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}
			if CompareUser(*user, tt.wantAddressUpdated) {
				t.Errorf("got: %v, want: %v", user, tt.wantAddressUpdated)
			}

			// delete
			err = UserRequests.UserDelete(tt.userId, UserDB)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
}
