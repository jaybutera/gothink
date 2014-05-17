package gothink

func bp (N *FFNet, filepath string) {
    d, err := NewDataSet(filepath)

    if err != nil {
        panic(err)
    }

    for _, L := range N.Layers {
    }
}
