package fileread

import (
    "io"
    "os"
)

func ReadData() string {
    // open input file
    fi, err := os.Open("test.txt")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // open output string
    contents := ""

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        contents = contents + string(buf[:n])
    }

    return contents;
}