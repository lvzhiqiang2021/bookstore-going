package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
)

func AddCart(cart *model.Cart) error {
	sql := "insert into carts(id, total_count, total_amount, user_id) values(?,?,?,?)"
	_, err := db.Db.Exec(sql, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		AddCartItem(cartItem)
	}
	return nil
}

func GetCartByUserID(userID int) (*model.Cart, error) {
	sql := "select id, total_count, total_amount, user_id from carts where user_id = ?"
	row := db.Db.QueryRow(sql, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems

	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_count = ?, total_amount = ? where id = ?"
	_, err := db.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//清空购物车内容
func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		return err
	}
	sql := "delete from carts where id = ?"
	_, err = db.Db.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}
