package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
)

func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cart_items(count, amount, book_id, cart_id) values(?,?,?,?)"
	_, err := db.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sql := "select id, count, amount, cart_id from cart_items where book_id = ? and cart_id = ?"
	row := db.Db.QueryRow(sql, bookID, cartID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	book, err := GetBookByID(bookID)
	if err != nil {
		return nil, err
	}
	cartItem.Book = book
	return cartItem, nil
}

func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sql := "select id, count, amount, book_id, cart_id from cart_items where cart_id = ?"
	rows, err := db.Db.Query(sql, cartID)
	if err != nil {
		return nil, err
	}

	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		book, _ := GetBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

func UpdateBookCount(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ?, amount = ? where book_id = ? and cart_id = ?"
	_, err := db.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := db.Db.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByID(cartItemID string) error {
	sql := "delete from cart_items where id = ?"
	_, err := db.Db.Exec(sql, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
