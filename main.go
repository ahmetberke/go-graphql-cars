package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

func init() {
	err := getCars()
	if err != nil {
		panic(err)
	}
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), carSchema)
		json.NewEncoder(w).Encode(result)
	})
	// Serve static files
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single car: curl -g 'http://localhost:8080/graphql?query={car(racing_number:\"b\"){name,speed,racing_number,country,racing_type}}'")
	fmt.Println("Load car list: curl -g 'http://localhost:8080/graphql?query={carList{name,speed,racing_number,country,racing_type}}'")

	http.ListenAndServe(":8080", nil)
}
