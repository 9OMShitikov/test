package bitmapFileIndex

import (
	"github.com/neganovalexey/bitmap-search"
)

type FileIndex struct {
	index *roaringIndex.Index
}
