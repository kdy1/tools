// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package composite defines an Analyzer that checks for unkeyed
// composite literals.
package composite

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/kdy1/tools/go/analysis"
	"github.com/kdy1/tools/go/analysis/passes/inspect"
	"github.com/kdy1/tools/go/ast/inspector"
	"github.com/kdy1/tools/internal/typeparams"
)

const Doc = `check for unkeyed composite literals

This analyzer reports a diagnostic for composite literals of struct
types imported from another package that do not use the field-keyed
syntax. Such literals are fragile because the addition of a new field
(even if unexported) to the struct will cause compilation to fail.

As an example,

	err = &net.DNSConfigError{err}

should be replaced by:

	err = &net.DNSConfigError{Err: err}
`

var Analyzer = &analysis.Analyzer{
	Name:             "composites",
	Doc:              Doc,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	RunDespiteErrors: true,
	Run:              run,
}

var whitelist = true

func init() {
	Analyzer.Flags.BoolVar(&whitelist, "whitelist", whitelist, "use composite white list; for testing only")
}

// runUnkeyedLiteral checks if a composite literal is a struct literal with
// unkeyed fields.
func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		cl := n.(*ast.CompositeLit)

		typ := pass.TypesInfo.Types[cl].Type
		if typ == nil {
			// cannot determine composite literals' type, skip it
			return
		}
		typeName := typ.String()
		if whitelist && unkeyedLiteral[typeName] {
			// skip whitelisted types
			return
		}
		var structuralTypes []types.Type
		switch typ := typ.(type) {
		case *typeparams.TypeParam:
			terms, err := typeparams.StructuralTerms(typ)
			if err != nil {
				return // invalid type
			}
			for _, term := range terms {
				structuralTypes = append(structuralTypes, term.Type())
			}
		default:
			structuralTypes = append(structuralTypes, typ)
		}
		for _, typ := range structuralTypes {
			under := deref(typ.Underlying())
			if _, ok := under.(*types.Struct); !ok {
				// skip non-struct composite literals
				continue
			}
			if isLocalType(pass, typ) {
				// allow unkeyed locally defined composite literal
				continue
			}

			// check if the CompositeLit contains an unkeyed field
			allKeyValue := true
			for _, e := range cl.Elts {
				if _, ok := e.(*ast.KeyValueExpr); !ok {
					allKeyValue = false
					break
				}
			}
			if allKeyValue {
				// all the composite literal fields are keyed
				continue
			}

			pass.ReportRangef(cl, "%s composite literal uses unkeyed fields", typeName)
			return
		}
	})
	return nil, nil
}

func deref(typ types.Type) types.Type {
	for {
		ptr, ok := typ.(*types.Pointer)
		if !ok {
			break
		}
		typ = ptr.Elem().Underlying()
	}
	return typ
}

func isLocalType(pass *analysis.Pass, typ types.Type) bool {
	switch x := typ.(type) {
	case *types.Struct:
		// struct literals are local types
		return true
	case *types.Pointer:
		return isLocalType(pass, x.Elem())
	case *types.Named:
		// names in package foo are local to foo_test too
		return strings.TrimSuffix(x.Obj().Pkg().Path(), "_test") == strings.TrimSuffix(pass.Pkg.Path(), "_test")
	case *typeparams.TypeParam:
		return strings.TrimSuffix(x.Obj().Pkg().Path(), "_test") == strings.TrimSuffix(pass.Pkg.Path(), "_test")
	}
	return false
}
