package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/mingzhehao/scloud/g"
)

func GetCatalogs(currPage, pageSize int) ([]*Catalog, int64, error) {
	catalogs := make([]*Catalog, 0)
	total, err := Catalogs().OrderBy("-display_order").Limit(pageSize, (currPage-1)*pageSize).All(&catalogs)
	if err != nil {
		return nil, 0, err
	}
	return catalogs, total, err
}

func GetCatalogById(id int64) *Catalog {
	if id == 0 {
		return nil
	}
	c := Catalog{Id: id}
	err := orm.NewOrm().Read(&c, "Id")
	if err != nil {
		return nil
	}
	return &c
}

func GetCatalogIdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}
	c := Catalog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return 0
	}
	return c.Id
}

func GetCatalogByIdent(ident string) *Catalog {
	c := Catalog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}
	return &c
}

func CheckCatalogIdentExists(ident string) bool {
	id := GetCatalogIdByIdent(ident)
	return id > 0
}

func SaveCatalog(this *Catalog) (int64, error) {
	if CheckCatalogIdentExists(this.Ident) {
		return 0, fmt.Errorf("catalog english identity exists")
	}
	num, err := orm.NewOrm().Insert(this)
	if err == nil {
		g.CatalogCacheDel("ids")
	}

	return num, err
}

func DelCatalog(c *Catalog) error {
	num, err := orm.NewOrm().Delete(c)
	if err != nil {
		return err
	}

	if num > 0 {
		g.CatalogCacheDel("ids")
	}
	return nil
}

func UpdateCatalog(this *Catalog) error {
	if this.Id == 0 {
		return fmt.Errorf("primary key id not set")
	}
	_, err := orm.NewOrm().Update(this)
	if err == nil {
		g.CatalogCacheDel(fmt.Sprintf("%d", this.Id))
	}
	return err
}

func Catalogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Catalog))
}
