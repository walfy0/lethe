package common

type Resp struct{
	Code 	string 		`json:"code"`
	Message string 		`json:"message"`
	Data 	interface{} `json:"data"`
}

func SuccessResp(data interface{}) Resp {
	return Resp{
		Code: "0",
		Message: "success",
		Data: data,
	}
}

func ErrorResp(err ErrorCode, data interface{}) Resp {
	return Resp{
		Code: err.Code,
		Message: err.Message,
		Data: data,
	}
}