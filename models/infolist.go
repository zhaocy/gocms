package models

import "github.com/astaxie/beego/orm"

type InfoList struct {
	Id         int
	ClassId    int
	Title      string
	Author     string
	Keywords   string
	Desc       string
	Content    string
	Picurl     string
	Media      string
	Posttime   int64
	Updatetime int64
	Status     int
	Orderid    int
}

/*
 *
 */
func NewsGetList(page, pageSize int, filters ...interface{}) ([]*InfoList, int64) {
	offset := (page - 1) * pageSize
	list := make([]*InfoList, 0)
	query := orm.NewOrm().QueryTable(TableName("info_list")) //pp_info_list
	if len(filters) > 0 {
		length := len(filters)
		for k := 0; k < length; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	count, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, count
}

func NewsGetById(id int) (*InfoList, error) {
	r := new(InfoList)
	err := orm.NewOrm().QueryTable("info_list").Filter("id", id).One(r)
	if err != nil {
		return nil,err
	}
	return r,nil
}
