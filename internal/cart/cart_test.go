package cart

import (
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func TestCartRequests_Get(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := users.UserDB{DB: db}
	ManufacturerDB := manufacturer.ManufacturerDB{DB: db}
	CartDB := CartDB{DB: db}
	Items := []manufacturer.ItemRegisterPayload{
		{
			Name:  "test1",
			Price: 2000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       2,
				Size:        3,
				Description: "test",
				Tags:        []string{"aaa", "bbb"},
			},
		},
		{
			Name:  "test2",
			Price: 3000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       3,
				Size:        4,
				Description: "test",
				Tags:        []string{"aaa", "ccc"},
			},
		},
	}

	if err = UserDB.CreateUser("aaa", 1); err != nil {
		t.Errorf("error")
	}
	for _, item := range Items {
		err = ManufacturerDB.RegisterItem(item.Name, item, 1, "aaa")
		if err != nil {
			t.Errorf("error")
		}
		err = ManufacturerDB.UpdateItem(map[string]interface{}{"status": items.ItemStatusAvailable}, 2, item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}

	CartRequests := new(CartRequests)
	CartUtils := new(CartUtils)

	Cases := []struct {
		name    string
		payload []CartRequestPayload
		want    *[]Cart
		err     error
	}{
		{
			name: "正常",
			payload: []CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
				{
					ItemId:   "test2",
					Quantity: 2,
				},
			},
			want: &[]Cart{
				{
					ItemId:   "test1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test1",
						Price: 2000,
					},
				},
				{
					ItemId:   "test2",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test2",
						Price: 3000,
					},
				},
			},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest("POST", "/cart", nil)
			body, err := json.Marshal(tt.payload)
			if err != nil {
				t.Errorf("error")
			}

			req.Body = io.NopCloser(strings.NewReader(string(body)))
			log.Println(req.Body)
			ctx.Request = req

			result, err := CartRequests.Get("test", CartDB, CartUtils)
			log.Print(result, tt.want)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, result, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
		})
	}
	for _, item := range Items {
		err = ManufacturerDB.DeleteItem(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	err = UserDB.DeleteUser("aaa")
	if err != nil {
		t.Errorf("error")
	}

}
