package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request){
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}
func DeleteBook(w http.ResponseWriter, r *http.Request){
	bookID := r.FormValue("bookId")
	dao.DeleteBook(bookID)
	GetPageBooks(w, r)
}
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request){
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	//如果图书ID能找到,则将数据渲染到更新页面,否则渲染新增
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	}else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request){
	bookId := r.PostFormValue("bookId")
	title  := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price  := r.PostFormValue("price")
	sales  := r.PostFormValue("sales")
	stock  := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales,10, 0)
	iStock, _ := strconv.ParseInt(stock,10,0)
	iBookID, _ := strconv.ParseInt(bookId, 10, 0)
	book := &model.Book{
		ID: int(iBookID),
		Title: title,
		Author: author,
		Price: fPrice,
		Sales: iSales,
		Stock: iStock,
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		dao.UpdateBook(book)
	}else {
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}
func GetPageBooks(w http.ResponseWriter, r *http.Request){
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request){
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == ""{
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}
func IndexHandler(w http.ResponseWriter, r *http.Request){
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}
