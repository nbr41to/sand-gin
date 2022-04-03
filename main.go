package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Author struct {
	Id   int    `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func Database() *gorm.DB {
	dsn := "fvfmgq0050t5:pscale_pw_IuYFSVdP5bsQeFj7hoDVGDi-seDFM5M4lZEL6pQ9v6o@tcp(os6jsax9p8am.ap-northeast-2.psdb.cloud)/app-db?tls=true"
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

// var author []Author

// var authors []Author

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
		newAuthor := Author{
			Name: "test name",
		}
		result := db.Create(&newAuthor)
		fmt.Println(result)
		c.JSON(200, result)
	})

	router.Run()
}
