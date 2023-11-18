package common

type IError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type FieldErrorHandlerResp struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Error   []*IError `json:"error"`
}

type SuccessHandlerResp struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}
