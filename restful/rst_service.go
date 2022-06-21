package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-restful-example/conf"
	"go-gin-restful-example/models"
	"reflect"
)

var Ctx = context.Background()

// DbFindAll 查询所有数据
func DbFindAll[T models.ModelType](q *T) (items []*T, err error) {
	if q == nil {
		err = conf.DB.Find(&items).Error // 查询所有数据
	} else {
		err = conf.DB.Where(&q).Find(&items).Error // 根据条件查询
	}
	return
}

// DbFind 查询单条数据
func DbFind[T models.ModelType](id int) (item *T, err error) {
	// 先读缓存
	item, err = RdbGet[T](id)
	if err == nil {
		return
	}
	// 缓存没有，从数据库读
	err = conf.DB.First(&item, id).Error
	if err != nil {
		return
	}
	// 写缓存
	err = RdbSet[T](id, item)
	return
}

// DbAdd 添加数据
func DbAdd[T models.ModelType](in *T) (err error) {
	// 写数据库
	err = conf.DB.Create(&in).Error
	if err != nil {
		return
	}
	// 反射获取id
	id := int(reflect.ValueOf(*in).FieldByName("ID").Int())
	// 写缓存
	err = RdbSet[T](id, in)
	return
}

// DbUpdate 更新数据
func DbUpdate[T models.ModelType](id int, in *T) (out *T, err error) {
	out, err = DbFind[T](id)
	if err != nil {
		return
	}
	// 更新数据库
	err = conf.DB.Model(&out).Updates(in).Error
	if err != nil {
		return
	}
	// 更新缓存
	err = RdbSet[T](id, out)
	return
}

// DbDelete 删除数据
func DbDelete[T models.ModelType](id int) (err error) {
	var m T
	err = conf.DB.Delete(&m, id).Error
	if err != nil {
		return
	}
	// 删除缓存
	err = RdbDel[T](id)
	return
}

// genCacheKey 生成缓存key
func genCacheKey(name string, id int) string {
	return fmt.Sprintf("%s:id-%d", name, id)
}

// RdbSet 写缓存（函数的id参数为了减少反射读取属性值）
func RdbSet[T models.ModelType](id int, in *T) (err error) {
	key := genCacheKey(reflect.TypeOf(*in).Name(), id)
	// 序列化
	data, err := json.Marshal(in)
	if err != nil {
		fmt.Println("RdbSet:", err)
		return
	}
	// 写redis
	err = conf.Rdb.Set(Ctx, key, string(data), conf.Cfg.Redis.Expiration).Err()
	return
}

// RdbGet 读缓存
func RdbGet[T models.ModelType](id int) (out *T, err error) {
	var tmp T
	key := genCacheKey(reflect.TypeOf(tmp).Name(), id)
	// 读redis
	data, err := conf.Rdb.Get(Ctx, key).Result()
	if err != nil {
		fmt.Println("RdbGet:", err)
		return
	}
	// 反序列化
	err = json.Unmarshal([]byte(data), &out)
	return
}

// RdbDel 删除缓存
func RdbDel[T models.ModelType](id int) (err error) {
	var t T
	key := genCacheKey(reflect.TypeOf(t).Name(), id)
	err = conf.Rdb.Del(Ctx, key).Err()
	return
}
