package gothink

import (
    "encoding/json"
    "io/ioutil"
)

type Net interface {
    Epoch()
}

type Layer struct {
    Activation  string
    Weights     [][]float64
}

/*
A feed forward type neural network
*/
type FFNet struct {
    //Net
    Layers []Layer
}

func NewFFNet (filepath string) (*FFNet, error) {
    b, err := ioutil.ReadFile(filepath)

    if err != nil {
        panic(err)
    }

    ff := FFNet{}
    return &ff, json.Unmarshal(b, &ff)
}

func (ff *FFNet) ToJson (filepath string) ([]byte, error) {
    if filepath != "" {
        d, err := json.Marshal(&ff)

        if err != nil{
            panic(err)
        }

        err1 := ioutil.WriteFile(filepath, d, 0644)

        return d, err1
    }
    return json.Marshal(&ff)
}

/*
func EncFFNet () ([]byte, error) {
    f := FFNet{}

    f.Layers = make(map[string]interface{})

    f.Layers["one"] = []float64{.5, .2}
    f.Layers["two"] = []float64{.0, .1}
    return json.Marshal(f)
}
*/
