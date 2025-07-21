// This struct holds book information.
type book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
// These vars hold the library data and the file path.
var (
    library     = []book{}
    libraryFile = "library.json"
)
// loadLibrary loads the library from the JSON file.
func loadLibrary() {
    _, err := os.Stat(libraryFile)
    if err == nil {
        data, err := os.ReadFile(libraryFile)
        if err != nil {
            log.Fatalf("Error reading library file! - %v", err)
        }
        json.Unmarshal(data, &library)
    }
}
// saveLibrary saves the library to the JSON file.
func saveLibrary() {
    data, err := json.MarshalIndent(library, "", "  ")
    if err != nil {
        log.Fatalf("Error saving library file! - %v", err)
    }
    os.WriteFile(libraryFile, data, 0644)
}
// deleteBook removes a book from the library by index.
func deleteBook(index int) {
    if index < 0 || index >= len(library) {
        fmt.Println("Invalid book index")
        return
    }
    library = append(library[:index], library[index+1:]...)
    saveLibrary()
}