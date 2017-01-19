package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id         int64
	Bid        int64
	Name       string
	Count      int64
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(Tag))
}

func GetTagCount() (int64, error) {
	tags := make([]*Tag, 0)
	nums, err := OrmTag().All(&tags)
	return nums, err
}

func GetTags(currPage, pageSize int) ([]*Tag, int64, error) {
	tags := make([]*Tag, 0)
	total, err := OrmTag().OrderBy("-created_at").Limit(pageSize, (currPage-1)*pageSize).All(&tags)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, err
}

func DeleteTag(id int64, tagName string) error {
	c := Tag{Name: tagName}
	o := orm.NewOrm()
	err := o.Read(&c, "Name")
	// SELECT `id`, `bid`, `name`, `count`, `created_at` FROM `bb_tag` WHERE `name` = ?] - `test`
	if err == orm.ErrNoRows {
		return nil
	} else if c.Count == 1 {
		tag := &Tag{
			Id: c.Id,
		}
		_, err = o.Delete(tag)
		if err != nil {
			return err
		}
		return nil
	} else {
		count := c.Count - int64(1)
		tagInfo := Tag{Count: count}
		_, err := o.Update(tagInfo)
		return err
	}

}

func InsertTag(bid int64, tagName string) error {

	c := Tag{Name: tagName}
	err := orm.NewOrm().Read(&c, "Name")
	// SELECT `id`, `bid`, `name`, `count`, `created_at` FROM `bb_tag` WHERE `name` = ?] - `test`
	if err == orm.ErrNoRows {
		stringTime := time.Now().Format("2006-01-02 15:04:05")
		datetime, _ := time.Parse("2006-01-02 15:04:05", stringTime)
		o := orm.NewOrm()
		tag := &Tag{
			Bid:        bid,
			Name:       tagName,
			Count:      1,
			Created_at: datetime,
		}
		_, err := o.Insert(tag)
		fmt.Println(err)
		return err
	} else {
		count := c.Count + int64(1)
		tagInfo := Tag{Count: count}
		_, err := orm.NewOrm().Update(tagInfo)
		fmt.Println(err)
		return err
	}
}

func OrmTag() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Tag))
}
