package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func (app *application) showResource(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    fmt.Fprintf(w, "Display a specific resource with ID %d...", id)
}

func (app *application) createResource(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    _, err := w.Write([]byte("Create a new resource..."))
    if err != nil {
        app.errorLog.Println(err)
    }
}
