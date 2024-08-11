package book

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type Handler struct{}

type PageData struct {
	Title string
	Books []Book
}

var bookList []Book = loadBooks()

func (h *Handler) renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	layoutDir := "layouts"
	partialsDir := "partials"

	t := template.Must(template.ParseFiles(
		filepath.Join("views", layoutDir, "base.html"),
		filepath.Join("views", partialsDir, "search.html"),
		filepath.Join("views", partialsDir, "book.html"),
		filepath.Join("views", partialsDir, "add.html"),
		filepath.Join("views", templateName),
	))

	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "My Reading List",
	}

	h.renderTemplate(w, "index.html", data)
}

func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	partialsDir := "partials"

	data := struct {
		Books []Book
	}{
		Books: bookList,
	}

	t := template.Must(template.ParseFiles(
		filepath.Join("views", "list.html"),
		filepath.Join("views", partialsDir, "book.html"),
	))

	err := t.ExecuteTemplate(w, "list", data)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")

	newBook := Book{
		ID:     uuid.NewString(),
		Title:  title,
		Author: author,
	}

	bookList = append(bookList, newBook)

	redirectRoute := fmt.Sprintf("/books/%s", newBook.ID)

	http.Redirect(w, r, redirectRoute, http.StatusSeeOther)
}

func (h *Handler) Details(w http.ResponseWriter, r *http.Request) {
	bookIndex := -1
	id := r.PathValue("id")

	for i := range bookList {
		if bookList[i].ID == id {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Book Book
	}{
		Book: bookList[bookIndex],
	}

	partialsDir := "partials"

	t := template.Must(template.ParseFiles(
		filepath.Join("views", "details.html"),
		filepath.Join("views", partialsDir, "book.html"),
	))

	err := t.ExecuteTemplate(w, "details", data)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	bookIndex := -1
	id := r.PathValue("id")

	for i, book := range bookList {
		if book.ID == id {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Book Book
	}{
		Book: bookList[bookIndex],
	}

	partialsDir := "partials"

	t := template.Must(template.ParseFiles(
		filepath.Join("views", "edit.html"),
		filepath.Join("views", partialsDir, "book.html"),
	))

	err := t.ExecuteTemplate(w, "edit", data)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	bookIndex := -1
	id := r.PathValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")

	updatedBook := Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	for i := range bookList {
		if bookList[i].ID == id {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		http.NotFound(w, r)
		return
	}

	bookList[bookIndex] = updatedBook

	redirectRoute := fmt.Sprintf("/books/%s", updatedBook.ID)

	http.Redirect(w, r, redirectRoute, http.StatusSeeOther)
}

func (h *Handler) Destroy(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	filteredBooks := []Book{}

	for _, book := range bookList {
		if book.ID != id {
			filteredBooks = append(filteredBooks, book)
		}
	}

	bookList = filteredBooks
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	searchTerm := strings.ToLower(r.FormValue("search"))

	filteredBookList := []Book{}

	for _, book := range bookList {
		if strings.Contains(strings.ToLower(book.Title), searchTerm) || strings.Contains(strings.ToLower(book.Author), searchTerm) {
			filteredBookList = append(filteredBookList, book)
		}
	}

	log.Printf("Filtered Books:\n%v", filteredBookList)

	t, _ := template.ParseFiles("views/list.html")
	err := t.Execute(w, filteredBookList)
	if err != nil {
		log.Println(err)
	}
}
