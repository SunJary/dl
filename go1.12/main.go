// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.12 command runs the go command from Go 1.12.
//
// To install, run:
//
//     $ go install github.com/SunJary/dl/go1.12@latest
//     $ go1.12 download
//
// And then use the go1.12 command as if it were your normal go
// command.
//
// See the release notes at https://github.com/SunJary/doc/go1.12
//
// File bugs at https://github.com/SunJary/issues/new
package main

import "github.com/SunJary/dl/internal/version"

func main() {
	version.Run("go1.12")
}
