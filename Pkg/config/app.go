package config

var (
	db *gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
