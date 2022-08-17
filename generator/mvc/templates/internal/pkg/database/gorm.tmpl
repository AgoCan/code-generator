package database

// https://gorm.io/zh_CN/
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 预留给其他的数据库，所以这么写，例如redis可以继续添加
type Options struct {
	Mysql struct {
		DbName   string
		Password string
		Username string
		Port     string
		Host     string
	}
}

type Client struct {
	DB *gorm.DB
}

func NewClient(o *Options) *Client {
	Connect := o.Mysql.Username + ":" + o.Mysql.Password + "@tcp(" +
		o.Mysql.Host + ":" + o.Mysql.Port + ")/" + o.Mysql.DbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", Connect)
	if err != nil {
		panic(err)
	}
	// 创建数据库
	migrate()

	// 尝试连接
	err = db.DB().Ping()
	return &Client{
		DB: db,
	}

}

func migrate() {
}
