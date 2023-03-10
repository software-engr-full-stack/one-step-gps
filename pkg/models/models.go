package models

import (
    "errors"
    "time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Customer struct {
    ID               int       `json:"id"`
    Name             string    `json:"name"`
    BusinessCategory string    `json:"businessCategory"`
    Payload          string    `json:"payload"`
    Color            string    `json:"color"`
    ImgURL           string    `json:"imgURL"`
    Created          time.Time `json:"created"`
    Updated          time.Time `json:"updated"`
}
