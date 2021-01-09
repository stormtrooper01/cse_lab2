package lab2

import (
    "fmt"
    "io"
    "io/ioutil"
)

type ComputeHandler struct {
    In  io.Reader
    Out io.Writer
}

func (ch *ComputeHandler) Compute() error {
    expr, err := ioutil.ReadAll(ch.In)
    if err != nil {
        return err
    }

    if res, err := ComputePrefix(string(expr)); err != nil {
        return err
    } else {
        buf := []byte(fmt.Sprintf("%f", res))
        if _, e := ch.Out.Write(buf); e != nil {
            return e
        }
    }

    return nil
}
