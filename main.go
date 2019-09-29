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

    start := info.ModTime()

    fmt.Println(time.Since(start).String())
}
