package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    // "os"
    "github.com/rs/cors"

)
type Response struct {
    Valid bool `json:"valid"`
}

func main() {

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Allows all origins
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })


    http.HandleFunc("/", creditCardValidator)
    handler := c.Handler(http.DefaultServeMux)

    port := "8080"
    fmt.Println("Listening on port:", port)
    err := http.ListenAndServe(":"+port, handler)
    if err != nil {
        fmt.Println("Error:", err)
    }

    // args := os.Args
    // port := args[1]

    // http.HandleFunc("/", creditCardValidator)
    // fmt.Println("Listening on port:", port)
    // err := http.ListenAndServe(":"+port, nil)
    // if err != nil {
    //     fmt.Println("Error:", err)
    // }
}

func creditCardValidator(writer http.ResponseWriter, request *http.Request) {
    // Check if the request method is POST.
    if request.Method != http.MethodPost {
        // if not, throw an error
        http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Create a struct to hold the incoming JSON payload.
    var cardNumber struct {
        Number string `json:"number"` // Number field holds the credit card number.
    }

    // Decode the JSON payload from the request body into the cardNumber struct.
    err := json.NewDecoder(request.Body).Decode(&cardNumber)
    if err != nil {
        http.Error(writer, "Invalid JSON payload", http.StatusBadRequest)
        return
    }
    // Validate the credit card number using the Luhn algorithm.
    isValid := luhnAlgorithm(cardNumber.Number)
    // Create a response struct with the validation result.
    response := Response{Valid: isValid}

    // Marshal the response struct into JSON format.
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(writer, "Error creating response", http.StatusInternalServerError)
        return
    }

    // Set the content type header to indicate JSON response.
    writer.Header().Set("Content-Type", "application/json")

    // Write the JSON response back to the client.
    writer.Write(jsonResponse)
}

