package dal

import (
	"fmt"
)

// 添加数据

func (user *models.User) Add() {
	conn := GetDb()
	defer conn.Close()

	err := conn.Create(user).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}

//修改数据
func (user *models.User) Update() {
	conn := GetDb()
	defer conn.Close()

	err := conn.Model(user).Update(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

//删除数据
func (user *models.User) Del() {
	conn := GetDb()
	defer conn.Close()

	err := conn.Delete(user).Error
	if err != nil {
		fmt.Println("删除失败")
	}
}
