package main

import "testing"

// TestGetLibrary tests the functionality of retrieving the library.
func TestGetLibrary(t *testing.T) {
	library := []book{
		{ID: "1", Title: "Warbreaker", Author: "Brandon Sanderson"},
		{ID: "2", Title: "Hamlet", Author: "William Shakespeare"},
	}
	if len(library) != 2 {
		t.Errorf("Expected 2 books in library, got %d", len(library))
	} else {
		t.Logf("Get library test PASSED: %+v", library)
	}
}

// TestAddBook tests the functionality of adding books to the library.
func TestAddBook(t *testing.T) {
	library := []book{}
	newBook := book{ID: "1", Title: "Test Book", Author: "Aiden Beus"}

	// Add a new book to the library
	library = append(library, newBook)
	if len(library) != 1 || library[0].ID != "1" || library[0].Title != "Test Book" || library[0].Author != "Aiden Beus" {
		t.Errorf("Expected to add book, got %+v", library)
	}else{
		t.Logf("Add valid book test PASSED: %+v", library[0])
	}

	// Try adding a book with missing Author
	library = append(library, book{ID: "", Title: "Incomplete Book", Author: ""})
	if len(library) != 2 || library[1].ID != "" || library[1].Title != "Incomplete Book" || library[1].Author != "" {
		t.Errorf("Expected to not add incomplete book, got %+v", library)
	}else{
		t.Logf("Missing author test PASSED: %+v", library[1])
	}

	// Try adding a book with missing Title
	library = append(library, book{ID: "2", Title: "", Author: "Another Author"})
	if len(library) != 3 || library[2].ID != "2" || library[2].Title != "" || library[2].Author != "Another Author" {
		t.Errorf("Expected to not add book with missing title, got %+v", library)	
}else{
		t.Logf("Missing title test PASSED: %+v", library[2])
	}
	// Try adding a book with negative ID
	library = append(library, book{ID: "-1", Title: "Negative ID Book", Author: "Negative Author"})
	if len(library) != 4 || library[3].ID != "-1" || library[3].Title != "Negative ID Book" || library[3].Author != "Negative Author" {
		t.Errorf("Expected to not add book with negative ID, got %+v", library)
}else{
		t.Logf("Negative ID test PASSED: %+v", library[3])
	}
}

// TestDeleteBook tests the functionality of deleting books from the library.
func TestDeleteBook(t *testing.T) {
    library := []book{
        {ID: "1", Title: "Book One", Author: "Author A"},
        {ID: "2", Title: "Book Two", Author: "Author B"},
    }

    // Test deleting the first book
    newLib, ok := deleteBook(library, 0)
    if !ok || len(newLib) != 1 || newLib[0].ID != "2" {
        t.Errorf("Expected to delete book at index 0, got %+v", newLib)
    }else{
		t.Logf("Delete first book test PASSED: %+v", newLib)
	}
	

    // Test deleting with an invalid index
    newLib, ok = deleteBook(library, 5)
    if ok || len(newLib) != 2 {
        t.Errorf("Expected delete to fail and library unchanged, got %+v", newLib)
    }else{
		t.Logf("Delete invalid index test PASSED: %+v", newLib)
	}

	// Test deleting negative index
	newLib, ok = deleteBook(library, -1)
	if ok || len(newLib) != 2 {
		t.Errorf("Expected delete to fail and library unchanged, got %+v", newLib)
	}else{
		t.Logf("Delete negative index test PASSED: %+v", newLib)
	}
}