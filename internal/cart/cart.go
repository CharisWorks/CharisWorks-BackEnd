package cart

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
