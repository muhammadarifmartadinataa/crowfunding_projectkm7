package helper

type Response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonresponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonresponse
}
