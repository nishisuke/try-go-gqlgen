package depth

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	errLimit = "DEPTH_LIMIT_EXCEEDED"
	name     = "MaxDepthLimit"
)

type (
	MaxDepthLimit struct {
		limit int
		dfs   DFS
	}
	DFS interface {
		MaxDepth(ast.SelectionSet) int
	}
)

func NewFixedMaxDepthLimit(limit int) *MaxDepthLimit {
	return &MaxDepthLimit{
		limit: limit,
		dfs:   dfs{},
	}
}

func (e MaxDepthLimit) ExtensionName() string {
	return name
}

func (e MaxDepthLimit) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (e MaxDepthLimit) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	op := rc.Doc.Operations.ForName(rc.OperationName)
	depth := e.dfs.MaxDepth(op.SelectionSet)

	if depth > e.limit {
		err := gqlerror.Errorf("operation has maximum depth %d, which exceeds the limit of %d", depth, e.limit)
		errcode.Set(err, errLimit)
		return err
	}

	return nil
}
