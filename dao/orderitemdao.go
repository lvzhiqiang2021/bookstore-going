package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items(count, amount, title, author, price, img_path, order_id) values(?,?,?,?,?,?,?)"
	_, err := db.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sql := "select id, count, amount, title, author, price, img_path, order_id from order_items where order_id=?"
	rows, err := db.Db.Query(sql, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
