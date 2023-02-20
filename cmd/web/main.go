package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "database/sql"

    _ "github.com/go-sql-driver/mysql"

    "github.com/software-engr-full-stack/one-step-gps/pkg/models/mysql"
)

type application struct {
    errorLog  *log.Logger
    infoLog   *log.Logger
    customers *mysql.CustomerModel
}

type argsType struct {
    addr   string
    dbname string
    dbuser string
    dbpw   string
}

func main() {
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    args := &argsType{}
    err := args.validate()
    if err != nil {
        errorLog.Fatal(err)
    }

    db, err := openDB(fmt.Sprintf("%s:%s@/%s?parseTime=true", args.dbuser, args.dbpw, args.dbname))
    if err != nil {
        errorLog.Fatal(err)
    }
    defer db.Close()

    cm := &mysql.CustomerModel{DB: db}
    if err := cm.Validate(); err != nil {
        errorLog.Fatal(err)
    }
    app := &application{
        errorLog:  errorLog,
        infoLog:   infoLog,
        customers: &mysql.CustomerModel{DB: db},
    }

    server := &http.Server{
        Addr:              args.addr,
        ErrorLog:          errorLog,
        Handler:           app.routes(),
        ReadHeaderTimeout: 5 * time.Second,
    }

    infoLog.Printf("Starting server on %s", args.addr)

    if err := server.ListenAndServe(); err != nil {
        errorLog.Print(err)
        return
    }
}

func (args *argsType) validate() error {
    flag.StringVar(&args.addr, "addr", ":4000", "HTTP network address")
    args.addr = strings.TrimSpace(args.addr)

    flag.StringVar(&args.dbname, "dbname", "", "database name")
    args.dbname = strings.TrimSpace(args.dbname)

    flag.StringVar(&args.dbuser, "dbuser", "", "database username")
    args.dbuser = strings.TrimSpace(args.dbuser)

    flag.StringVar(&args.dbpw, "dbpw", "", "database password")
    args.dbpw = strings.TrimSpace(args.dbpw)

    flag.Parse()

    if args.dbname == "" {
        return fmt.Errorf("must pass valid dbname")
    }

    if args.dbuser == "" {
        return fmt.Errorf("must pass valid dbuser")
    }

    if args.dbpw == "" {
        return fmt.Errorf("must pass valid dbpw")
    }

    return nil
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
