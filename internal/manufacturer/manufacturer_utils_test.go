package manufacturer

import (
	"log"
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestInspectedRegisterPayload(t *testing.T) {
	ManufacturerUtils := ManufacturerUtils{}
	Cases := []struct {
		name    string
		payload RegisterPayload
		err     error
	}{
		{
			name: "正常",
			payload: RegisterPayload{
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
			payload: RegisterPayload{
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
			payload: RegisterPayload{
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
			payload: RegisterPayload{
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
			payload: RegisterPayload{
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
			payload: RegisterPayload{
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
		payload UpdatePayload
		want    map[string]interface{}
	}{
		{
			name: "正常",
			payload: UpdatePayload{
				Price:       1000,
				Name:        "test",
				Stock:       2,
				Size:        3,
				Description: "test",
				Tags:        []string{"aaa", "bbb"},
			},
			want: map[string]interface{}{
				"price":       1000,
				"name":        "test",
				"stock":       2,
				"size":        3,
				"description": "test",
				"tags":        []string{"aaa", "bbb"},
			},
		},
		{
			name: "price error",
			payload: UpdatePayload{
				Price: 0,
				Name:  "test",
			},
			want: map[string]interface{}{
				"name": "test",
			},
		},
		{
			name: "name error",
			payload: UpdatePayload{
				Name:  "",
				Price: 1000,
			},
			want: map[string]interface{}{
				"price": 1000,
			},
		},
		{
			name: "stock error",
			payload: UpdatePayload{
				Stock: 0,
				Tags:  []string{"aaa", "bbb"},
			},
			want: map[string]interface{}{
				"tags": []string{"aaa", "bbb"},
			},
		},
		{
			name: "size error",
			payload: UpdatePayload{
				Size: 0,
				Tags: []string{"aaa", "bbb"},
			},
			want: map[string]interface{}{
				"tags": []string{"aaa", "bbb"},
			},
		},
		{
			name: "description error",
			payload: UpdatePayload{
				Description: "",
				Tags:        []string{"aaa", "bbb"},
			},
			want: map[string]interface{}{
				"tags": []string{"aaa", "bbb"},
			},
		},
		{
			name: "tags error",
			payload: UpdatePayload{
				Tags:  []string{},
				Price: 1000,
			},
			want: map[string]interface{}{
				"price": 1000,
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := ManufacturerUtils.Update(tt.payload)
			log.Print(payload, err)
			if err != nil {
				log.Print(err.Error())
			}
			if !reflect.DeepEqual(payload, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, payload, tt.want)
			}

		})
	}
}
