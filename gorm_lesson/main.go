package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	dsn := "root:lpf123456@tcp(127.0.0.1:3306)/gormLesson?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	// 创建
	//db.Create(&User{
	//	Name: "Mary",
	//	Age:  67,
	//	Sex:  false,
	//})

	// 查询
	// 查询一条
	//var u User
	//db.First(&u) // 默认的查询第一条
	//db.Where("id = ?", 2).First(&u)
	//db.First(&u, "id = ?", 3)
	//db.Where("id in (?)", []int{3, 2}).First(&u) // 根据条件查询第一条
	//fmt.Println(u)

	// 查询多条
	//var u []User
	//db.Find(&u)
	//db.Find(&u, "id > ?", 1)
	//db.Where("id in (?)", []int{1, 3}).Find(&u)
	//db.Where("id = ?", 3).Or("id = ?", 1).Find(&u)
	//fmt.Println(u)

	// 更新
	//db.Model(&User{}).Where("id = ?", 1).Update("name", "tom")

	// 删除
	//var u User
	//db.Where("id = ?", 1).Delete(&u)
	//db.Where("id = ?", 3).Unscoped().Delete(&u)
}
