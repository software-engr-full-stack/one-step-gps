## About

This is a small toy demo app using Go, React, Leaflet and Material UI. When the page is loaded, the app will get coordinates data from a back-end API every 2 seconds. These coordinates are random fake values representing the location of a GPS device attached to an object being tracked. The app shows that last 4 locations of the object. The most current location is represented by the marker with the biggest radius. Clicking on any of the markers will show information about the object being tracked.

## How to build

1. Install Go

2. Install NodeJS

3. Install MySQL

4. Clone this repo and `cd` into this repo

5. Run `./db/reset.sh`

6. Run `./db/user.sh $PASSWORD` - this will create a MySQL user named `go` with the password `$PASSWORD`

7. Run `go mod tidy`

8. Run `cd ui && npm install && npm run build`

9. Run `go build -o tmp/web ./cmd/web`

10. Run `./tmp/web -dbname=one_step_gps -dbuser=go -dbpw=$PASSWORD`

11. Point your browser to `http://localhost:4000`
