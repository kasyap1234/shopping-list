package shoppinglist
import (
	"encore.dev/storage/sqldb"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)
type Service struct {
	db *gorm.DB 

}
var db = sqldb.NewDatabase("shoppinglist",sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

