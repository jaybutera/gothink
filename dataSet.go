package gothink

import (
    "encoding/json"
    "io/ioutil"
)

type DataSet struct {
    Inputs, Outputs [][]float64
}

func (d *DataSet) Inputs (i int) []float64{
    return d.inputs[i]
}

func (d *DataSet) Outputs (i int) []float64{
    return d.outputs[i]
}

func NewDataSet (filepath string) (*DataSet, error) {
    b, err := ioutil.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    d := DataSet{}
    return &d, json.Unmarshal(b, &d)
}
