// package main

// import (
//     "encoding/json"
//     "fmt"
//     "net/http"
// )

// type Response struct {
//     Valid bool `json:"valid"`
// }

// func main() {
//     http.HandleFunc("/", creditCardValidator)
//     fmt.Println("Listening on port: 8080")
//     err := http.ListenAndServe(":8080", nil)
//     if err != nil {
//         fmt.Println("Error:", err)
//     }
// }

// func creditCardValidator(writer http.ResponseWriter, request *http.Request) {
//     writer.Header().Set("Access-Control-Allow-Origin", "*")
//     writer.Header().Set("Access-Control-Allow-Methods", "POST")
//     writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

//     if request.Method == http.MethodOptions {
//         return
//     }

//     if request.Method != http.MethodPost {
//         http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
//         return
//     }

//     err := request.ParseForm()
//     if err != nil {
//         http.Error(writer, "Error parsing form", http.StatusBadRequest)
//         return
//     }

//     cardNumber := request.FormValue("number")
//     fmt.Println("Card Number Received:", cardNumber)
//     if cardNumber == "" {
//         http.Error(writer, "Missing card number", http.StatusBadRequest)
//         return
//     }

//     isValid := luhnAlgorithm(cardNumber)
//     response := Response{Valid: isValid}

//     jsonResponse, err := json.Marshal(response)
//     if err != nil {
//         http.Error(writer, "Error creating response", http.StatusInternalServerError)
//         return
//     }

//     writer.Header().Set("Content-Type", "application/json")
//     writer.Write(jsonResponse)
// }


package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/rs/cors"

)

type Response struct {
    Valid bool `json:"valid"`
}

func main() {
    // http.HandleFunc("/", creditCardValidator)
    // fmt.Println("Listening on port: 8080")
    // err := http.ListenAndServe(":8080", nil)
    // if err != nil {
    //     fmt.Println("Error:", err)
    // }
    
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
}

func creditCardValidator(writer http.ResponseWriter, request *http.Request) {
    // writer.Header().Set("Access-Control-Allow-Origin", "*")
    // writer.Header().Set("Access-Control-Allow-Methods", "POST")
    // writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if request.Method == http.MethodOptions {
        return
    }

    if request.Method != http.MethodPost {
        http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Log request headers and content type for debugging
    fmt.Println("Request Headers:", request.Header)
    fmt.Println("Content-Type:", request.Header.Get("Content-Type"))

    err := request.ParseMultipartForm(32 << 20) // 32 MB limit
    if err != nil {
        http.Error(writer, "Error parsing form", http.StatusBadRequest)
        return
    }

    cardNumber := request.FormValue("number")
    fmt.Println("Card Number Received:", cardNumber) // Log the received card number
    if cardNumber == "" {
        http.Error(writer, "Missing card number", http.StatusBadRequest)
        return
    }

    isValid := luhnAlgorithm(cardNumber)
    response := Response{Valid: isValid}

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(writer, "Error creating response", http.StatusInternalServerError)
        return
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.Write(jsonResponse)
}
