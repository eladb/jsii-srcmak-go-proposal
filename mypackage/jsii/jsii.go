// Package jsii contains the functionaility needed for jsii packages to
// initialize their dependencies and themselves. Users should never need to use this package
// directly. If you find you need to - please report a bug at
// https://github.com/aws/jsii/issues/new/choose
package jsii

import (
	_      "embed"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

//go:embed 0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605-0.0.0.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
// The implementation is idempotent (and hence safe to be called over and over).
func Initialize() {
	// Load this library into the kernel
	_jsii_.Load("0d6e4079e36703ebd37c00722f5891d28b0e2811dc114b129215123adcce3605", "0.0.0", tarball)
}
