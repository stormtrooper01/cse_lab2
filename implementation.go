package lab2

import (
    "errors"
    "math"
    "strconv"
    "strings"
)

type operation struct {
    left   expression
    right  expression
    opType int
}

type expression interface {
    Compute() float64
}

type operand struct {
    value float64
}

const (
    MUL = iota
    DIV = iota
    ADD = iota
    SUB = iota
    POW = iota
)

func (o operand) Compute() float64 { return o.value }

func parseOp(opType int, exp []string) (expression, error, []string) {
    left, err1, rest1 := parse(exp)
    if err1 != nil {
        return nil, err1, nil
    } else if len(rest1) <= 0 {
        return nil, errors.New("syntax error"), nil
    }

    right, err2, rest2 := parse(rest1)
    if err2 != nil {
        return nil, err2, nil
    }

    return operation{
        left:   left,
        right:  right,
        opType: opType,
    }, nil, rest2
}

func parse(input[] string) (expression, error, []string) {
    if len(input) <= 0 {
        return nil, errors.New("syntax error"), nil
    }

    switch input[0] {
    case "*":
        return parseOp(MUL, input[1:])
    case "/":
        return parseOp(DIV, input[1:])
    case "+":
        return parseOp(ADD, input[1:])
    case "-":
        return parseOp(SUB, input[1:])
    case "^":
        return parseOp(POW, input[1:])
    default:
        num, err := strconv.ParseFloat(input[0], 64)
        if err != nil {
            return nil, err, nil
        }
        return operand{value: num}, nil, input[1:]
    }
}

func (o operation) Compute() float64 {
    res := 0.0
    left := o.left.Compute()
    right := o.right.Compute()

    switch o.opType {
    case ADD:
        res = left + right
    case SUB:
        res = left - right
    case MUL:
        res = left * right
    case DIV:
        res = left / right
    case POW:
        res = math.Pow(left, right)
    }

    return res
}

func ComputePrefix(input string) (float64, error) {
    expression, err, rest := parse(strings.Split(input, " "))
    if len(rest) > 0 {
        return 0, errors.New("redundant characters")
    } else if err != nil {
        return 0, err
    } else {
        return expression.Compute(), nil
    }
}
