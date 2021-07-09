package common

type ErrorCode struct{
	Code 	string
	Message string
}

var (
	Success 	 		= ErrorCode{Code: "0", Message: "success"}
	ParamsError  		= ErrorCode{Code: "1", Message: "params error"}
	ServiceError 		= ErrorCode{Code: "2", Message: "service error"}
	RegisterTimeout 	= ErrorCode{Code: "3", Message: "register timeout"}
	RegisterCodeError 	= ErrorCode{Code: "4", Message: "register code error"}
)
