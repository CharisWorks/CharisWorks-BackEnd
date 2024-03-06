package users

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestInspectProfileUpdatePayload(t *testing.T) {
	UserUtils := UserUtils{}
	Cases := []struct {
		name    string
		payload UserProfile
		want    map[string]interface{}
	}{
		{
			name: "正常",
			payload: UserProfile{
				DisplayName: "test",
				Description: "test",
			},
			want: map[string]interface{}{
				"display_name": "test",
				"description":  "test",
			},
		},
		{
			name: "変更できないカラムがあっても無視される",
			payload: UserProfile{
				DisplayName:     "test",
				Description:     "test",
				StripeAccountId: "test",
			},
			want: map[string]interface{}{
				"display_name": "test",
				"description":  "test",
			},
		},
		{
			name: "どちらかだけの変更も可能",
			payload: UserProfile{
				DisplayName: "test",
			},
			want: map[string]interface{}{
				"display_name": "test",
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			got := UserUtils.InspectProfileUpdatePayload(tt.payload)
			if got["display_name"] != tt.want["display_name"] {
				t.Errorf("got: %v, want: %v", got["display_name"], tt.want["display_name"])
			}
			if got["description"] != tt.want["description"] {
				t.Errorf("got: %v, want: %v", got["description"], tt.want["description"])
			}
		})
	}
}

func TestInspectAddressUpdatePayload(t *testing.T) {
	UserUtils := UserUtils{}
	Cases := []struct {
		name    string
		payload UserAddress
		want    map[string]interface{}
		err     error
	}{
		{
			name: "正常",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "テスト",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"address_3":       "test",
				"phone_number":    "000-0000-0000",
			},
		},
		{
			name: "郵便番号が変換されるか",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "0123456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "テスト",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
		},
		{
			name: "電話番号が変換されるか",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "00000000000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "テスト",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
		},
		{
			name: "電話番号の変化エラー",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-00000000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "テスト",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "カタカナじゃない",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "test",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "00000000000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "test",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "カタカナじゃない",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "test",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "00000000000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "テスト",
				"last_name":       "test",
				"last_name_kana":  "test",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "郵便番号エラー",
			payload: UserAddress{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-34567",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "00000000000",
			},
			want: map[string]interface{}{
				"first_name":      "test",
				"first_name_kana": "test",
				"last_name":       "test",
				"last_name_kana":  "テスト",
				"zip_code":        "012-3456",
				"address_1":       "test",
				"address_2":       "test",
				"phone_number":    "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserUtils.InspectAddressUpdatePayload(tt.payload)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("got: %v, want: %v", err, tt.err)
				}
				return
			}
			if got["first_name"] != tt.want["first_name"] {
				t.Errorf("got: %v, want: %v", got["first_name"], tt.want["first_name"])
			}
			if got["first_name_kana"] != tt.want["first_name_kana"] {
				t.Errorf("got: %v, want: %v", got["first_name_kana"], tt.want["first_name_kana"])
			}
			if got["last_name"] != tt.want["last_name"] {
				t.Errorf("got: %v, want: %v", got["last_name"], tt.want["last_name"])
			}
			if got["last_name_kana"] != tt.want["last_name_kana"] {
				t.Errorf("got: %v, want: %v", got["last_name_kana"], tt.want["last_name_kana"])
			}
			if got["zip_code"] != tt.want["zip_code"] {
				t.Errorf("got: %v, want: %v", got["zip_code"], tt.want["zip_code"])
			}
			if got["address_1"] != tt.want["address_1"] {
				t.Errorf("got: %v, want: %v", got["address_1"], tt.want["address_1"])
			}
			if got["address_2"] != tt.want["address_2"] {
				t.Errorf("got: %v, want: %v", got["address_2"], tt.want["address_2"])
			}
			if got["phone_number"] != tt.want["phone_number"] {
				t.Errorf("got: %v, want: %v", got["phone_number"], tt.want["phone_number"])
			}
		})
	}
}

func TestInspectAddressRegisterPayload(t *testing.T) {
	UserUtils := UserUtils{}
	Cases := []struct {
		name    string
		payload AddressRegisterPayload
		want    AddressRegisterPayload
		err     error
	}{
		{
			name: "正常",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
		},
		{
			name: "カタカナじゃない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "test",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "カタカナじゃない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "test",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:    "test",
				LastName:     "test",
				LastNameKana: "テスト",
				ZipCode:      "012-3456",
				Address1:     "test",
				Address2:     "test",
				PhoneNumber:  "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "なにかない",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "電話番号エラー",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "0000000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "郵便番号エラー",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "01234567",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		}, {
			name: "電話番号・郵便番号が変換されるか",
			payload: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "0123456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "00000000000",
			},
			want: AddressRegisterPayload{
				FirstName:     "test",
				FirstNameKana: "テスト",
				LastName:      "test",
				LastNameKana:  "テスト",
				ZipCode:       "012-3456",
				Address1:      "test",
				Address2:      "test",
				PhoneNumber:   "000-0000-0000",
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserUtils.InspectAddressRegisterPayload(tt.payload)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("got: %v, want: %v", err, tt.err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
