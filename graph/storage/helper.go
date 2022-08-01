package storage

// import graph gophers with your other imports
import (
	"errors"

	"github.com/graph-gophers/dataloader"
)

func conv(keys dataloader.Keys) []string {
	ids := make([]string, len(keys))
	for i, k := range keys {
		ids[i] = k.String()
	}
	return ids
}

func packResult[T any](keys dataloader.Keys, look map[string]T) []*dataloader.Result {

	res := make([]*dataloader.Result, len(keys))

	for i, key := range keys {
		if u, ok := look[key.String()]; ok {
			res[i] = &dataloader.Result{Data: u}
		} else {
			res[i] = &dataloader.Result{Error: errors.New("NotFound")}
		}
	}

	return res
}
