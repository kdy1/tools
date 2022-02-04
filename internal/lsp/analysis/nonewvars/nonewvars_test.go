// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nonewvars_test

import (
	"testing"

	"github.com/kdy1/tools/go/analysis/analysistest"
	"github.com/kdy1/tools/internal/lsp/analysis/nonewvars"
	"github.com/kdy1/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"a"}
	if typeparams.Enabled {
		tests = append(tests, "typeparams")
	}
	analysistest.RunWithSuggestedFixes(t, testdata, nonewvars.Analyzer, tests...)
}
