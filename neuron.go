package gothink

type Neuron struct {
    Val        float64
    Weights    []float64
    connected  []uint32
}

func (n *Neuron) Eval (inp *Layer) {
    sum := 0

    for i, ni := range inp.Neurons {
        sum += ni.Val * n.Weights[i]
    }

    n.Val = sum
}
