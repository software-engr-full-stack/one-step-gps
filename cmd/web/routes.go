package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/resource", app.showResource)
    mux.HandleFunc("/resource/create", app.createResource)

    fileServer := http.FileServer(http.Dir("./ui/build/"))
    mux.Handle("/", fileServer)

    return mux
}
