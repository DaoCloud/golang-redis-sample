package main

import (
    "os"

    "github.com/garyburd/redigo/redis"
)

var (
    address  string
    port     string
    password string
    proto    string
)

func Config() {
    proto = os.Getenv("REDIS_PORT_6379_TCP_PROTO")
    address = os.Getenv("REDIS_PORT_6379_TCP_ADDR")
    port = os.Getenv("REDIS_PORT_6379_TCP_PORT")
    password = os.Getenv("REDIS_PASSWORD")
}

var conn redis.Conn

func ConnectRedis() (err error) {
    conn, err = redis.Dial(proto, address+":"+port)
    if err != nil {
        return err
    }

    if password != "" {
        if _, err = conn.Do("AUTH", password); err != nil {
            conn.Close()
            return err
        }
    }

    return nil
}

func MustConnectRedis() {
    if err := ConnectRedis(); err != nil {
        panic(err)
    }
}

func GetNum() (request int, err error) {

    request, err = redis.Int(conn.Do("GET", "request"))

    return request, err
}

func IncNum() error {
    _, err := conn.Do("INCR", "request")
    return err
}

func Drop() error {
    _, err := conn.Do("DEL", "request")
    return err
}
