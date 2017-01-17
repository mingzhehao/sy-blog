package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Message struct {
	Id         int64
	Uid        int64
	Uname      string `orm:"size(255),unique"`
	Status     int8
	Content    string    `orm:size(255)`
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(Message))
}

func GetMessageCount() (int64, error) {
	messages := make([]*Message, 0)
	nums, err := OrmMessage().All(&messages)
	return nums, err
}

func GetMessages(currPage, pageSize int) ([]*Message, int64, error) {
	messages := make([]*Message, 0)
	total, err := OrmMessage().OrderBy("-created_at").Limit(pageSize, (currPage-1)*pageSize).All(&messages)
	if err != nil {
		return nil, 0, err
	}
	return messages, total, err
}

func DeleteMessage(id string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	message := &Message{
		Id: aid,
	}
	_, err = o.Delete(message)
	if err != nil {
		return err
	}
	return nil
}

func MessageDetail(id string) (*Message, error) {
	aid, err := strconv.ParseInt(id, 10, 64)
	message := &Message{}
	err = OrmMessage().Filter("id", aid).One(message)
	return message, err
}

func SearchMessageCount(keyword string) (int64, error) {
	message := make([]*Message, 0)
	total, err := OrmMessage().Filter("uname", keyword).OrderBy("-created_at").All(&message)
	return total, err
}

func SearchMessageByName(currPage, pageSize int, keyword string) ([]*Message, error) {
	var resultRecs []*Message
	var err error
	resultRecs = make([]*Message, 0)
	if len(keyword) > 0 {
		_, err = OrmMessage().Filter("uname", keyword).OrderBy("created_at").Limit(pageSize, (currPage-1)*pageSize).All(&resultRecs)
	}
	return resultRecs, err
}

func InsertMessage(uid int64, status int8, uname, content string) error {

	stringTime := time.Now().Format("2006-01-02 15:04:05")
	datetime, _ := time.Parse("2006-01-02 15:04:05", stringTime)
	o := orm.NewOrm()
	message := &Message{
		Uid:        uid,
		Uname:      uname,
		Status:     status,
		Content:    content,
		Created_at: datetime,
		Updated_at: datetime,
	}
	_, err := o.Insert(message)
	return err
}

func OrmMessage() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Message))
}
