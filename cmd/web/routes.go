package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/v1/customer", app.showCustomer)
    mux.HandleFunc("/api/v1/customers/latest", app.showCustomersLatest)
    mux.HandleFunc("/api/v1/customers/coords", app.showCustomersCoords)

    fileServer := http.FileServer(http.Dir("./ui/build/"))
    mux.Handle("/", fileServer)

    return mux
}
