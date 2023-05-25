package main


import (
    "fmt"
    "log"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful\n")
	recipeName := r.FormValue("recipeName")
	cuisine := r.FormValue("cuisine")
	ingredients := r.FormValue("ingredients")
	instructions := r.FormValue("instructions")
	source := r.FormValue("source")
	cookTime := r.FormValue("cookTime")

    fmt.Fprintf(w, "Recipe Name = %s\n", recipeName)
    fmt.Fprintf(w, "Cuisine = %s\n", cuisine)
    fmt.Fprintf(w, "Ingredients = %s\n", ingredients)
    fmt.Fprintf(w, "Instructions= %s\n", instructions)
    fmt.Fprintf(w, "Source = %s\n", source)
    fmt.Fprintf(w, "Cook Time (min) = %s\n", cookTime)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}


func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}