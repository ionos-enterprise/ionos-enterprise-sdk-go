package profitbricks

import (
	"log"
	"reflect"
)
type cacheEntry struct {
	item cachable
	depth int
}
type cache map[string]cacheEntry

type cachable interface {
	GetMetadata() *Metadata
}

func newCache() *cache {
	c := make(cache)

	return &c
}
func (c cache) Add(key string, depth int, obj cachable) {
	log.Printf("cachy")
	if cp := reflect.ValueOf(obj).MethodByName("DeepCopy"); ! cp.IsNil() {
		rsp := cp.Call(nil)
		c[key] = cacheEntry{item: rsp[0].Interface().(cachable), depth: depth}
	}
}

func (c cache) Get(key string, depth int) cachable {
	if have, ok := c[key]; ok && have.depth >= depth {
		return have.item
	}
	return nil
}
