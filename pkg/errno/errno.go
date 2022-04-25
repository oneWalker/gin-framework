package errno

// codes
var (
	OK            = &Errno{Code: 0, Message: "OK"}
	FAILED        = &Errno{Code: -1, Message: "FAILED"}
	ParamsInvalid = &Errno{Code: -2, Message: "参数错误"}
	TokenMissing  = &Errno{Code: -3, Message: "token缺失"}
	TokenExpired  = &Errno{Code: -4, Message: "登录失效,请重新登录"}

	DataIllegal  = &Errno{Code: -5, Message: "数据查询失败，请联系管理员处理"}
)
