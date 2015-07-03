package main

import (
    "fmt"
    // "html/template"
    "log"
    "net/http"
    // "os"
    // "strings"
)

func init() {
    Config()
    MustConnectRedis()
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/drop", drop)

    log.Println("Start listening...")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        panic(err)
    }
}

func index(res http.ResponseWriter, req *http.Request) {
    defer func() {
        if e := recover(); e != nil {
            log.Println(e)
            res.WriteHeader(http.StatusInternalServerError)
        }
    }()

    log.Println("Index home")

    err := IncNum()
    if err != nil {
        log.Println(err)
    }
    request, err := GetNum()
    if err != nil {
        res.WriteHeader(http.StatusInternalServerError)
        log.Println(err)
        return
    }

    res.Write([]byte(fmt.Sprintf("Request number is %d", request)))
}

func drop(res http.ResponseWriter, req *http.Request) {
    log.Println("drop collection")

    Drop()

    http.Redirect(res, req, "/", 302)
}
