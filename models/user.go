package models

import (
	"time"
)

type User struct {
	Id         int       `json:"id" form:"id"`
	UserName   string    `json:"user_name" form:"user_name"`
	WxID       string    `json:"wx_id" form:"wx_id"`
	WxNickName string    `json:"wx_nick_name" form:"wx_nick_name"`
	IDCode     string    `json:"id_code" form:"id_code"`
	Tags       string    `json:"tags" form:"tags"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
	Update     time.Time `json:"update_time" form:"update_time"`
	Extra      string    `json:"extra" form:"extra"`
}

type TradeOrder struct {
	Id         int       `json:"id" form:"id"`
	UserName   string    `json:"user_name" form:"user_name"`
	WxID       string    `json:"wx_id" form:"wx_id"`
	IDCode     string    `json:"id_code" form:"id_code"`
	Tags       string    `json:"tags" form:"tags"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
	Update     time.Time `json:"update_time" form:"update_time"`
	Extra      string    `json:"extra" form:"extra"`
}
