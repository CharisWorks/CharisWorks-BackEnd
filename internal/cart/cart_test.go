package cart

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGET(t *testing.T) {
	e := new(ExapleCartRequest)
	ctx := new(gin.Context)
	_, err := e.Get(ctx)
	if err != nil {
		t.Errorf("error")
	}
}
func TestCartRegister(t *testing.T) {
	e := new(ExapleCartRequest)
	ctx := new(gin.Context)
	goodCases := []struct {
		name     string
		ItemId   string
		Quantity int
	}{
		{
			"正常なパターン", "a", 1,
		},
	}
	for _, tt := range goodCases {
		t.Run(tt.name, func(t *testing.T) {
			p := CartRequestPayload{tt.name, tt.Quantity}
			err := e.Register(p, ctx)
			if err != nil {
				t.Errorf("error")
			}
		})
	}
	badCases := []struct {
		name     string
		ItemId   string
		Quantity int
	}{
		{
			"0はエラー", "b", 0,
		},
		{
			"負数なのでエラー", "c", -1,
		},
	}
	for _, tt := range badCases {
		t.Run(tt.name, func(t *testing.T) {
			p := CartRequestPayload{tt.name, tt.Quantity}
			err := e.Register(p, ctx)
			if err == nil {
				t.Errorf("error")
			}
		})
	}
}
