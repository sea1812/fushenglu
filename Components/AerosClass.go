package Components

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

/*
	Aerospike 数据库操作类
	为简化起见，先使用Redis作为数据库后端
*/

type TStorage struct {
	Redis *gredis.Redis //Redis对象的名称
}

// Init 初始化内置的Redis对象指针
func (p *TStorage) Init(AName string) {
	p.Redis = g.Redis(AName)
}

// Get 取值，并转换成JSON对象指针
func (p *TStorage) Get(AKey string) (error, *gjson.Json) {
	res, err := p.Redis.Do("GET", AKey)
	if err != nil {
		return err, nil
	} else {
		mJson := gjson.New(res)
		return nil, mJson
	}
}

// Set Set赋值，把传入的JSON对象转为字符串后保存
func (p *TStorage) Set(AKey string, AJson *gjson.Json) error {
	mJson := AJson.Export()
	_, err := p.Redis.Do("SET", AKey, mJson)
	return err
}
