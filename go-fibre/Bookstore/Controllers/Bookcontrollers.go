package controllers

import (
	database "Bookstore/Database"
	models "Bookstore/Models"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func Test(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Test url working")
}

func GetBooks(c *fiber.Ctx) error {
	book_list := models.Books
	fmt.Println(len(book_list))

	// retrive book name from request url
	bookname := "'%" + string(c.Request().URI().QueryString()) + "%'"
	// fecth from database related data
	query := fmt.Sprintf("Select name,author from %s where name LIKE %s", "bookapi_books", bookname)
	// return the list
	fmt.Println(query)
	Db := database.DbInfo{}
	Db.Query = query
	if conn, err := database.ConnectDatabase(); err != nil {
		defer Db.DBConn.Close()
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "books creation failed",
		})
	} else {
		Db.DBConn = conn
		defer Db.DBConn.Close()
		if book_list, erro := Db.FetchRows(); erro == nil {
			fmt.Println("dataa", book_list)
			if len(book_list) > 0 {
				return c.JSON(book_list)
			}
		} else {
			// snet to senrty
			fmt.Println("fetch row error", erro)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "No books Found",
	})
}

func CreateBooks(c *fiber.Ctx) error {

	var data struct {
		Books []models.Book `json:"books"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	books_list := make(map[string]interface{})
	for _, item := range data.Books {
		books_list["name"] = item.Title
		books_list["author"] = item.Author
	}
	Db := database.DbInfo{}

	if conn, err := database.ConnectDatabase(); err != nil {
		defer Db.DBConn.Close()
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "books creation failed",
		})
	} else {
		Db.DBConn = conn
		defer Db.DBConn.Close()
		if erro := Db.InsertRow("bookapi_books", books_list); erro != nil {
			fmt.Println("eroo", erro)
		}
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "books created successfully",
	})
}
