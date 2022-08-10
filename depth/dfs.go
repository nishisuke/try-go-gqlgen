package gqlgendepth

import (
	"github.com/vektah/gqlparser/v2/ast"
)

type (
	dfs struct{}
)

func (g dfs) MaxDepth(set ast.SelectionSet) int {
	return maxDepth(set) - 1
}

func maxDepth(set ast.SelectionSet) int {
	maxChildrenDepth := 0
	for _, selection := range set {
		depth := walk(selection)
		if depth > maxChildrenDepth {
			maxChildrenDepth = depth
		}
	}
	return 1 + maxChildrenDepth
}

func walk(selection ast.Selection) int {
	switch typed := selection.(type) {
	case *ast.Field:
		return maxDepth(typed.SelectionSet)
		//case *ast.FragmentSpread:
		//	complexity = safeAdd(complexity, cw.recursiveWalk(s.Definition.SelectionSet))

		//case *ast.InlineFragment:
		//	complexity = safeAdd(complexity, cw.recursiveWalk(s.SelectionSet))
	default:
		panic("Not implemented")
	}
}
