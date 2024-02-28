package user

import (
	"log"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_UserDB_Create_User(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")

	}
	UserDB := UserDB{DB: db}
	Cases := []struct {
		name     string
		userId   string
		want     User
		hasError bool
	}{
		{
			name:   "正常",
			userId: "aaa",
			want: User{
				UserId: "aaa",
			},

			hasError: false,
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := UserDB.CreateUser(tt.userId, 1)
			if (err != nil) != tt.hasError {
				t.Errorf("%v,got,%v,want%v", tt.name, err, tt.hasError)
			}
			User, err := UserDB.GetUser(tt.userId)
			if err != nil {
				t.Errorf("error")
			}
			log.Print(&User)
			if User.UserId != tt.want.UserId {
				t.Errorf("%v,got,%v,want%v", tt.name, User.UserId, tt.want.UserId)
			}
		})
	}
}
