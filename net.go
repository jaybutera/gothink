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
    //Vector
    ID          string
    Activation  string
    Weights     [][]float64
}

/*
A feed forward type neural network
*/
type FFNet struct {
    //Net
    //Layers []InputLayer // Change to interface{}
    Layers map[string]InputLayer // Change to interface{}
}

/*
func NewFFNet (filePath string) *FFNet {
    f := FFNet{}
    //f.Layers = make([]InputLayer, 1) // Change to interface{}
    _, err := ioutil.ReadFile(filePath)

    if err != nil {
        panic("failed to load" + filePath)
    }

    return (&f)
}
*/

func FromJson (filepath string) (*FFNet, error) {
    b, _ := ioutil.ReadFile(filepath)

    ff := FFNet{}
    return &ff, json.Unmarshal(b, &ff)
}

/*
func (ff *FFNet) ToJson (jsonStr string) error {
    var data = &ff.Layers
}
*/

/*
func EncFFNet () ([]byte, error) {
    f := FFNet{}

    f.Layers = make(map[string]interface{})

    f.Layers["one"] = []float64{.5, .2}
    f.Layers["two"] = []float64{.0, .1}
    return json.Marshal(f)
}
*/
