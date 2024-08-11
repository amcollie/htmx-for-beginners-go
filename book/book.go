package book

type Book struct {
	ID     string
	Title  string
	Author string
}

func books() []Book {
	return []Book{
		{ID: "1", Title: "The Catcher in the Rye", Author: "J. D. Salinger"},
		{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		{ID: "3", Title: "1984", Author: "George Orwell"},
		{ID: "4", Title: "Pride and Prejudice", Author: "Jane Austen"},
		{ID: "5", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
		{ID: "6", Title: "The Hobbit", Author: "J. R. R. Tolkien"},
		{ID: "7", Title: "The Adventures of Huckleberry Finn", Author: "Mark Twain"},
		{ID: "8", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
		{ID: "9", Title: "The Picture of Dorian Gray", Author: "Oscar Wilde"},
		{ID: "10", Title: "The Grapes of Wrath", Author: "John Steinbeck"},
	}
}

func loadBooks() []Book {
	return books()
}
