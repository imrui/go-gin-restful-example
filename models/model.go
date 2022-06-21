package models

import (
	"go-gin-restful-example/conf"
	"time"
)

type Model struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Publisher struct {
	Model
	Name string `json:"name" form:"addr"` // 名称
	Addr string `json:"addr" form:"addr"` // 地址
}

type Book struct {
	Model
	Name        string `json:"name" form:"name"`                                // 名称
	ISBN        string `json:"isbn" form:"isbn" gorm:"unique"`                  // ISBN编号
	PublishDate string `json:"publishDate" form:"publishDate" gorm:"type:date"` // 出版日期
	PublisherId int    `json:"publisherId" form:"publisherId"`                  // 出版社ID
}

type Author struct {
	Name   string `json:"name" form:"name"`     // 名称
	Gender int8   `json:"gender" form:"gender"` // 性别
	Age    int    `json:"age" form:"age"`       // 年龄
	Tel    string `json:"tel" form:"tel"`       // 电话
}

type AuthorBook struct {
	Model
	AuthorId int `json:"authorId" form:"authorId"` // 作者ID
	BookId   int `json:"bookId" form:"bookId"`     // 书籍ID
}

// ModelType 泛型 类型参数
type ModelType interface {
	Publisher | Book | Author | AuthorBook
}

func AutoMigrate() (err error) {
	err = conf.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Publisher{}, &Book{}, &Author{}, &AuthorBook{})
	return
}
