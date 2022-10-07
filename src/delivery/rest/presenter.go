package rest

// Ok generates success response payload.
func Ok(config *Config, data interface{}) (int, *Response) {
	return config.DefaultOkStatus, &Response{
		Code:    config.DefaultOkCode,
		Message: config.DefaultOkMessage,
		Data:    data,
	}
}

// NotOk generates error response payload.
func NotOk(config *Config, status int, code, message string) (int, *Response) {
	if status != 0 {
		return status, &Response{
			Code:    code,
			Message: message,
		}
	} else {
		return config.DefaultNotOkStatus, &Response{
			Code:    code,
			Message: message,
		}
	}
}
