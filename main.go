package functionalgo

import (
	"golang.org/x/exp/constraints"
)

type OrderedNumber interface {
	constraints.Float | constraints.Integer
}

type Number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

type Ordered = constraints.Ordered
