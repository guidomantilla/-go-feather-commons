package collections

import (
	feather_commons_constraints "github.com/guidomantilla/go-feather-commons/pkg/constraints"
)

type ComparableAllowedTypes interface {
	comparable
}

type IDAllowedTypes interface {
	feather_commons_constraints.UnsignedInteger | string
}

type ValueAllowedTypes interface {
	any
}

type Item[ID IDAllowedTypes, V ValueAllowedTypes] interface {
	ID() *ID
	Value() *V
}

type IterateeFunc[ID IDAllowedTypes, V ValueAllowedTypes] func(index int, id ID, item V)

type FilterFunc[ID IDAllowedTypes, V ValueAllowedTypes] func(index int, id ID, item V) bool

type CountFunc[ID IDAllowedTypes, V ValueAllowedTypes] func(index int, id ID, item V) bool

type Items[ID IDAllowedTypes, V ValueAllowedTypes] interface {
	IDs() []ID
	Values() []*V
	Contains(id ID) bool
	Find(id ID) (*V, error)
	Filter(filterFunc FilterFunc[ID, V]) Items[ID, V]
	Count(countFunc CountFunc[ID, V]) int
	Select(ids []ID) []*V
	ForEach(iteratee IterateeFunc[ID, V])
}

func NewItems[ID IDAllowedTypes, V ValueAllowedTypes](items ...Item[ID, V]) Items[ID, V] {
	return NewDefaultItems(items...)
}
