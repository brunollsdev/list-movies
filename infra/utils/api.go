package utils

type ResponseAPI struct {
	Status bool            `json:"status"`
	Data   DataResponseAPI `json:"data"`
}

type DataResponseAPI struct {
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

func Response(status bool, message string, response interface{}) ResponseAPI {
	return ResponseAPI{
		Status: status,
		Data: DataResponseAPI{
			Message:  message,
			Response: response,
		},
	}
}
