package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id         int
	ClassId    int
	Title      string
	SubTitle   string
	Keywords   string
	Content    string
	Desc       string
	Author     string
	Used       int
	Picurl     string
	Linkurl    string
	Media      string
	Orderid    int
	Hits       int
	Status     int
	Posttime   int64
	Updatetime int64
}

func (self *Article) TableName() string {
	return TableName("article")
}

func ArticleGetList() []*Article {
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(TableName("article"))
	query.All(&list)
	return list
}
