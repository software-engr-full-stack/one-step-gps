package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/software-engr-full-stack/one-step-gps/pkg/models"
    "github.com/software-engr-full-stack/one-step-gps/pkg/services"
)

func (app *application) showCustomer(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    cm, err := app.customers.Get(id)
    if err == models.ErrNoRecord {
        app.notFound(w)
        return
    } else if err != nil {
        app.serverError(w, err)
        return
    }

    enc, err := json.Marshal(cm)
    if err != nil {
        app.serverError(w, err)
        return
    }

    fmt.Fprint(w, string(enc))
}

func (app *application) showCustomersLatest(w http.ResponseWriter, r *http.Request) {
    customers, err := app.customers.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    enc, err := json.Marshal(customers)
    if err != nil {
        app.serverError(w, err)
        return
    }

    fmt.Fprint(w, string(enc))
}

// Param format: ids=[1, 2, 3, ...]
func (app *application) showCustomersCoords(w http.ResponseWriter, r *http.Request) {
    idsStr := r.URL.Query().Get("ids")
    var ids []int
    if err := json.Unmarshal([]byte(idsStr), &ids); err != nil {
        app.serverError(w, err)
        return
    }

    coords, err := services.Coordinates(ids)
    if err != nil {
        app.serverError(w, err)
        return
    }

    enc, err := json.Marshal(coords)
    if err != nil {
        app.serverError(w, err)
        return
    }

    fmt.Fprint(w, string(enc))
}
