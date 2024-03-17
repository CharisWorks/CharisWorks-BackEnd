package admin

import (
	"context"
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"

	"github.com/charisworks/charisworks-backend/internal/utils"
	transactionpb "github.com/charisworks/charisworks-backend/pkg/grpc"
)

type TransactionServiceServer struct {
	transactionpb.UnimplementedTransactionServiceServer
}

func (r *TransactionServiceServer) All(ctx context.Context, rep *transactionpb.VoidRequest) (res *transactionpb.AllTransactionResponse, err error) {
	db, err := utils.DBInit()
	res = new(transactionpb.AllTransactionResponse)
	if err != nil {
		return res, err
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		return res, err
	}
	users := new([]string)
	transactionList := new([]transaction.TransactionPreview)
	db.Table("users").Where("1=1").Select("id").Find(&users)
	for _, user := range *users {
		transactionRepository := transaction.Repository{DB: trdb}
		transaction, err := transactionRepository.GetList(user)
		if err != nil {
			return res, err
		}
		for _, t := range transaction {
			*transactionList = append(*transactionList, t)
		}

	}
	bytes, err := json.Marshal(transactionList)
	if err != nil {
		return res, err
	}
	res.Transaction = string(bytes)
	return res, nil

}

func (r *TransactionServiceServer) ById(ctx context.Context, req *transactionpb.SpecificTransactionRequest) (res *transactionpb.SpecificTransactionResponse, err error) {

	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		return res, err
	}
	res = new(transactionpb.SpecificTransactionResponse)
	transactionRepository := transaction.Repository{DB: trdb}
	transaction, _, _, err := transactionRepository.GetDetails(req.GetTransaction())
	if err != nil {
		return res, err
	}
	bytes, err := json.Marshal(transaction)
	if err != nil {
		return res, err
	}
	res.Transaction = string(bytes)
	return res, nil
}

func (r *TransactionServiceServer) RegisterTrackingId(ctx context.Context, req *transactionpb.RegisterTrackingIdRequest) (res *transactionpb.VoidResponse, err error) {
	res = new(transactionpb.VoidResponse)
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		log.Print(err)
		return res, err
	}
	if err := trdb.Table("transactions").Where("transaction_id = ?", req.GetTransaction()).Updates(map[string]interface{}{"tracking_id": req.GetTrackingId()}).Error; err != nil {
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return res, err
	}
	db, err := utils.DBInitTest()
	if err != nil {
		return res, err
	}
	transactionDetails, _, _, err := transaction.Repository{DB: trdb, UserRepository: users.UserRepository{DB: db}}.GetDetails(req.GetTransaction())
	if err != nil {
		return res, err
	}
	SendShippedEmail(transactionDetails)
	return res, nil
}
func (r *TransactionServiceServer) RegisterStatus(ctx context.Context, req *transactionpb.UpdateTransactionStatusRequest) (res *transactionpb.VoidResponse, err error) {
	res = new(transactionpb.VoidResponse)
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		return res, err
	}
	if err := trdb.Table("transactions").Where("transaction_id = ?", req.GetTransaction()).Updates(map[string]interface{}{"status": req.GetStatus()}).Error; err != nil {
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return res, err
	}
	return res, nil
}
