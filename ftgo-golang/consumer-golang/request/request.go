package request

type CreateConsumerRequest struct {
	Name    string `json:"name"  binding:"required"`
	Address string `json:"address" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

type UpdateConsumerRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
}
