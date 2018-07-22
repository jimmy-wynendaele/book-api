package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"html/template"
	"io/ioutil"
	//"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

// Book Structs (MODEL)

type Book struct {
	ID     int    `json:"id"`
	Isbn   int    `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books var as a slice Book struct

// type books struct {
// 	Something []Book `json:"bookjes"`
// }

// var booking books

var booking []Book

// GET ALL BOOKS

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// can I put following lines into a func outside the getBooks func so it would be
	// accessible for other functions as well? Defer?

	//var toBooking books
	configFile, err := os.Open("config.json")

	defer configFile.Close()

	if err != nil {

		fmt.Println("come on")
		//return booking, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&booking)

	json.NewEncoder(w).Encode(booking)
	fmt.Println(err)
	//return toBooking, err

}

//Get Single Book

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	configFile, err := os.Open("config.json")

	defer configFile.Close()

	if err != nil {

		fmt.Println("problem")
		//return booking, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&booking)

	var isbnBook = params["isbn"]

	var isbnToInt, _ = strconv.Atoi(isbnBook)

	for _, item := range booking {

		fmt.Println(item.ID, item, item.Author.Lastname, reflect.TypeOf(item.Author), reflect.TypeOf(isbnToInt))

		if item.Isbn == isbnToInt {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{}) // This returns an empty book struct when it doesn't find the id query. You can omit this line
	// //when you want the page to be blank when not finding the id query

	fmt.Println("heeeeeeeyyy")

}

// DELETE A CERTAIN BOOK

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	params := mux.Vars(r)

	configFile, err := os.Open("config.json")

	defer configFile.Close()

	if err != nil {

		fmt.Println("problem")
		//return booking, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&booking)

	var isbnToInt, _ = strconv.Atoi(params["isbn"])

	for index, item := range booking {
		if item.Isbn == isbnToInt {

			booking = append(booking[:index], booking[index+1:]...)
			fmt.Println("book to delete detected", booking)

			// Marshal encodes booking into json format
			out, err := json.Marshal(booking)

			if err != nil {
				fmt.Println(err)
			}

			// WriteFile writes the out(the json format of booking) into the file 'config.json'
			ioutil.WriteFile("config.json", out, 0777)
			break
		}
	}
	json.NewEncoder(w).Encode(booking)
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	configFile, err := os.Open("config.json")

	defer configFile.Close()

	if err != nil {

		fmt.Println("problem")
		//return booking, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&booking)

	r.ParseForm()

	fmt.Println(r.Form["authorLastName"])

	var lastname = r.Form["authorLastName"]
	var firstname = r.Form["authorFirstName"]
	var title = r.Form["title"]
	var isbn = r.Form["isbn"]
	var bookId = rand.Intn(100) //converts integer to a string. This is a Mock ID, because it could generaate the same ID
	var isbnConv, _ = strconv.Atoi(isbn[0])

	booking = append(booking, Book{ID: bookId, Isbn: isbnConv, Title: title[0], Author: Author{Firstname: firstname[0], Lastname: lastname[0]}})

	fmt.Println(booking)

	// Marshal encodes booking into json format
	out, err := json.Marshal(booking)
	if err != nil {
		fmt.Println(err)
	}
	// WriteFile writes the out(the json format of booking) into the file 'config.json'
	ioutil.WriteFile("config.json", out, 0777)

	//http.Redirect(w, r, "ile:///Users/jimmy/go/src/github.com/jimmy-wynendaele/restapi2/index.html", 301)

}

// UPDATE A BOOK

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	r.ParseForm()

	var lastname = r.Form["authorLastName"][0]
	var firstname = r.Form["authorFirstName"][0]
	var titleBook = r.Form["title"][0]
	var isbn, _ = strconv.Atoi(r.Form["isbn"][0])

	fmt.Println(lastname, firstname, titleBook, isbn)

	configFile, err := os.Open("config.json")

	defer configFile.Close()

	if err != nil {

		fmt.Println("problem")
		//return booking, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&booking)

	fmt.Println(booking)

	for index, item := range booking {

		fmt.Println(item, index)

		if item.Isbn == isbn || item.Author.Firstname == firstname || item.Author.Lastname == lastname || item.Title == titleBook {
			booking = append(booking[:index], booking[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			var book Book
			var author Author
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = rand.Intn(100) //converts integer to a string. This is a Mock ID, not to be used in production because it could generate the same ID
			book.Isbn = isbn
			book.Title = titleBook
			author.Firstname = firstname
			author.Lastname = lastname
			book.Author = author
			//fmt.Fprintln(author, book.Author)
			// book.Author.Firstname = firstname
			// book.Author.Lastname = lastname
			booking = append(booking, book)
			fmt.Println(booking)

			// Marshal encodes booking into json format
			out, err := json.Marshal(booking)
			if err != nil {
				fmt.Println(err)
			}
			// WriteFile writes the out(the json format of booking) into the file 'config.json'
			ioutil.WriteFile("config.json", out, 0777)
			return
		}
	}

}

func main() {
	//init Router

	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Route Handlers / Endpoints

	router.HandleFunc("/api/books", getBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{isbn}", getBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{isbn}", deleteBook).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/books", createBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update", updateBook).Methods("POST", "OPTIONS")

	http.ListenAndServe(":"+port, router)

}
