package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "系统繁忙, 请稍后重试")
	InvalidParams             = NewError(10001, "无效输入")
	NotFound                  = NewError(10002, "未找到数据")
	UnauthorizedAuthNotExist  = NewError(10003, "账户不存在")
	UnauthorizedAuthFailed    = NewError(10004, "账户密码错误")
	UnauthorizedTokenError    = NewError(10005, "未登录")
	UnauthorizedTokenTimeout  = NewError(10006, "登录信息已过期")
	UnauthorizedTokenGenerate = NewError(10007, "Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")
	InvalidLink               = NewError(10009, "链接失效")
	OutDateWeiXinAccessToken  = NewError(10010, "微信token超时")
	FunctionInService         = NewError(10011, "功能维护中,请先试用替代功能")

	GatewayMethodsLimit    = NewError(10109, "网关仅接受GET或POST请求")
	GatewayLostSign        = NewError(10110, "网关请求缺少签名")
	GatewayLostAppKey      = NewError(10111, "网关请求缺少APP KEY")
	GatewayAppKeyInvalid   = NewError(10112, "网关请求无效APP KEY")
	GatewayAppKeyClosed    = NewError(10113, "网关请求APP KEY已停用")
	GatewayParamSignError  = NewError(10114, "网关请求参数签名错误")
	GatewayTooManyRequests = NewError(10115, "网关请求频次超限")

	FileUploadFailed = NewError(10200, "文件上传失败")
	FileInvalidExt   = NewError(10201, "文件类型不合法")
	FileInvalidSize  = NewError(10202, "文件大小超限")

	UsernameHasExisted           = NewError(20001, "用户名已存在")
	UsernameLengthLimit          = NewError(20002, "用户名长度3~12")
	UsernameCharLimit            = NewError(20003, "用户名只能包含字母、数字")
	PasswordLengthLimit          = NewError(20004, "密码长度6~16")
	UserRegisterFailed           = NewError(20005, "用户注册失败")
	UserHasBeenBanned            = NewError(20006, "该账户已被封停")
	NoPermission                 = NewError(20007, "无权限请求")
	UserHasBindOTP               = NewError(20008, "当前用户已绑定二次验证")
	UserOTPInvalid               = NewError(20009, "二次验证码验证失败")
	UserNoBindOTP                = NewError(20010, "当前用户未绑定二次验证")
	ErrorOldPassword             = NewError(20011, "当前用户密码验证失败")
	ErrorCaptchaPassword         = NewError(20012, "图形验证码验证失败")
	AccountNoPhoneBind           = NewError(20013, "拒绝操作: 账户未绑定手机号")
	TooManyLoginError            = NewError(20014, "登录失败次数过多，请稍后再试")
	GetPhoneCaptchaError         = NewError(20015, "短信验证码获取失败")
	TooManyPhoneCaptchaSend      = NewError(20016, "短信验证码获取次数已达今日上限")
	ExistedUserPhone             = NewError(20017, "该手机号已被绑定")
	ErrorPhoneCaptcha            = NewError(20018, "手机验证码不正确")
	PhoneCaptchaExpired          = NewError(20019, "手机验证码已过期")
	NicknameLengthLimit          = NewError(20020, "昵称长度2~12")
	NoExistUsername              = NewError(20021, "用户不存在")
	NoAdminPermission            = NewError(20022, "无管理权限")
	NoUserJoinedGroups           = NewError(20023, "无法查询到用户已加入的群组")
	UpdateGroupInfoFailed        = NewError(20024, "无法更新当前群组信息")
	GetHotSearchFailed           = NewError(20025, "获取热门搜索失败")
	DeleteSearchHistoryFailed    = NewError(20026, "删除搜索历史失败")
	GetSuggestSearchFailed       = NewError(20027, "获取推荐列表失败")
	PlaceUserGoodsRequestFailed  = NewError(20028, "用户发送商品请求失败")
	PlaceUserBidFailed           = NewError(20029, "用户发送商品出价失败")
	UpdateGoodsProviderFailed    = NewError(20030, "更新商品信息失败")
	DeleteGoodsProviderFailed    = NewError(20031, "删除商品信息失败")
	CreateZhaiHubFailed          = NewError(20032, "创建圈失败")
	ValidateZhaiHubFailed        = NewError(20033, "校验圈失败")
	UserNoGenderInfo             = NewError(20034, "用户未设置性别")
	NotEnoughCraftPoints         = NewError(20035, "圈块点不足")
	NotAllowDuplicateAccept      = NewError(20036, "你已经参与了, 不能再参与喽")
	QrCodeExpired                = NewError(20037, "二维码过期")
	NoEndTime                    = NewError(20038, "没有填写结束时间")
	MissionOrderToPay            = NewError(20039, "存在未支付的订单，请前往我的订单支付")
	NotVerifiedUser              = NewError(20040, "请先完成用户认证")
	NoRepeatUserApply            = NewError(20041, "不能重复提交申请")
	NotEnoughBalance             = NewError(20042, "余额不足")
	PleaseContactCustomerService = NewError(20043, "请联系客服解决")
	ExceedEndTime                = NewError(20044, "已经超出结束时间, 无法取消")
	UnknownLink                  = NewError(20045, "未知的链接类型")
	NotAllowedDuplicate          = NewError(20046, "不允许重复创建")
	PendingApprove               = NewError(20047, "审核中,请等待审核完成后再查看")
	AlreadyApprove               = NewError(20048, "已经通过申请,无法重复通过")
	NoZhaihubPermission          = NewError(20049, "你还没有加入圈哦")
	NoExistPost                  = NewError(20050, "不存在的动态")
	NoDuplicateRedeem            = NewError(20051, "不允许重复获取")
	NotValidCDK                  = NewError(20052, "无效的cdk")
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedAuthFailed.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
