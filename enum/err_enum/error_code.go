package err_enum

type ErrorCode int

const (
	SUCCESS     = 200
	SYSTEM_ERR  = 400
	PARAM_ERR   = 401
	IS_LOCK_ERR = 402

	DATA_NOT_EXIST_CODE      = 501
	DATA_CAT_NOT_UPDATE_CODE = 502
	DATA_EXPIRED             = 503
	LOGIN_ERROR              = 504

	TOKEN_INVALID = 505
	TOKEN_EXPIRED = 506

	ORDER_ERROR                    = 601
	ORDER_STATE_ERROR              = 602
	ORDER_AMOUNT_ERROR             = 603
	PAY_TYPE_ERR                   = 604
	ORDER_MUST_ONLY_BARBER_ERROR   = 605
	ORDER_MUST_FINISHED_ERROR      = 607
	ORDER_PAY_YET_ERROR            = 608
	CASH_PAY_REMARK_IS_EMPTY_ERROR = 608

	// Barber error
	NOT_SET_MEAL_TIME           = 701
	PAUSE_REASON_EMPTY          = 703
	MEAL_TIME_SET_OUT           = 704
	MEAL_TIME_SET_ERROR         = 705
	NO_PERMISSION               = 706
	HAS_ORDER                   = 707
	NOT_IN_MEAL_TIME            = 708
	MEAL_TIME_YET               = 709
	NOT_IN_WORK                 = 710
	HAS_PAUSE_RECORD_NOT_HANDLE = 711

	//location
	USER_LOCATION_POSITION_EXISTS_CODE = 801
)

var (
	OK *ErrorObj = &ErrorObj{Code: SUCCESS, Msg: "success"}

	SystemErr  *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "系统异常"}
	RequestErr *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "请求异常"}
	ParamErr   *ErrorObj = &ErrorObj{Code: PARAM_ERR, Msg: "param error"}
	LockErr    *ErrorObj = &ErrorObj{Code: IS_LOCK_ERR, Msg: "lock error"}

	DataNotExistErr  *ErrorObj = &ErrorObj{Code: DATA_NOT_EXIST_CODE, Msg: "数据不存在"}
	DataCanNotUpdate *ErrorObj = &ErrorObj{Code: DATA_CAT_NOT_UPDATE_CODE, Msg: "数据不能修改"}
	DataExpired      *ErrorObj = &ErrorObj{Code: DATA_EXPIRED, Msg: "过期了"}
	LoginErr         *ErrorObj = &ErrorObj{Code: LOGIN_ERROR, Msg: "登录异常"}
	UserNotExistErr  *ErrorObj = &ErrorObj{Code: LOGIN_ERROR, Msg: "用户不存在"}

	// token
	TokenInvalid *ErrorObj = &ErrorObj{Code: TOKEN_INVALID, Msg: "token invalid"}
	TokenExpired *ErrorObj = &ErrorObj{Code: TOKEN_EXPIRED, Msg: "token expired"}

	// order
	OrderError              *ErrorObj = &ErrorObj{Code: ORDER_ERROR, Msg: "order error"}
	OrderPayYetError        *ErrorObj = &ErrorObj{Code: ORDER_PAY_YET_ERROR, Msg: "订单已经被抢了"}
	OrderCanNotCancleErr    *ErrorObj = &ErrorObj{Code: 601, Msg: "订单非排队状态不能取消"}
	OrderNotExistErr        *ErrorObj = &ErrorObj{Code: 602, Msg: "订单不存在"}
	OrderMustOnlyBarberErr  *ErrorObj = &ErrorObj{Code: ORDER_MUST_ONLY_BARBER_ERROR, Msg: "不能同时取两个发型师的号"}
	OrderMustFinishedErr    *ErrorObj = &ErrorObj{Code: ORDER_MUST_FINISHED_ERROR, Msg: "存在未支付的订单不能再取号"}
	OrderStateErr           *ErrorObj = &ErrorObj{Code: ORDER_STATE_ERROR, Msg: "订单状态异常"}
	OrderAmountErr          *ErrorObj = &ErrorObj{Code: ORDER_AMOUNT_ERROR, Msg: "订单金额异常"}
	PayTypeErr              *ErrorObj = &ErrorObj{Code: PAY_TYPE_ERR, Msg: "订单支付方式错误"}
	CashPayRemarkIsEmptyErr *ErrorObj = &ErrorObj{Code: CASH_PAY_REMARK_IS_EMPTY_ERROR, Msg: "现金支付时备注不能为空"}
	PayFrequentlyErr        *ErrorObj = &ErrorObj{Code: PAY_TYPE_ERR, Msg: "支付过于频繁，请稍等"}
	OrderNotSettled         *ErrorObj = &ErrorObj{Code: PAY_TYPE_ERR, Msg: "订单未结算"}
	CardBuyErr              *ErrorObj = &ErrorObj{Code: PAY_TYPE_ERR, Msg: "只能购买一张"}
	CardNotExistErr         *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "订单不存在次卡"}
	CardMoreThanOneErr      *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "订单次卡超过1张"}
	CardTimesIsNone         *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "次卡的次数已经用完"}

	//meituan
	MeituanAppShopIdNull       *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团绑定数据异常，需要后台管理录入"}
	MeituanFetchSessionFail    *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "根据AuthCode换取美团的Token失败"}
	MeituanFetchOpenUUIDFail   *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "获取美团Open UUID失败"}
	MeituanVerifyCouponFail    *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团验券失败"}
	MeituanTokenExpires        *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团Token过期，无法刷新，需要客户端重新授权"}
	MeituanFetchCouponInfoFail *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团获取券信息失败"}
	MeituanCouponBinded        *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团券已经被使用并绑定了其他订单"}
	MeituanCouponErrState      *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "美团券状态异常"}

	// coupon
	CouponReceivedErr         *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "优惠券已经领取"}
	CouponReceivedNotNewErr   *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "不是新人"}
	CouponReceivedNotPhoneErr *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "没有手机号"}

	// user
	UserBirthdayIsExistErr *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "生日不能重新修改"}
	AccountNotEnough       *ErrorObj = &ErrorObj{Code: SYSTEM_ERR, Msg: "账户余额不足"}
	AuthorizationFailedErr *ErrorObj = &ErrorObj{Code: 601, Msg: "授权失效"}

	UserLocationPositionExistsErr *ErrorObj = &ErrorObj{Code: USER_LOCATION_POSITION_EXISTS_CODE, Msg: "位置信息已存在"}
)

type ErrorObj struct {
	Code ErrorCode
	Msg  string
}

func (err ErrorObj) Error() string {
	return err.Msg
}

func NewErrorObj(err error) *ErrorObj {
	return &ErrorObj{
		Code: SYSTEM_ERR,
		Msg:  err.Error(),
	}
}

func NewWithMsg(msg string) *ErrorObj {
	return &ErrorObj{
		Code: SYSTEM_ERR,
		Msg:  msg,
	}
}
