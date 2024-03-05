package cart

/*
	func ExampleCart() *[]Cart {
		e := Cart{
			ItemId:   "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
			Quantity: 1,
			ItemProperties: CartItemPreviewProperties{
				Name:  "クラウディ・エンチャント",
				Price: 2000,
				Details: CartItemPreviewDetails{
					Status: CartItemStatusAvailable,
				},
			},
		}
		re := new([]Cart)
		cart := append(*re, e)
		return &cart

}

type ExapleCartRequest struct {
}

	func (p ExapleCartRequest) Get(ctx *gin.Context, i IcartRepository, UserId string) (*[]Cart, error) {
		Cart := ExampleCart()
		return Cart, nil
	}

	func (c ExapleCartRequest) Register(p CartRequestPayload, i IcartRepository, ctx *gin.Context) error {
		log.Print("CartRequestPayload: ", p)
		if p.Quantity <= 0 {
			return &utils.InternalError{Message: utils.InternalErrorInvalidQuantity}
		}
		return nil
	}

	func (c ExapleCartRequest) Delete(itemId string, ctx *gin.Context) error {
		log.Print("itemId: ", itemId)
		return nil
	}

	func (c ExapleCartRequest) GetItem(itemid string) (*itemStatus, error) {
		return nil, nil
	}
*/
/* type ExamplecartRepository struct {
	ItemStatus      *itemStatus
	InternalCarts   *[]InternalCart
	ItemSelectError error
	SelectError     error
	UpdateError     error
	RegisterError   error
	DeleteError     error
}

func (c ExamplecartRepository) GetItem(itemId string) (*itemStatus, error) {
	if c.ItemSelectError != nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	return c.ItemStatus, nil
}
func (c ExamplecartRepository) GetCart(UserId string) (*[]InternalCart, error) {
	if c.SelectError != nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorMessage(c.SelectError.Error())}
	}
	return c.InternalCarts, nil
}
func (c ExamplecartRepository) RegisterCart(UserId string, CartRequestPayload CartRequestPayload) error {
	if c.RegisterError != nil {
		return c.RegisterError
	}
	return nil
}
func (c ExamplecartRepository) UpdateCart(UserId string, CartRequestPayload CartRequestPayload) error {
	if c.UpdateError != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (c ExamplecartRepository) DeleteCart(UserId string, itemId string) error {
	if c.DeleteError != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
*/
