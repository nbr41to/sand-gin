package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Author struct { // こいつがテーブルになる
	Id   int    `gorm:"column:id;primary_key" json:"id"` // タグにtableのnameとjsonにしたときのnameを指定
	Name string `gorm:"column:name" json:"name"`
}

func Database() *gorm.DB {
	err := godotenv.Load(".env") // env 読み込み
	dsn := os.Getenv("DSN")      // 変数を取得
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		/* テーブル名に関する options */
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 勝手に複数形にしない
			NoLowerCase:   true, // 勝手に小文字にしない
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!!",
		})
	})

	db := Database()

	router.GET("/test", func(c *gin.Context) {
		author := Author{}
		authors := []Author{}

		/* 単体取得 */
		db.First(&author, 2)
		fmt.Println(author)

		/* 全て取得 */
		db.Find(&authors)
		fmt.Println(authors)

		c.JSON(200, authors)
	})

	/* 新規作成 */
	router.POST("/test", func(c *gin.Context) {
		/* データを取得 */
		// var author Author
		author := Author{}
		c.BindJSON(&author)
		fmt.Println(author)

		// newAuthor := Author{
		// 	Name: "test name",
		// }
		// result := db.Create(&newAuthor)
		// fmt.Println(result)
		// c.JSON(200, result)
	})

	router.Run()
}
