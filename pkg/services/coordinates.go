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
    for _, id := range ids {
        coords = append(coords, CoordinatesType{
            ID:   id,
            Lat:  minLat + rand.Float64()*(maxLat-minLat),    //nolint:gosec
            Long: minLong + rand.Float64()*(maxLong-minLong), //nolint:gosec
        })
    }

    return coords, nil
}
