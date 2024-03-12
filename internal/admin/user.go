package admin

import (
	"context"
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	userpb "github.com/charisworks/charisworks-backend/pkg/grpc"
	"github.com/charisworks/charisworks-backend/validation"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

type GetAllUserServiceServer struct {
	userpb.UnimplementedGetAllUserServiceServer
}

func (r *GetAllUserServiceServer) GetAll(ctx context.Context, req *userpb.AllUserRequest) (res *userpb.AllUserResponse, err error) {
	db, err := utils.DBInit()
	if err != nil {
		return nil, err
	}
	userIds := new([]string)
	userList := new([]users.User)
	db.Table("users").Where("1=1").Select("id").Find(&userIds)
	for _, userId := range *userIds {
		userRequests := users.Requests{UserUtils: users.UserUtils{}, UserRepository: users.UserRepository{DB: db}}
		user, err := userRequests.Get(userId)
		if err != nil {
			return nil, err
		}
		*userList = append(*userList, *user)
	}
	jsonBytes, err := json.Marshal(userList)
	if err != nil {
		return nil, err
	}
	res = new(userpb.AllUserResponse)
	res.Message = "ok"
	res.User = string(jsonBytes)

	return res, nil
}

func AuthUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptor 1: ", info.FullMethod) // ハンドラの前に割り込ませる前処理
	idToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	app, err := validation.NewFirebaseApp()
	if err != nil {
		return nil, err
	}
	fApp := validation.FirebaseApp{App: app.App}
	userId, _, _, err := fApp.Verify(ctx, idToken)
	if err != nil {
		return nil, err
	}
	if userId != "cowatanabe26@gmail.com" {
		return nil, err
	}
	log.Printf("idToken: %v\n", idToken)
	res, err := handler(ctx, req)                         // 本来の処理
	log.Println("[post] my unary server interceptor 1: ") // ハンドラの後に割り込ませる後処理
	return res, err
}
