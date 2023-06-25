package durable

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func InitMysqlDb() *gorm.DB {
	var connectionSecret = "417927vb147o9m8wh5ls:pscale_pw_nORDgMmM6ppUq2b1yUDnuDELVJ9OINCULSJyu1WmUDw@tcp(aws.connect.psdb.cloud)/streamsquad?tls=true"
	Mysqldb, err := gorm.Open(mysql.Open(connectionSecret), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale Mysql Db: %v", err)
	}
	//Mysqldb.Table("user").AutoMigrate(&models.User{})

	return Mysqldb

}
func CloseDbConn(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		log.Fatal("could not close conn:", err)
	}
	dbSql.Close()
}
