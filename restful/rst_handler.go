package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-restful-example/models"
	"net/http"
	"reflect"
	"strconv"
)

type RstHandler[T models.ModelType] struct{}

func (h *RstHandler[T]) FindAll(c *gin.Context) {
	var q T // 查询条件
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	items, err := DbFindAll[T](&q)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *RstHandler[T]) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	item, err := DbFind[T](id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *RstHandler[T]) Add(c *gin.Context) {
	var in T
	err := c.BindJSON(&in)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = DbAdd[T](&in)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, in)
}

func (h *RstHandler[T]) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	var in T
	err = c.BindJSON(&in)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	inId := int(reflect.ValueOf(in).FieldByName("ID").Int())
	if inId != id {
		fmt.Println("error: path id != body id")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	out, err := DbUpdate[T](id, &in)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *RstHandler[T]) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = DbDelete[T](id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
