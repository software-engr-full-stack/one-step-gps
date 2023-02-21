package services

import (
    "fmt"
    "math/rand"
)

type CoordinatesType struct {
    ID   int     `json:"id"`
    Lat  float64 `json:"lat"`
    Long float64 `json:"long"`
}

const (
    // Min/max lat/long forms a rectangular box in the LA area
    minLat = 33.780486
    maxLat = 34.162326

    minLong = -118.331057
    maxLong = -117.291232
)

func Coordinates(ids []int) ([]CoordinatesType, error) {
    if len(ids) < 1 {
        return nil, fmt.Errorf("must pass at least one id")
    }

    coords := []CoordinatesType{}
    for ix, id := range ids {
        // First id is always updated, the rest are updated randomly.
        if ix != 0 {
            if rand.Int()%(ix+3) == 0 { //nolint:gosec
                continue
            }
        }
        coords = append(coords, CoordinatesType{
            ID:   id,
            Lat:  minLat + rand.Float64()*(maxLat-minLat),    //nolint:gosec
            Long: minLong + rand.Float64()*(maxLong-minLong), //nolint:gosec
        })
    }
    return coords, nil
}
