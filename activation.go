package gothink

type Sigmoid struct {}

func (s *Sigmoid) Activate (sum float64) float64 {
    return 1.0 / (1.0 + math.Exp(-sum))
}
func (s *Sigmoid) Dactivate (sum float64) float64 {
    return output * (1 - output)
}

type Linear struct {}

func (s *Linear) Activate (sum float64) float64 {
    return sum
}
func (s *Linear) Dactivate (sum float64) float64 {
    return 1.0
}

type Tanh struct {}

func (a *Tanh) Activate(sum float64) float64 {
	return 1.7159 * math.Tanh(2.0/3.0*sum)
}
func (a *Tanh) Dactivate(sum, output float64) float64 {
	return 2.0 / 3.0 * 1.7159 * (1.0 - math.Pow(math.Tanh(2.0/3.0*sum), 2))
}
