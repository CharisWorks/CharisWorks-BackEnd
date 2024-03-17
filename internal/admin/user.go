package admin

import (
	"context"
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	userpb "github.com/charisworks/charisworks-backend/pkg/grpc"
	"github.com/charisworks/charisworks-backend/validation"
)

type UserServiceServer struct {
	userpb.UserServiceServer
}

func (r *UserServiceServer) All(ctx context.Context, req *userpb.VoidRequest) (res *userpb.AllUserResponse, err error) {
	db, err := utils.DBInit()
	res = new(userpb.AllUserResponse)
	if err != nil {
		return res, err
	}
	userIds := new([]string)
	userList := new([]users.User)
	db.Table("users").Where("1=1").Select("id").Find(&userIds)
	log.Print("userIds:", userIds)
	for _, userId := range *userIds {
		userRequests := users.Requests{UserUtils: users.UserUtils{}, UserRepository: users.UserRepository{DB: db}}
		user, err := userRequests.Get(userId)
		if err != nil {
			return res, err
		}
		*userList = append(*userList, user)
	}
	jsonBytes, err := json.Marshal(userList)
	if err != nil {
		return res, err
	}

	res.User = string(jsonBytes)

	return res, nil
}

func (r *UserServiceServer) Remove(ctx context.Context, req *userpb.RemoveUserRequest) (res *userpb.VoidResponse, err error) {
	res = new(userpb.VoidResponse)
	db, err := utils.DBInit()
	if err != nil {
		return res, err
	}
	userId := req.GetUser()
	userRequests := users.Requests{UserUtils: users.UserUtils{}, UserRepository: users.UserRepository{DB: db}}
	err = userRequests.Delete(userId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *UserServiceServer) Privilege(ctx context.Context, req *userpb.PrivilegeUserRequest) (res *userpb.VoidResponse, err error) {
	db, err := utils.DBInit()
	res = new(userpb.VoidResponse)
	if err != nil {
		return res, err
	}
	userId := req.GetUser()
	userRepository := users.UserRepository{DB: db}
	err = userRepository.UpdateProfile(userId, map[string]interface{}{"stripe_account_id": "acct_unregistered"})
	if err != nil {
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return res, err
	}
	app, err := validation.NewFirebaseApp()
	if err != nil {
		return res, err
	}

	SendPrivilegedEmail(userId, app)
	return res, nil
}

func (r *UserServiceServer) Transaction(ctx context.Context, req *userpb.SpecificUserTransactionRequest) (res *userpb.SpecificUserTransactionResponse, err error) {
	db, err := utils.DBInit()
	res = new(userpb.SpecificUserTransactionResponse)
	if err != nil {
		return res, err
	}
	userId := req.GetUser()
	transactionRepository := transaction.Repository{DB: db}
	transactions, err := transactionRepository.GetList(userId)
	if err != nil {
		return res, err
	}
	jsonBytes, err := json.Marshal(transactions)
	if err != nil {
		return res, err
	}
	res.Transaction = string(jsonBytes)
	return res, nil
}
