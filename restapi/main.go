package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	_ "github.com/go-sql-driver/mysql"
)

var books []Book
var dbHost, dbPass string

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Isbn  string `json:"isbn"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	books = nil
	connectString := "root:" + dbPass + "@(" + dbHost + ")/testlocal"
	fmt.Println("Connection string is ", connectString)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		fmt.Println("There was an error connecting to the mysql datbase and the error is ", err)
	} else {
		fmt.Println("Connections sucessful.")
	}
	defer db.Close()
	tx, e := db.Begin()
	if e != nil {
		fmt.Println("There was an error opening the tx to read the data")
	}
	rows, err := tx.Query("select * from books")
	if err != nil {
		fmt.Println("There was an error getting the users ", err)
	}

	defer rows.Close()

	fmt.Println("Listing all the books that are there in the database")
	for rows.Next() {
		var id, title, isbn string
		err := rows.Scan(&id, &title, &isbn)
		fmt.Println("Values of id, title and isbn is ", id, title, isbn)
		aBook := Book{
			Id:    id,
			Title: title,
			Isbn:  isbn,
		}

		books = append(books, aBook)

		if err != nil {
			fmt.Println("There was na error getting the rows ", err)
		}
	}
	//_ = rows.Err()
	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func postBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	aBook := &Book{}
	err := json.NewDecoder(r.Body).Decode(aBook)
	if err != nil {
		fmt.Println("There is an error getting the request body ", err)
	}

	fmt.Println("The id of the book that we've got is ", aBook.Id)

	connectString := "root:" + dbPass + "@(" + dbHost + ")/testlocal"
	fmt.Println("Connection string is ", connectString)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		fmt.Println("There was an error connecting to the mysql datbase and the error is ", err)
	} else {
		fmt.Println("Connections sucessful.")
	}
	defer db.Close()
	dbQuery := "insert into books values ('" + aBook.Id + "', '" + aBook.Title + "', '" + aBook.Isbn + "') "

	tx, e := db.Begin()
	if e != nil {
		fmt.Println("There was an error opening the transaction")
	}

	_, errr := tx.Exec(dbQuery)

	if errr != nil {
		fmt.Println("There was an error posting the Book", errr)
	}

	tx.Commit()

}

func main() {
	path := os.Getenv("APIPATH")

	if path == "" {
		path = "/api/books"
	}
	r := mux.NewRouter()

	dbHost = os.Getenv("DBHOST")
	if dbHost == "" {
		fmt.Println("Environment variable for DBHOST is not defined, setting the default one, that is localhost:3306.")
		dbHost = "localhost:3306"
	}

	dbPass = os.Getenv("DBPASS")
	if dbPass == "" {
		fmt.Println("Environment variable DBPASS is not defined, using the default one, that is 1Vivek$ingh")
		dbPass = "1Vivek$ingh"
	}

	r.HandleFunc(path, getBooks).Methods("GET")
	r.HandleFunc(path, postBook).Methods("POST")

	//log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
