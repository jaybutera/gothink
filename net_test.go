package gothink

import (
    "testing"
    "fmt"
)

func TestDecode (t *testing.T) {
    f, _ := NewFFNet("net.json")
    fmt.Println(f)

    x, err := f.ToJson("")
    fmt.Println(err, string(x))
}
