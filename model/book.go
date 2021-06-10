package model

type Book struct {
	ID      int
	Title   string
	Author  string  //作者
	Price   float64 //价格
	Sales   int64   //促销价
	Stock   int64   //库存
	ImgPath string  //图片路径
}
