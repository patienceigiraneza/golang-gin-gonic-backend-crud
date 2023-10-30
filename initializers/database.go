package initializers

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
    "log"
)

var DB *gorm.DB // Export the DB variable

func ConnectToDb() (*gorm.DB, error) {
    // postgres://drijhxnt:Cd8gvEzVvXLRq4zUIJFaDa0oNWiy4zy8@flora.db.elephantsql.com/drijhxnt
    dsn := "host=flora.db.elephantsql.com user=drijhxnt password=Cd8gvEzVvXLRq4zUIJFaDa0oNWiy4zy8 dbname=drijhxnt port=5432 sslmode=disable TimeZone=Africa/Cairo"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect")
    } else {
        fmt.Println("Connected")
    }

    DB = db // Set the DB variable to the connected database

    return db, err
}
