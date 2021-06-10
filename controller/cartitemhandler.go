package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func DeleteCartItem(w http.ResponseWriter, r *http.Request){
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}
func UpdateCartItem(w http.ResponseWriter, r *http.Request){
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems

	for _, v := range cartItems {
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)

	//刷新购物车
	cart, _ = dao.GetCartByUserID(userID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount

	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			amount = v.Amount
		}
	}
	data := model.Data{
		Amount: amount,
		TotalAmount: totalAmount,
		TotalCount: int64(totalCount),
	}
	json, _ := json.Marshal(data)
	w.Write(json)
}