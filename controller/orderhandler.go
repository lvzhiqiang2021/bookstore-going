package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/pkg/utils"
	"html/template"
	"net/http"
	"time"
)
//去结账
func Checkout(w http.ResponseWriter, r *http.Request){
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	orderID := utils.CreateUUID()
	order := &model.Order{
		OrderID: orderID,
		CreateTime: timeStr,
		TotalCount: cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State: 0,
		UserID: int64(userID),
	}
	dao.AddOrder(order)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		orderItem := &model.OrderItem{
			Count: v.Count,
			Amount: v.Amount,
			Title: v.Book.Title,
			Author: v.Book.Author,
			Price: v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		dao.AddOrderItem(orderItem)
		book := v.Book
		book.Sales = book.Sales + v.Count
		book.Stock = book.Stock - v.Count
		dao.UpdateBook(book)
	}
	dao.DeleteCartByCartID(cart.CartID)
	session.Order = order
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}
func GetOrders(w http.ResponseWriter, r *http.Request){
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}
func GetOrderInfo(w http.ResponseWriter, r *http.Request){
	orderID := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func GetOrderByUserID(w http.ResponseWriter, r *http.Request){
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetOrderByUserID(userID)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

//发货
func SendOrder(w http.ResponseWriter, r *http.Request){
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID, 1)
	GetOrders(w, r)
}

//确认收货
func TakeOrder(w http.ResponseWriter, r *http.Request){
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID, 2)
	GetOrderByUserID(w, r)
}
