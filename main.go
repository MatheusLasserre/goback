package main

import (
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/MatheusLasserre/goback/my_models"
	_ "github.com/go-sql-driver/mysql"
	sqlcommentercore "github.com/google/sqlcommenter/go/core"
	gosql "github.com/google/sqlcommenter/go/database/sql"
	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	db, err := gosql.Open("mysql", goDotEnvVariable("DSN"), sqlcommentercore.CommenterOptions{
		Config: sqlcommentercore.CommenterConfig{
			EnableRoute:       true,
			EnableController:  true,
			EnableAction:      true,
			EnableDBDriver:    true,
			EnableFramework:   true,
			EnableTraceparent: true,
			EnableApplication: true,
		},
		Tags: sqlcommentercore.StaticTags{
			Application: "GOLANG Commenter",
			DriverName:  "mysql",
			CustomTag:   "customtag",
		},
	})
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// ctx := context.Background()

	start := time.Now()

	var us models.User

	us.Email = null.StringFrom("sss")
	res, err := db.Query("SELECT * FROM golangapp /*runtime='golangapp'*/")
	defer res.Close()

	if err != nil {
		panic(err)
	}

	if res.Next() {
		err := res.Scan(&us.ID, &us.Name, &us.Email, &us.Password, &us.Image, &us.EmailVerified)
		if err != nil {
			panic(err)
		}
		fmt.Println(us)
	}

	elapsed := time.Since(start)
	log.Printf("Query took %s", elapsed)

}
