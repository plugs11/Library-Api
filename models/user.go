package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"manan.tola/config"
)

var db *sql.DB

type Book struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear string `json:"publishedyear"`
	Genre         string `json:"genre"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (u *Book) CreateBook() *Book {
	QueryCreate := "INSERT INTO book (id,title,author,publishedyear,genre) VALUES (?,?,?,?,?)"
	if _, err := db.Exec(QueryCreate, u.Id, u.Title, u.Author, u.PublishedYear, u.Genre); err != nil {
		log.Println("Error in creating book", err)
	}
	return u
} //Lead Create

func GetAllBook() []Book {
	var books []Book
	list, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Println("Error in getting all the books", err)
	}
	for list.Next() {
		var book Book
		err := list.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedYear, &book.Genre)
		if err != nil {
			log.Println("Error in getting all the books", err)
		}
		books = append(books, book)
	}
	return books
} //Lead Listing(getting all the books)

func GetBookByID(id string) []Book {
	var books []Book
	list, err := db.Query("SELECT id,title,author,publishedyear,genre FROM book WHERE id = ?", id)
	if err != nil {
		log.Println("Error in getting book", err)
	}
	for list.Next() {
		var book Book
		err := list.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedYear, &book.Genre)
		if err != nil {
			log.Println("Error in getting book", err)
		}
		books = append(books, book)
	}
	return books
} // Get all the data of the single book and then update the specific row:column

func UpdateBook(id string, data Book, c *gin.Context) []Book {
	switch {
	case data.Title != "":
		query := " UPDATE book SET title = ? WHERE id = ?"
		_, err := db.Exec(query, data.Title, id)
		if err != nil {
			log.Println("Error in Updating password", err)
		}
	case data.Author != "":
		query := " UPDATE book SET author = ? WHERE id = ?"
		_, err := db.Exec(query, data.Author, id)
		if err != nil {
			log.Println("Error in Updating fullname", err)
		}

	case data.PublishedYear != "":
		query := " UPDATE book SET publishedyear = ? WHERE id = ?"
		_, err := db.Exec(query, data.PublishedYear, id)
		if err != nil {
			log.Println("Error in Updating phonenumber", err)
		}

	case data.Genre != "":
		query := " UPDATE book SET genre = ? WHERE id = ?"
		_, err := db.Exec(query, data.Genre, id)
		if err != nil {
			log.Println("Error in Updating dateofbirth", err)
		}

	default:
		c.JSON(400, gin.H{"error": "Invalid INPUT/s"})
		return nil
	}
	DisplayUpdatedBook := GetBookByID(id)
	return DisplayUpdatedBook
} // update the specific row:column

func DeleteBook(id string) []Book {
	_, err := db.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		log.Println("Error in deleting book", err)
	}
	DisplayAllBooks := GetAllBook()
	return DisplayAllBooks
} // delete a book(and print the remaining books)
