package transaction

type TransactionRequests struct {
	TransactionRepository ITransactionHistoryRepository
}

func (r TransactionRequests) GetList(userId string) (*[]TransactionPreview, error) {
	transactionPreview, err := r.TransactionRepository.GetList(userId)
	if err != nil {
		return nil, err
	}
	return transactionPreview, nil
}

func (r TransactionRequests) GetDetails(userId string, transactionId string) (*TransactionDetails, error) {
	transactionDetails, transactionUserId, err := r.TransactionRepository.GetDetails(transactionId)
	if err != nil {
		return nil, err
	}
	if transactionUserId != userId {
		return nil, nil
	}
	return transactionDetails, nil
}
