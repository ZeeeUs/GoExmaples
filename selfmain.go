// Пример тривиального API сервиса
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
}

func deleteBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
}

func main()  {
	books = append(books,Book{Id: "1", Title: "Война и Мир", Author: &Author{FirstName: "Лев", LastName: "Толстой"}})
	books = append(books,Book{Id: "2", Title: "Капитанская дочка", Author: &Author{FirstName: "Александр", LastName: "Пушкин"}})
	http.HandleFunc("/", getBooks)
	http.HandleFunc("/add", addBook)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
