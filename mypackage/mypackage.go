// 0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605
package mypackage

import (
	_init_ "com.example/mymodule/mypackage/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"
)

// A sophisticaed multi-language calculator.
type Calculator interface {
	Add(ops *Operands) *float64
	Mul(ops *Operands) *float64
	Sub(ops *Operands) *float64
}

// The jsii proxy struct for Calculator
type jsiiProxy_Calculator struct {
	_ byte // padding
}

func NewCalculator() Calculator {
	_init_.Initialize()

	j := jsiiProxy_Calculator{}

	_jsii_.Create(
		"0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605.Calculator",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewCalculator_Override(c Calculator) {
	_init_.Initialize()

	_jsii_.Create(
		"0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605.Calculator",
		nil, // no parameters
		c,
	)
}

// Adds the two operands.
func (c *jsiiProxy_Calculator) Add(ops *Operands) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"add",
		[]interface{}{ops},
		&returns,
	)

	return returns
}

// Multiplies the two operands.
func (c *jsiiProxy_Calculator) Mul(ops *Operands) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"mul",
		[]interface{}{ops},
		&returns,
	)

	return returns
}

// Subtracts the two operands.
func (c *jsiiProxy_Calculator) Sub(ops *Operands) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"sub",
		[]interface{}{ops},
		&returns,
	)

	return returns
}

// Math operands.
type Operands struct {
	// Left-hand side operand.
	Lhs *float64 `json:"lhs"`
	// Right-hand side operand.
	Rhs *float64 `json:"rhs"`
}
