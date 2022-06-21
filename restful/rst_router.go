package restful

import (
	"github.com/gin-gonic/gin"
	"go-gin-restful-example/models"
)

func LoadRouters(e *gin.Engine) {
	r := e.Group("/api")
	// publishers
	ph := RstHandler[models.Publisher]{}
	r.GET("/publishers", ph.FindAll)
	r.POST("/publishers", ph.Add)
	r.GET("/publishers/:id", ph.Find)
	r.PUT("/publishers/:id", ph.Update)
	r.DELETE("/publishers/:id", ph.Delete)
	// books
	bh := RstHandler[models.Book]{}
	r.GET("/books", bh.FindAll)
	r.POST("/books", bh.Add)
	r.GET("/books/:id", bh.Find)
	r.PUT("/books/:id", bh.Update)
	r.DELETE("/books/:id", bh.Delete)
	// authors
	ah := RstHandler[models.Author]{}
	r.GET("/authors", ah.FindAll)
	r.POST("/authors", ah.Add)
	r.GET("/authors/:id", ah.Find)
	r.PUT("/authors/:id", ah.Update)
	r.DELETE("/authors/:id", ah.Delete)
	// author books
	abh := RstHandler[models.AuthorBook]{}
	r.GET("/author/books", abh.FindAll)
	r.POST("/author/books", abh.Add)
	r.GET("/author/books/:id", abh.Find)
	r.PUT("/author/books/:id", abh.Update)
	r.DELETE("/author/books/:id", abh.Delete)
}
