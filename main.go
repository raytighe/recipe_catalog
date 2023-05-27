package main


import (
    "fmt"
    "log"
    "net/http"
    "recipe_catalog/dynamodb"
    "strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    recipeId, _ := strconv.Atoi(r.FormValue("recipeId"))
    recipeName := r.FormValue("recipeName")
    cuisine := r.FormValue("cuisine")
    ingredients := r.FormValue("ingredients")
    instructions := r.FormValue("instructions")
    source := r.FormValue("source")
    cookTime, _ := strconv.Atoi(r.FormValue("cookTime"))

    entry :=  table_operations.Recipe {
        RecipeId: recipeId,
        RecipeName: recipeName,
        Cuisine: cuisine,
        Ingredients: ingredients,
        Instructions: instructions,
        Source: source,
        CookTime: cookTime, 
        }

    // Query Dynamodb and display results
    if r.Method == "GET" {
        fmt.Fprintf(w, "Search Results\n-------------------\n",)
        response := table_operations.ScanItems(entry)
        for _, r := range response {
        fmt.Fprintf(w, "Recipe ID: %v\n", r.RecipeId)
        fmt.Fprintf(w, "Recipe Name: %s\n", r.RecipeName)
        fmt.Fprintf(w, "Cuisine: %s\n", r.Cuisine)
        fmt.Fprintf(w, "Ingredients:\n %s\n", r.Ingredients)
        fmt.Fprintf(w, "Cooking Instructions:\n %s\n", r.Instructions)
        fmt.Fprintf(w, "Source: %s\n", r.Source)
        fmt.Fprintf(w, "Cook Time (min): %v\n", r.CookTime)
        fmt.Fprintf(w, "-------------------\n",)
        }
        fmt.Fprintf(w, "\n%v recipes returned", len(response))


    // Insert record in Dynamodb
    } else if r.Method == "POST" {
        table_operations.WriteItem(entry)
        fmt.Fprintf(w, "POST request successful\n")
        fmt.Fprintf(w, "Recipe ID = %v\n", recipeId)
        fmt.Fprintf(w, "Recipe Name = %s\n", recipeName)
        fmt.Fprintf(w, "Cuisine = %s\n", cuisine)
        fmt.Fprintf(w, "Ingredients = %s\n", ingredients)
        fmt.Fprintf(w, "Cooking Instructions = %s\n", instructions)
        fmt.Fprintf(w, "Source = %s\n", source)
        fmt.Fprintf(w, "Cook Time (min) = %v\n", cookTime)

    } else {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
    }
}

func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/search", formHandler)
    // http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}