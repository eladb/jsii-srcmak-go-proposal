package mypackage

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605.Calculator",
		reflect.TypeOf((*Calculator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "mul", GoMethod: "Mul"},
			_jsii_.MemberMethod{JsiiMethod: "sub", GoMethod: "Sub"},
		},
		func() interface{} {
			return &jsiiProxy_Calculator{}
		},
	)
	_jsii_.RegisterStruct(
		"0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605.Operands",
		reflect.TypeOf((*Operands)(nil)).Elem(),
	)
}
