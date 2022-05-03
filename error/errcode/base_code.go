package errcode

var (
	OK                         int32 = 200  //请求成功
	UNDEFINE                   int32 = 2    //错误未定义
	BAD_REQUEST                int32 = 400  //参数错误:当前请求无法被服务器理解
	TOKEN_INVALID              int32 = 401  //鉴权错误:无法获取Token/AppKey
	SERVICE_FORBIDDEN          int32 = 403  //鉴权错误:服务拒绝处理请求
	INTERNAL_SERVER_ERROR      int32 = 500  //服务异常:内部服务器错误
	PARAMETER_VALIDATE_ERROR   int32 = 1010 //参数错误:参数验证错误
	BUSINESS_LOGIC_ERROR       int32 = 1014 //服务异常:因业务逻辑错误导致请求不能被正常处理
	EXTERNAL_DEPENDENCY_REJECT int32 = 1015 //服务异常:依赖其他项目错误
	EXTERNAL_DEPENDENCY_ERROR  int32 = 1016 //服务异常:依赖服务返回拒绝
)

var errorMsg = map[int32]string{
	OK:                         "请求成功",
	UNDEFINE:                   "错误未定义",
	BAD_REQUEST:                "参数错误:当前请求无法被服务器理解",
	TOKEN_INVALID:              "鉴权错误:无法获取Token/AppKey",
	SERVICE_FORBIDDEN:          "鉴权错误:服务拒绝处理请求",
	INTERNAL_SERVER_ERROR:      "服务异常:内部服务器错误",
	PARAMETER_VALIDATE_ERROR:   "参数错误:参数验证错误",
	BUSINESS_LOGIC_ERROR:       "服务异常:因业务逻辑错误导致请求不能被正常处理",
	EXTERNAL_DEPENDENCY_REJECT: "服务异常:依赖其他项目错误",
	EXTERNAL_DEPENDENCY_ERROR:  "服务异常:依赖服务返回拒绝",
}

func GetErrMsg(key int32) string {
	return errorMsg[key]
}
