package cart

type CartRequest struct {
}

func (c CartRequest) Get() *[]Cart {
	Cart := ExampleCart()
	return &Cart
}
func (c CartRequest) Register(p CartRequestPayload) string {
	return ""
}
func (c CartRequest) Update(p CartRequestPayload) string {
	return ""
}
func (c CartRequest) Delete(ItemId string) string {
	return ""
}

func GetCart(i ICartRequest) []Cart {
	Cart := i.Get()
	return *Cart
}
func PostCart(c CartRequestPayload, i ICartRequest) []Cart {
	Cart := i.Get()
	return *Cart
}
func UpdateCart(c CartRequestPayload, i ICartRequest) []Cart {
	Cart := i.Get()
	return *Cart
}
func DeleteCart(Item_id string, i ICartRequest) []Cart {
	Cart := i.Get()
	return *Cart
}
