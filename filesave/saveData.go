// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package filesave

import (
    "fmt"
    "os"
)

func SaveData(s string) {

    f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(s)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

}
