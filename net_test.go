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

func TestDataSet (t *testing.T) {
    d, err := NewDataSet("testData.json")

    fmt.Println(err, d)
}

/*
func TestEval (t *testing.T) {
    f, _ := NewFFNet("net.json")
    d, _ := NewDataSet("testData.json")

    fmt.Println(f.Weights[0])
    f.Weights[0].Eval()
    fmt.Println(f.Weights[0])
}
*/
