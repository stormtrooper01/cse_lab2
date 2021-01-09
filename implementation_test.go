package lab2

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPrefixComputeCorrect(t *testing.T) {
    res, err := ComputePrefix("7")
    if assert.Nil(t, err) {
        assert.Equal(t, 7.0, res)
    }

    res, err = ComputePrefix("+ 4 * - 4 3 2")
    if assert.Nil(t, err) {
        assert.Equal(t, 6.0, res)
    }

    res, err = ComputePrefix("* 10.25 / 100 ^ - 5 6 3")
    if assert.Nil(t, err) {
        assert.Equal(t, -1025.0, res)
    }
// 3 × 10 + 1 × 15 + 2 × 40 + 6 × 20 + 4 × 30 = 365
    res, err = ComputePrefix("+ * 3 10 + * 1 15 + * 2 40 + * 6 20 * 4 30")
    if assert.Nil(t, err) {
        assert.Equal(t, 365.0, res)
    }
}

func TestPrefixComputeIncorrect(t *testing.T) {
    _, err := ComputePrefix("+")
    assert.NotNil(t, err)

    _, err = ComputePrefix("+ 5")
    assert.NotNil(t, err)

    _, err = ComputePrefix("5 +")
    assert.NotNil(t, err)

    _, err = ComputePrefix("")
    assert.NotNil(t, err)

    _, err = ComputePrefix("aaa")
    assert.NotNil(t, err)
}

func ExampleComputePrefix() {
    res, _ := ComputePrefix("+ 2 2")
    fmt.Println(res)

    // Output:
    // 4
}
