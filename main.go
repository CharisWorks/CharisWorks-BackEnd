package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/charisworks/charisworks-backend/handler"
	"github.com/charisworks/charisworks-backend/internal/admin"
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/images"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	userpb "github.com/charisworks/charisworks-backend/pkg/grpc"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	R2Conns := images.R2Conns{Crud: nil}
	R2Conns.Init()
	utils.CORS(r)
	h := handler.NewHandler(r)
	app, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}
	db, err := utils.DBInit()
	if err != nil {
		return
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		return
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(admin.AuthUnaryServerInterceptor),
	)
	userpb.RegisterUserServiceServer(s, &admin.UserServiceServer{})
	userpb.RegisterItemServiceServer(s, &admin.ItemServiceServer{})
	userpb.RegisterTransactionServiceServer(s, &admin.TransactionServiceServer{})

	go func() {
		port := 50051
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}

		reflection.Register(s)
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	go func() {
		cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
		itemRequests := items.Requests{ItemRepository: items.ItemRepository{DB: db}, ItemUtils: items.ItemUtils{}}
		userRequests := users.Requests{UserUtils: users.UserUtils{}, UserRepository: users.UserRepository{DB: db}}
		manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db, Crud: R2Conns}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
		stripeRequests := cash.Requests{CartRequests: cartRequests, UserRequests: userRequests}
		transactionRequests := transaction.TransactionRequests{TransactionRepository: transaction.Repository{DB: trdb, UserRepository: users.UserRepository{DB: db}}, CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, StripeRequests: cash.Requests{CartRequests: cartRequests, UserRequests: userRequests}, StripeUtils: cash.Utils{}}
		webhookRequests := transaction.Webhook{StripeUtils: cash.Utils{}, TransactionRepository: transaction.Repository{DB: trdb}, ItemUpdater: items.Updater{DB: db}, App: app.App}
		h.SetupRoutesForWebhook(webhookRequests, app)
		h.SetupRoutesForItem(itemRequests)
		h.SetupRoutesForUser(app, userRequests)
		h.SetupRoutesForCart(app, cartRequests, userRequests)
		h.SetupRoutesForManufacturer(app, manufacturerRequests, userRequests)
		h.SetupRoutesForStripe(app, userRequests, stripeRequests, transactionRequests)
		h.SetupRoutesForImages(app, manufacturerRequests, itemRequests, userRequests, R2Conns)
		h.Router.Run(":8080")
	}()
	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
