package handlers

const (
	Error     = "Error"
	MessageOK = "OK"
)

type Response struct {
	MessageType string      `json:"messageType"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(ms string, m string, data interface{}) Response {
	return Response{
		MessageType: ms,
		Message:     m,
		Data:        data,
	}
}
