package main

import (
    "testing"
)

func Test_Redis(t *testing.T) {
    _IncNum(t)
    _GetNum(t)
    _Drop(t)
}

func _GetNum(t *testing.T) {
    if _, err := GetNum(); err != nil {
        t.Errorf("get num meets error %v\n", err)
    }
}

func _IncNum(t *testing.T) {
    if err := IncNum(); err != nil {
        t.Errorf("inc num meets error %v\n", err)
    }
}

func _Drop(t *testing.T) {
    if err := Drop(); err != nil {
        t.Errorf("drop request meets error %v\n", err)
    }
}
