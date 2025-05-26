package infrastructure

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func dbSetup() {
    var err error
    l := gormLogger.Default.LogMode(gormLogger.Silent)

    switch cfg.Database.Driver {
    case "mysql":
        db, err = gorm.Open(mysql.New(mysql.Config{
            DSN: cfg.Database.DSN,
        }), &gorm.Config{
            Logger: l,
        })
    case "sqlite":
        db, err = gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{
            Logger: l,
        })
    default:
        panic("Unsupported DB driver: " + cfg.Database.Driver)
    }

    if err != nil {
        panic(err)
    }


    // Test query minimal untuk cek koneksi
    sqlDB, err := db.DB()
    if err != nil {
        panic(err)
    }
    if err := sqlDB.Ping(); err != nil {
        panic(err)
    }
}

