package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Book struct {
	// TODO: Finish struct
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}
var cnt = 1
var bookshelf = []Book{
	// TODO: Init bookshelf
	{1, "Blue Bird", 500},
}

func getBooks(c *gin.Context) {
	c.JSON(200, bookshelf)
}
func getBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, book := range bookshelf {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}
	c.JSON(404, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(404, gin.H{"message": "book not found"})
		return
	}

	for _, book := range bookshelf {
		if book.Name == newBook.Name {
			c.JSON(409, gin.H{"message": "duplicate book name"})
			return
		}
	}
	newBook.ID = cnt + 1
	cnt = cnt + 1
	bookshelf = append(bookshelf, newBook)
	c.JSON(201, newBook)
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, book := range bookshelf {
		if book.ID == id {
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			c.JSON(204, nil)
			return
		}
	}
	c.JSON(204, nil)
}
func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(404, gin.H{"message": "book not found"})
		return
	}

	for _, book := range bookshelf {
		if book.Name == updatedBook.Name {
			c.JSON(409, gin.H{"message": "duplicate book name"})
			return
		}
	}

	for i, book := range bookshelf {
		if book.ID == id {
			bookshelf[i] = updatedBook
        	bookshelf[i].ID = id
			c.JSON(200, bookshelf[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "book not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}
