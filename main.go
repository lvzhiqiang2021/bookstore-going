package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//用户处理
	http.HandleFunc("/", controller.GetPageBooksByPrice)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	//购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)         // 添加图书到购物车
	http.HandleFunc("/getCartInfo", controller.GetCartInfo) 		  // 获取购物车信息
	http.HandleFunc("/deleteCart", controller.DeleteCart)			  // 清空购物车
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)     // 删除购物车项目
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)     // 更新购物车项目

	//订单处理
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getOrderByUserID", controller.GetOrderByUserID)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":8080", nil)

}
