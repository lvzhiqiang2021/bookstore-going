package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/pkg/utils"
	"html/template"
	"net/http"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request){
	flag, session := dao.IsLogin(r)

	if flag {
		bookID := r.FormValue("bookId")
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID)
		//如果购物车非空
		if cart != nil {
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				cts := cart.CartItems
				for _, v := range cts {
					//如果购物车已有同样的项目,则累加
					if v.Book.ID == cartItem.Book.ID {
						v.Count = v.Count + 1
						dao.UpdateBookCount(v)
					}
				}
			}else {
				cartItem := &model.CartItem{
					Book: book,
					Count: 1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				dao.AddCartItem(cartItem)
			}
			dao.UpdateCart(cart)
		} else {
			//创建一个购物车
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: int64(userID),
			}
			//购物车商品项目
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book: book,
				Count: 1,
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			dao.AddCart(cart)
		}
		w.Write([]byte("您添加了" + book.Title + " 到购物车!"))
	}else{
		w.Write([]byte("请先登陆!"))
	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request){
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		cart.UserName = session.UserName
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, cart)
	} else {
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

func DeleteCart(w http.ResponseWriter, r *http.Request){
	cartID := r.FormValue("cartId")
	dao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}