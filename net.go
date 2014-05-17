package gothink

import (
    "encoding/json"
    "io/ioutil"
)

type Net interface {
    Epoch()
}

type Layer struct {
    Activator
    Neurons     map[uint32]Neuron
}

func (L *Layer) Eval () {
    sem := make(chan float64, len(L))

    for i, n := range n.Weights {
        go func () {
            n.Eval()
            delta := n.Val * (1 - n.Val) * (t.i - n.Val)
        }()
    }
}

/*
A feed forward type neural network
*/
type FFNet struct {
    //Net
    Input []float64
    Layers []Layer
}

func (nn *FFNet) Update () {

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
