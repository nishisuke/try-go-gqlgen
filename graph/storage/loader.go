package storage

// import graph gophers with your other imports
import (
	"github.com/graph-gophers/dataloader"
)

type (
	// Loaders wrap your data loaders to inject via middleware
	Loaders struct {
		UserLoader *dataloader.Loader
	}
)

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	// define the data loader
	userReader := &UserReader{}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}
