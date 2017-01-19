package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/mingzhehao/scloud/g"
	"time"
)

func GetArticleCount() (int64, error) {
	dbRecs := make([]*Blog, 0)
	nums, err := Blogs().All(&dbRecs)
	return nums, err
}

func GetArticles(currPage, pageSize int) ([]*Blog, int64, error) {
	dbRecs := make([]*Blog, 0)
	total, err := Blogs().OrderBy("-Created_at").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	if err != nil {
		return nil, 0, err
	}
	return dbRecs, total, err
}

func GetHotArticles() ([]*Blog, int64, error) {
	dbRecs := make([]*Blog, 0)
	total, err := Blogs().OrderBy("-views").Limit(10).All(&dbRecs)
	if err != nil {
		return nil, 0, err
	}
	return dbRecs, total, err
}

func GetArticleById(id int64) *Blog {
	if id <= 0 {
		return nil
	}
	o := Blog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func GetArticleByIdent(ident string) *Blog {
	if ident == "" {
		return nil
	}
	c := Blog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}
	return &c
}

func GetArticleIdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}
	c := Blog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return 0
	}

	return c.Id
}

func CheckIdentExists(ident string) bool {
	id := GetArticleIdByIdent(ident)
	return id > 0
}

func GetArticleIds(catalog_id int64) []int64 {
	if catalog_id <= 0 {
		return []int64{}
	}

	var blogs []Blog
	Blogs().Filter("CatalogId", catalog_id).Filter("Status", 1).OrderBy("-Created_at").All(&blogs, "Id")
	size := len(blogs)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i] = blogs[i].Id
	}

	return ret
}

func ReadBlogContent(b *Blog) *BlogContent {
	if b.Id <= 0 || b.BlogContentId <= 0 {
		return nil
	}

	key := fmt.Sprintf("content_of_%d_%d", b.Id, b.Updated_at)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := readBlogContentInDB(b); p != nil {
			g.BlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(BlogContent)
	return &ret
}

func readBlogContentInDB(b *Blog) *BlogContent {
	o := BlogContent{Id: b.BlogContentId}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

type CustomData struct {
	Id         int64
	Ident      string
	Title      string
	Content    string
	Views      int64
	Created_at time.Time
}

func GetArticlesByCatalog(catalog_id int64, offset, limit int) []*CustomData {
	// 构建查询对象
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("bb_blog.*",
		"bb_blog_content.content ").
		From("bb_blog").
		InnerJoin("bb_blog_content").On("bb_blog.blog_content_id = bb_blog_content.id").
		Where("catalog_id = ?").
		OrderBy("created_at").Desc().
		Limit(limit).Offset(offset)

	// 导出SQL语句
	sql := qb.String()

	// 执行SQL语句
	o := orm.NewOrm()
	dbRecs := make([]*CustomData, 0)
	o.Raw(sql, catalog_id).QueryRows(&dbRecs)
	fmt.Println(dbRecs)
	return dbRecs
}

func SaveArticles(this *Blog, blogContent string) (int64, error) {
	if CheckIdentExists(this.Ident) {
		return 0, fmt.Errorf("blog english identity exists")
	}

	bc := &BlogContent{Content: blogContent}
	or := orm.NewOrm()
	blogContentId, e := or.Insert(bc)
	if e != nil {
		return 0, e
	}

	this.BlogContentId = blogContentId
	stringTime := time.Now().Format("2006-01-02 15:04:05")
	this.Updated_at, _ = time.Parse("2006-01-02 15:04:05", stringTime)

	id, err := or.Insert(this)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("article_ids_of_%d", this.CatalogId))
	}

	return id, err
}

func DelArticles(b *Blog) error {
	num, err := Blogs().Filter("Id", b.Id).Delete()
	if err != nil {
		return err
	}

	if num > 0 {
		g.BlogCacheDel(fmt.Sprintf("article_ids_of_%d", b.CatalogId))
		BlogContents().Filter("Id", b.BlogContentId).Delete()
	}

	return nil
}

func UpdateArticles(b *Blog, content string) error {
	if b.Id == 0 {
		return fmt.Errorf("primary key:id not set")
	}

	bc := ReadBlogContent(b)
	if content != "" && bc.Content != content {
		bc.Content = content
		_, e := orm.NewOrm().Update(bc)
		if e != nil {
			return e
		}
		stringTime := time.Now().Format("2006-01-02 15:04:05")
		b.Updated_at, _ = time.Parse("2006-01-02 15:04:05", stringTime)
	}

	_, err := orm.NewOrm().Update(b)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("%d", b.Id))
	}
	return err
}

func Blogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Blog))
}

func BlogContents() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(BlogContent))
}
