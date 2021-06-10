package model

type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物项目
	TotalCount  int64       //总数量
	TotalAmount float64     //总金额
	UserID      int64         //所属用户
	UserName    string
}

//获取图书的总数
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//获取图书总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
