package collections

import (
	"fmt"
)

type DefaultItems[ID IDAllowedTypes, V ValueAllowedTypes] struct {
	internalMap   map[ID]*V
	internalArray []Item[ID, V]
}

func NewDefaultItems[ID IDAllowedTypes, V ValueAllowedTypes](items ...Item[ID, V]) *DefaultItems[ID, V] {
	internalMap := ToMap(items...)
	return &DefaultItems[ID, V]{
		internalMap:   internalMap,
		internalArray: items,
	}
}

func (items *DefaultItems[ID, V]) IDs() []ID {
	return Keys(items.internalMap)
}

func (items *DefaultItems[ID, V]) Values() []*V {
	return Values(items.internalMap)
}

func (items *DefaultItems[ID, V]) Contains(id ID) bool {
	_, ok := items.internalMap[id]
	return ok
}

func (items *DefaultItems[ID, V]) Find(id ID) (*V, error) {
	model, ok := items.internalMap[id]
	if !ok {
		return nil, fmt.Errorf("item %v not found", id)
	}
	return model, nil
}

func (items *DefaultItems[ID, V]) Filter(filterFunc FilterFunc[ID, V]) Items[ID, V] {
	result := Filter(items.internalArray, func(item Item[ID, V], index int) bool {
		return filterFunc(index, *item.ID(), *item.Value())
	})
	return NewDefaultItems(result...)
}

func (items *DefaultItems[ID, V]) Count(countFunc CountFunc[ID, V]) int {
	return Count(items.internalArray, func(item Item[ID, V], index int) bool {
		return countFunc(index, *item.ID(), *item.Value())
	})
}

func (items *DefaultItems[ID, V]) Select(ids []ID) []*V {
	result := make([]*V, 0, len(ids))
	for _, id := range ids {
		value, _ := items.Find(id)
		result = append(result, value)
	}
	return result
}

func (items *DefaultItems[ID, V]) ForEach(iterateeFunc IterateeFunc[ID, V]) {
	for index, item := range items.internalArray {
		iterateeFunc(index, *item.ID(), *item.Value())
	}
}
