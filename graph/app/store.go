package app

import (
	"context"
	"example/graph/loader"

	"gorm.io/gorm"
)

const (
	key   = "loader-key"
	dbkey = "gorm-key"
)

func StoreLoader(ctx context.Context, con *gorm.DB) context.Context {
	return context.WithValue(ctx, key, loader.NewLoader(con))
}

func StoreDB(ctx context.Context, con *gorm.DB) context.Context {
	return context.WithValue(ctx, dbkey, con.WithContext(ctx).Debug())
}
