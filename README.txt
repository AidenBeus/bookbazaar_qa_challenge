##Setup

## Clone the project

```
$ git clone https://github.com/AidenBeus/bookbazaar_qa_challenge.git
$ cd bookbazaar_qa_challenge
```
## Get tview 
In order to use tview, download the github library.
```
$go mod init 
&go get github.com/rivo/tview@master
``` 

##Running the project
To run the application, simple run:
```
$cd ui
$go run .
```
From there, a window will pop up. Use the TAB key to navigate the fields.
To add a book, input book title and author name.. Then hit ENTER on the "Add Book" button.
To delete a book, input just the book id. Then hit ENTER on the "Delete Book" button.
To exit the app, hit ENTER on the "Quit" button.

##Running the testcases
Use the test command to run the test cases.
```
$go test -v
```
Each test case should show whether it passed, along with the runtime.

##Rationale
I choose to use tview for my application UI because it is a relatively simple UI library. 
I liked how it looked and there were plenty of guides and demos to help me build this application.
For testing, I used the basic test library. Since the app is small, using logic statements were enough.
