// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmdtest

import (
	"encoding/json"
	"testing"

	"github.com/kdy1/tools/internal/lsp/protocol"
	"github.com/kdy1/tools/internal/lsp/tests"
	"github.com/kdy1/tools/internal/span"
)

func (r *runner) Link(t *testing.T, uri span.URI, wantLinks []tests.Link) {
	m, err := r.data.Mapper(uri)
	if err != nil {
		t.Fatal(err)
	}
	out, _ := r.NormalizeGoplsCmd(t, "links", "-json", uri.Filename())
	var got []protocol.DocumentLink
	err = json.Unmarshal([]byte(out), &got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := tests.DiffLinks(m, wantLinks, got); diff != "" {
		t.Error(diff)
	}
}
