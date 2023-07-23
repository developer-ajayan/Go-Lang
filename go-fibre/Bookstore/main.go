package main

import (
	"log"

	Apps "Bookstore/Config"
	Routes "Bookstore/Routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	// conn, err := database.ConnectDatabase()
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Database connection not established")
	// }
	// execute a query
	// defer conn.Close()
	// dbi := database.DbInfo{}
	// dbi.DBConn = conn
	// dbi.Query = query
	// newTable := "bookslist"
	// model := "(id INT PRIMARY KEY , bookname CHAR(50),author CHAR(50))"
	// if _, err := dbi.CreateTable(newTable, model); err != nil {
	// 	panic("table creation failed")
	// }
	// row := map[string]interface{}{
	// 	"id":       01,
	// 	"bookname": "testone",
	// 	"author":   "ajayan",
	// }

	// Insert row into "employees" table
	// err = dbi.InsertRow("bookslist", row)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	config := Apps.FiberConfig()
	app := fiber.New(config)

	// routes
	Routes.Bookroutes(app)

	log.Fatal(app.Listen(":4000"))

}
