package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
	"log"
	"strconv"
)

func AddBook(b *model.Book) error {
	sql := "insert into books(title, author, price, sales, stock, img_path) values(?,?,?,?,?,?)"
	_, err := db.Db.Exec(sql, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)
	log.Println("b = ", b)
	if err != nil {
		return err
	}
	return nil
}

func GetBooks() ([]*model.Book, error) {
	sql := "select id, title, author, price, sales,stock, image_path from books"
	rows, err := db.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func DeleteBook(bookId string) error {
	sql := "delete from books where id=?"
	_, err := db.Db.Exec(sql, bookId)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(bookID string) (*model.Book, error) {
	sql := "select id, title, author, price, sales, stock, img_path from books where id=?"
	row := db.Db.QueryRow(sql, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

func UpdateBook(b *model.Book) error {
	sql := "update books set title=?, author=?,price=?,sales=?,stock=? where id=?"
	_, err := db.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql := "select count(*) from books"
	var totalRecord int64
	row := db.Db.QueryRow(sql)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	//如果记录数刚好除尽每页显示的数量,则总页数不需要+1
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//获取当前页的图书
	sql2 := "select id, title, author, price, sales, stock, img_path from books limit ?,?"
	rows, err := db.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}

//根据价格区间筛选图书
func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql := "select count(*) from books where price between ? and ?"
	var totalRecord int64
	row := db.Db.QueryRow(sql, minPrice, maxPrice)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	sql2 := "select id, title, author, price, sales, stock, img_path from books where price between ? and ? limit ?, ?"
	rows, err := db.Db.Query(sql2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil

}
