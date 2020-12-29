package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(100000, "服务器内部错误")
	InvalidParams             = NewError(100001, "入参错误")
	NotFound                  = NewError(100002, "找不到")
	UnauthorizedAuthNotExist  = NewError(100003, "鉴权失败，找不到对应的APP key和APP secret")
	UnauthorizedTokenError    = NewError(100004, "鉴权失败，Token错误")
	UnauthorizedTokenExpire   = NewError(100005, "鉴权失败，Token过期")
	UnauthorizedTokenGenerate = NewError(100006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(100007, "请求过多")
)
