package users

import (
	"regexp"
	"unicode"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

type UserUtils struct {
}

func (u UserUtils) InspectProfileUpdatePayload(profile UserProfile) map[string]interface{} {
	updatepayload := make(map[string]interface{})
	if len(profile.DisplayName) > 1 {
		updatepayload["display_name"] = profile.DisplayName
	}
	if len(profile.Description) > 1 {
		updatepayload["description"] = profile.Description
	}

	return updatepayload
}
func (u UserUtils) InspectAddressUpdatePayload(address UserAddress) (map[string]interface{}, error) {
	conditions := make(map[string]interface{})
	if len(address.FirstName) > 1 {
		conditions["first_name"] = address.FirstName
	}
	if len(address.FirstNameKana) > 1 {
		if IsKatakana(address.FirstNameKana) {
			conditions["first_name_kana"] = address.FirstNameKana
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	if len(address.LastName) > 1 {
		conditions["last_name"] = address.LastName
	}
	if len(address.LastNameKana) > 1 {
		if IsKatakana(address.LastNameKana) {
			conditions["last_name_kana"] = address.LastNameKana
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	if len(address.ZipCode) > 1 {
		if IsValidPostalCode(address.ZipCode) {
			conditions["zip_code"] = address.ZipCode
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	if len(address.Address1) > 1 {
		conditions["address_1"] = address.Address1
	}
	if len(address.Address2) > 1 {
		conditions["address_2"] = address.Address2
	}
	if len(address.Address3) > 1 {
		conditions["address_3"] = address.Address3
	}
	if len(address.PhoneNumber) > 1 {
		isValid, phoneNumber := ConvertPhoneNumber(address.PhoneNumber)
		if isValid {
			conditions["phone_number"] = phoneNumber
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	return conditions, nil
}
func (u UserUtils) InspectAddressRegisterPayload(address UserAddressRegisterPayload) (UserAddressRegisterPayload, error) {
	if len(address.FirstName) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.FirstNameKana) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if !IsKatakana(address.FirstNameKana) {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.LastName) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.LastNameKana) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if !IsKatakana(address.LastNameKana) {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.ZipCode) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if !IsValidPostalCode(address.ZipCode) {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.Address1) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.Address2) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(*address.Address3) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if len(address.PhoneNumber) < 1 {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	isValid, phoneNumber := ConvertPhoneNumber(address.PhoneNumber)
	if !isValid {
		return address, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	address.PhoneNumber = phoneNumber

	return address, nil
}
func IsKatakana(str string) bool {
	for _, r := range str {
		if !unicode.In(r, unicode.Katakana) {
			return false
		}
	}
	return true
}

func IsValidPostalCode(postalCode string) bool {
	// 正規表現パターン
	pattern := `^\d{3}-?\d{4}$`
	// 郵便番号の形式にマッチするかどうかを確認
	match, _ := regexp.MatchString(pattern, postalCode)
	return match
}
func ConvertPhoneNumber(phoneNumber string) (bool, string) {
	// 正規表現パターン
	pattern := `^\d{3}-\d{4}-\d{4}$`
	// 電話番号の形式にマッチするかどうかを確認
	match, _ := regexp.MatchString(pattern, phoneNumber)
	if match {
		// 2番目の形式の場合は3番目の形式に変換して返す
		if phoneNumber[3] == '-' {
			return true, phoneNumber
		}
		// 3番目の形式の場合は2番目の形式に変換して返す
		return true, phoneNumber[:3] + "-" + phoneNumber[4:8] + "-" + phoneNumber[9:]
	}
	// 有効な形式でない場合はそのまま返す
	return false, phoneNumber
}
