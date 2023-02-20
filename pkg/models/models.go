package models

import (
    "errors"
    "time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Customer struct {
    ID      int
    Name    string
    ImgURL  string
    Created time.Time
    Updated time.Time
}
