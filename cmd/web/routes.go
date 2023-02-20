package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/customer", app.showCustomer)
    mux.HandleFunc("/customers/latest", app.showCustomersLatest)
    mux.HandleFunc("/customers/coords", app.showCustomersCoords)

    fileServer := http.FileServer(http.Dir("./ui/build/"))
    mux.Handle("/", fileServer)

    return mux
}
