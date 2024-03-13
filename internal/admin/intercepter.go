package admin

import (
	"context"
	"log"

	"github.com/charisworks/charisworks-backend/validation"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

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
