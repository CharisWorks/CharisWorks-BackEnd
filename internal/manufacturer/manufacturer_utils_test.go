package manufacturer

import (
	"log"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestInspectedRegisterPayload(t *testing.T) {
	ManufacturerUtils := ManufacturerUtils{}
	Cases := []struct {
		name    string
		payload ItemRegisterPayload
		err     error
	}{
		{
			name: "正常",
			payload: ItemRegisterPayload{
				Name:  "abc",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: nil,
		},
		{
			name: "price error",
			payload: ItemRegisterPayload{
				Name:  "abc",
				Price: 0,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "name error",
			payload: ItemRegisterPayload{
				Name:  "",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "stock error",
			payload: ItemRegisterPayload{
				Name:  "test",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       0,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "size error",
			payload: ItemRegisterPayload{
				Name:  "test",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        0,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "description error",
			payload: ItemRegisterPayload{
				Name:  "test",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := ManufacturerUtils.Register(tt.payload)
			if err != nil {
				log.Print(err.Error())
				if err.Error() != tt.err.Error() {
					t.Errorf("got %v, want %v", err, tt.err)
				}

			}
		})
	}

}

func TestInspectedUpdatePayload(t *testing.T) {
	ManufacturerUtils := ManufacturerUtils{}
	Cases := []struct {
		name    string
		payload map[string]interface{}
		err     error
	}{
		{
			name: "正常",
			payload: map[string]interface{}{
				"Price":       1000,
				"Name":        "test",
				"Stock":       2,
				"Size":        3,
				"Description": "test",
				"Tags":        []string{"aaa", "bbb"},
			},
			err: nil,
		},
		{
			name: "price error",
			payload: map[string]interface{}{
				"Price": 0,
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "name error",
			payload: map[string]interface{}{
				"Name": "",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "stock error",
			payload: map[string]interface{}{
				"Stock": 0,
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "size error",
			payload: map[string]interface{}{
				"Size": 0,
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "description error",
			payload: map[string]interface{}{
				"Description": "",
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "tags error",
			payload: map[string]interface{}{
				"Tags": []string{},
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ManufacturerUtils.Update(tt.payload)
			if err != nil {
				log.Print(err.Error())
				if err.Error() != tt.err.Error() {
					t.Errorf("got %v, want %v", err, tt.err)
				}

			}
		})
	}
}
