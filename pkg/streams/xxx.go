package streams

import feather_commons_constraints "github.com/guidomantilla/go-feather-commons/pkg/constraints"

type IDAllowedTypes interface {
	feather_commons_constraints.UnsignedInteger | string
}

type ValueAllowedTypes interface {
	any
}

type ValueObject[ID IDAllowedTypes, V ValueAllowedTypes] interface {
	ID() *ID
	Value() *V
}
