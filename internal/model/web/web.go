package web

type WebResponse struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OkResponse(msg string, data interface{}) *WebResponse {
	return &WebResponse{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func CreatedResponse(msg string, data interface{}) *WebResponse {
	return &WebResponse{
		Code:    201,
		Message: msg,
		Data:    data,
	}
}

func BadRequestResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    400,
		Message: msg,
		Data:    err.Error(),
	}
}

func InternalServerErrorResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    500,
		Message: msg,
		Data:    err.Error(),
	}
}

func ForbiddenResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    403,
		Message: msg,
		Data:    err.Error(),
	}
}

func NotFoundResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    404,
		Message: msg,
		Data:    err.Error(),
	}
}

func UnauthorizedResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    401,
		Message: msg,
		Data:    err.Error(),
	}
}

func ConflictResponse(msg string, err error) *WebResponse {
	return &WebResponse{
		Code:    409,
		Message: msg,
		Data:    err.Error(),
	}
}
