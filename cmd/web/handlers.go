package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func (app *application) showCustomer(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    fmt.Fprintf(w, "Display a specific customer with ID %d...", id)
}

func (app *application) showCustomers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Display all customers...")
}

// func (app *application) createCustomer(w http.ResponseWriter, r *http.Request) {
//     if r.Method != "POST" {
//         w.Header().Set("Allow", "POST")
//         app.clientError(w, http.StatusMethodNotAllowed)
//         return
//     }
//     _, err := w.Write([]byte("Create a new customer..."))
//     if err != nil {
//         app.errorLog.Println(err)
//     }
// }
