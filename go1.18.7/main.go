// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.18.7 command runs the go command from Go 1.18.7.
//
// To install, run:
//
//	$ go install github.com/SunJary/dl/go1.18.7@latest
//	$ go1.18.7 download
//
// And then use the go1.18.7 command as if it were your normal go
// command.
//
// See the release notes at https://go.dev/doc/devel/release#go1.18.minor.
//
// File bugs at https://go.dev/issue/new.
package main

import "github.com/SunJary/dl/internal/version"

func main() {
	version.Run("go1.18.7")
}