package gothink

import (
    "encoding/json"
    "io/ioutil"
)

type Net interface {
    Epoch()
}

type Vector interface {
    Activation() *Vector
}

type InputLayer struct {
    Vector
    V []float64
}

/*
A feed forward type neural network
*/
type FFNet struct {
    Net
    Layers map[string]interface{}
}

func NewFFNet (filePath string) *FFNet {
    f := FFNet{}
    f.Layers = make(map[string]interface{})
    enc, err := ioutil.ReadFile(filePath)

    if err != nil {
        panic("failed to load" + filePath)
    }

    var objmap map[string]*json.RawMessage
    json.Unmarshal(enc, &objmap)

    //json.Unmarshal(*objmap["Layers"], &f.Layers)
    return &f


func EncFFNet () ([]byte, error) {
    f := FFNet{}

    f.Layers = make(map[string]interface{})

    f.Layers["one"] = []float64{.5, .2}
    f.Layers["two"] = []float64{.0, .1}
    return json.Marshal(f)
}
