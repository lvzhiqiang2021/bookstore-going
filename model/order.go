package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int64
	UserID      int64
}

//订单状态, 1.未提交、2. 提交、3. 完成
func (order *Order) NoSend() bool {
	return order.State == 0
}

func (order *Order) SendComplete() bool {
	return order.State == 1
}

func (order *Order) Complete() bool {
	return order.State == 2
}
