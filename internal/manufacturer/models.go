package manufacturer

type RegisterPayload struct {
	Name    string                     `json:"name" binding:"required"`
	Price   int                        `json:"price" binding:"required"`
	Details ItemRegisterDetailsPayload `json:"details" binding:"required"`
}
type ItemRegisterDetailsPayload struct {
	Stock       int      `json:"stock" binding:"required"`
	Size        int      `json:"size" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Tags        []string `json:"tags" binding:"required" `
}

type UpdatePayload struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Status      string   `json:"status"`
	Stock       int      `json:"stock"`
	Size        int      `json:"size"`
	Description string   `json:"description" `
	Tags        []string `json:"tags"`
}

type IItemRequests interface {
	Register(registerPayload RegisterPayload, userId string, itemId string) error
	Update(updatePayload UpdatePayload, userId string, itemId string) error
	Delete(itemId string, userId string) error
}

type IInspectPayloadUtils interface {
	Register(RegisterPayload) error
	Update(UpdatePayload) (map[string]interface{}, error)
}

type IItemRepository interface {
	Register(itemId string, i RegisterPayload, userId string) error
	Update(i map[string]interface{}, itemId string) error
	Delete(itemId string) error
}
