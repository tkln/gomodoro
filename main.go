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

    start := info.ModTime()

    for true {
        done := time.Since(start)
        todo := duration - done

        if (done.Nanoseconds() > duration.Nanoseconds()) {
            fmt.Print("done!")
        } else {
            fmt.Print(todo.String())
        }

        d, _ := time.ParseDuration("1s")
        time.Sleep(d)
        fmt.Print("\r")
    }
}
