package validation

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type IFirebaseApp interface {
	VerifyIDToken(*gin.Context) (string, error)
}

type FirebaseApp struct {
	App *firebase.App
}

// initialize app with ServiceAccountKey.json
func NewFirebaseApp() (*FirebaseApp, error) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return nil, err
	}

	return &FirebaseApp{app}, nil
}

func (app *FirebaseApp) VerifyIDToken(ctx *gin.Context) (string, error) {
	client, err := app.App.Auth(ctx.Request.Context())
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return "", err
	}
	idToken := ctx.Request.Header.Get("Authorization")
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return "", err
	}
	ctx.Set("UserEmail", token.Claims["email"].(string))
	log.Printf("Verified ID token: %v\n", token)

	return token.UID, nil
}
