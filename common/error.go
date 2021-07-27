package common

type ErrorCode struct{
	Code 	string
	Message string
}

var (
	Success 	 		= ErrorCode{Code: "0", Message: "success"}
	ParamsError  		= ErrorCode{Code: "1", Message: "params error"}
	DataBaseError  		= ErrorCode{Code: "2", Message: "database error"}
	ServiceError 		= ErrorCode{Code: "3", Message: "service error"}
	RegisterTimeout 	= ErrorCode{Code: "4", Message: "register timeout"}
	RegisterCodeError 	= ErrorCode{Code: "5", Message: "register code error"}
	OrderNotEnoughError = ErrorCode{Code: "6", Message: "order num not enough"}
	PasswordError		= ErrorCode{Code: "7", Message: "password error"}
)
