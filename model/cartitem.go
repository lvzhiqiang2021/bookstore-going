package model

type CartItem struct {
	CartItemID int64  //购物项ID
	Book       *Book   //购物车图书信息
	Count      int64   //购物车商品数量
	Amount     float64 //金额小计
	CartID     string  //购物ID
}

func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
