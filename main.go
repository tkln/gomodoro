// Copyright (C) 2019 Aapo Vienamo
// SPDX-License-Identifier: CC0-1.0

package main

import (
    "fmt"
    "os"
    "log"
    "path/filepath"
    "time"
)

func main() {
    session_file := ".pomodoro_session"
    home := os.Getenv("HOME")
    path := filepath.Join(home, session_file)
    file, err := os.Open(path)
    if (err != nil) {
        log.Fatal(err)
        return
    }

    info, err := file.Stat()
    if (err != nil) {
        log.Fatal(err)
        return
    }

    duration, err := time.ParseDuration("20m")
    if (err != nil) {
        log.Fatal(err)
        return
    }

    break_duration, err := time.ParseDuration("5m")
    if (err != nil) {
        log.Fatal(err)
        return
    }

    start := info.ModTime()

    for true {
        done := time.Since(start)
        todo := duration - done
        break_remain := break_duration - (done - duration)

        if (done > duration) {
            fmt.Print("break ")
            fmt.Print(break_remain.String())
        } else {
            fmt.Print(todo.String())
        }

        d, _ := time.ParseDuration("1s")
        time.Sleep(d)
        fmt.Print("\r")
    }
}
