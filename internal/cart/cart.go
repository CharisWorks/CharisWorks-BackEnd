package cart

func GetCart(i ICartRequest) ([]Cart, error) {
	Cart, err := i.Get()
	return *Cart, err
}
func PostCart(c CartRequestPayload, i ICartRequest) error {
	err := i.Register(c)
	return err
}
func UpdateCart(c CartRequestPayload, i ICartRequest) error {
	err := i.Update(c)
	return err
}
func DeleteCart(itemId string, i ICartRequest) error {
	err := i.Delete(itemId)
	return err
}
