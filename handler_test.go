package lab2

import (
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

type outputTest struct {
    called bool
}

func (o *outputTest) Write(p []byte) (n int, err error) {
    o.called = true
    return 0, nil
}

func TestHandlCompute(t *testing.T) {
    exp := "+ 5 * - 4 2 3"
    in := strings.NewReader(exp)
    out := outputTest{}

    handler := ComputeHandler{
        In: in,
        Out: &out,
    }
    handler.Compute()

    assert.Equal(t,true, out.called)
}

func TestHandlComputeIncorrect(t *testing.T) {
    exp := " + 5 * - 4 2 3"
    in := strings.NewReader(exp)
    out := outputTest{}

    handler := ComputeHandler{
        In: in,
        Out: &out,
    }
    err := handler.Compute()
    assert.NotNil(t, err)
}
