package validation

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type IFirebaseApp interface {
	VerifyIDToken(*gin.Context) (string, error)
	Verify(context.Context, string) (string, string, bool, error)
	GetEmail(string) (string, error)
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
	userId, email, emailVerified, err := app.Verify(ctx.Request.Context(), strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer "))
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		ctx.Abort()
		return "", err
	}

	ctx.Set("UserEmail", email)
	ctx.Set("UserId", userId)
	ctx.Set("EmailVerified", emailVerified)

	return userId, nil
}
func (app *FirebaseApp) Verify(ctx context.Context, idToken string) (userId string, email string, emailVerified bool, err error) {
	client, err := app.App.Auth(ctx)
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return userId, email, emailVerified, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return userId, email, emailVerified, err
	}
	userId = token.UID
	email = token.Claims["email"].(string)
	emailVerified = token.Claims["email_verified"].(bool)

	log.Printf("Verified ID token: %v\n", token)
	return userId, email, emailVerified, nil
}
func (app *FirebaseApp) GetEmail(userId string) (string, error) {
	client, err := app.App.Auth(context.Background())
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return "", err
	}

	user, err := client.GetUser(context.Background(), userId)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return "", err
	}

	return user.Email, nil
}
