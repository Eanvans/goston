package model

// 支付类型
type PayType int

const (
	PAY_TYPE_WALLET  PayType = iota //钱包
	PAY_TYPE_WECHAT                 //微信支付
	PAY_TYPE_ALIPAY                 //支付宝支付
	PAY_TYPE_OFFLINE                //线下支付
)

// 支付状态
type PayStatus int

const (
	PAY_STATUS_PENDING PayStatus = iota //待付款
	PAY_STATUS_PAID                     //已付款
	PAY_STATUS_CANCEL                   //已取消
	PAY_STATUS_REFUND                   //已退款
)

// 订单状态
type OrderStatus int

const (
	ORDER_STATUS_ALL              OrderStatus = -1 //获取全部
	ORDER_STATUS_CANCELLED        OrderStatus = 0  //已取消
	ORDER_STATUS_PENDING_PAY      OrderStatus = 10 //待付款
	ORDER_STATUS_PENDING_DELIVERY OrderStatus = 20 //待发货
	ORDER_STATUS_PENDING_CONFIRM  OrderStatus = 30 //待收货、待使用
	ORDER_STATUS_DONE             OrderStatus = 40 //已完成、已使用
	ORDER_STATUS_REFUND           OrderStatus = 50 //已退款
)

type Order struct {
	ID       int64  `json:"id"`
	OrderNo  string `json:"order_no"`
	UserID   int64  `json:"user_id"`
	UserMemo string `json:"user_memo"` //

	Amount     int       `json:"amount"`      //购买数量
	TotalPrice int64     `json:"total_price"` //商品总金额
	PayPrice   int64     `json:"pay_price"`   //实际付款金额(包含运费)
	PayType    PayType   `json:"pay_type"`    //支付方式: 0: 钱包? 1:微信 2:支付宝...
	PayStatus  PayStatus `json:"pay_status"`  //付款状态 0: 未付款; 1: 已付款;
	PayTime    int64     `json:"pay_time"`    //

	OrderStatus OrderStatus `json:"order_status"` //订单状态 10待付款 20待发货 30待收货 40已完成
}
