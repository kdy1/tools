// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The stringintconv command runs the stringintconv analyzer.
package main

import (
	"github.com/kdy1/tools/go/analysis/passes/stringintconv"
	"github.com/kdy1/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(stringintconv.Analyzer) }
