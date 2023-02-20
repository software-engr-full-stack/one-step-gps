package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/customer", app.showCustomer)
    mux.HandleFunc("/customers", app.showCustomers)

    // mux.HandleFunc("/customer/create", app.createCustomer)

    fileServer := http.FileServer(http.Dir("./ui/build/"))
    mux.Handle("/", fileServer)

    return mux
}
