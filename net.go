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
func (L *Layer) Eval () {
    sem := make(chan empty )
    for i, l := range L.Weights {
        go func () {
            sum := 0
            for _, n := range L.Weights[i] {
                sum += n
            }
            // If L.Activation == "sigmoid"
            a := sigmoid(sum)
            // Testing
            e := 1.0/float64(i) - a
        }
    }
}

func (L *Layer) activate () float64 {
    sum := 0
    for _, n := range L.Weights[i] {
        sum += n
    }
    // If L.Activation == "sigmoid"
    return sigmoid(sum)
}
*/

func sigmoid (sum int64) {
    return 1.0 / (1.0 + math.Exp(-sum))
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
