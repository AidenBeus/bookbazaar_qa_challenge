##Setup

## Clone the project or download zip

```
$ git clone https://github.com/AidenBeus/bookbazaar_qa_challenge.git
$ cd bookbazaar_qa_challenge
```
zip: https://github.com/AidenBeus/bookbazaar_qa_challenge/archive/refs/heads/master.zip

##Files
main.go is where the application is stored and run from.
library.go contains the logic and allows for testing.
library_test.go is the test code with the test cases.
go.sum and go.mod are dependencies.

##Running the project
To run the application, simply run:
```
$cd ui
$go run .
```
If that doesnt work, you may have to download tview.
In order to use tview, download the library.
```
$go mod init 
&go get github.com/rivo/tview@master
``` 
If the app is running correctly, a window will pop up. Use the TAB key to navigate the fields.
To add a book, input book title and author name. Then hit ENTER on the "Add Book" button.
To delete a book, input just the book id. Then hit ENTER on the "Delete Book" button.
To exit the app, hit ENTER on the "Quit" button.

##Running the testcases
Use the test command to run the test cases.
```
$go test -v
```
Each test case should show whether it passed, along with the runtime.
Tests are stored and run from library_test.go

##Rationale
I choose to use tview for my application UI because it is a relatively simple UI library. 
I liked how it looked, and there were plenty of guides and demos to help me build this application.
For testing, I used the basic test library. Since the app is small, using simple logic statements were enough.
For GET, I tested if I could retrieve the library without changing or modifying the contents of the slice. 
The ID of a book is based on its location in the library, so that is tested in POST and DELETE.
For POST (ADD), I tested if I could add a valid book to the library.
I also tested what would happen if a book had an invalid author, title, or ID.
For DELETE, I tested if I could delete a book with a valid ID. 
I also tested to make sure nothing happened to a delete attempt with an invalid or negative ID.

##Possible Improvements
One possible improvement is adding more features to the application, such as price of a book or quantity. 
For example, functionality can be added to check in and out a book, keeping track of total book count and number of available books.
Another improvement could be to improve the UI. It could be made to interact with the mouse instead of using just the keyboard.
The UI can also be made to look better and maybe run as a website.
