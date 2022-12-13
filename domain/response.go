package domain

type BaseResponse struct {
	Data      Pokemon   `json:"data,omitempty"`
	DataArray []Pokemon `json:"datarray,omitempty"`
	Error     string    `json:"error,omitempty"`
}

func (b *BaseResponse) SetErrorMessage(message string) {
	b.Error = message
}
