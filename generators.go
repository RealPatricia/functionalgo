package functionalgo

import (
	"fmt"
	"math"
)

// Generator

func Generator[T any, U func(in T) T, V func(in T) bool](i T, s U, q V) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		state := i
		for {
			if q(state) {
				break
			}

			out <- state
			state = s(state)
		}
	}()
	return out
}

// Quitters

func QuitIfLessThan[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing < val
	}
}

func QuitIfLessThanOrEqual[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing <= val
	}
}

func QuitIfGreaterThan[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing > val
	}
}

func QuitIfGreaterThanOrEqual[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing >= val
	}
}

func QuitIfEqual[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing == val
	}
}

func QuitIfNotEqual[T Ordered](val T) func(in T) bool {
	return func(thing T) bool {
		return thing < val
	}
}

// Sources

func Adder[T Number](in T) func(in T) T {
	return func(val T) T {
		return in + val
	}
}

func Subtractor[T Number](in T) func(in T) T {
	return func(val T) T {
		return in - val
	}
}

func Multiplier[T Number](in T) func(in T) T {
	return func(val T) T {
		return in * val
	}
}

func Divider[T Number](in T) func(in T) T {
	return func(val T) T {
		return in / val
	}
}

// Range Generator

func ValidateRange[T OrderedNumber](start, end, step T) bool {

	steppingDown := math.Signbit(float64(step))
	goingDown := start > end
	valid := goingDown == steppingDown
	return valid
}

func GenerateRange[T OrderedNumber](start, end, step T) (chan T, error) {
	out := make(chan T)
	valid := ValidateRange(start, end, step)
	goingDown := math.Signbit(float64(step))

	if !valid {
		close(out)
		return out, fmt.Errorf("Could not generate range [%v:%v] due to an invalid step size of [%v]", start, end, step)
	}

	source := func(in T) T {
		return in + step
	}

	var quitter func(in T) bool

	if goingDown {
		quitter = QuitIfLessThanOrEqual(end)
	} else {
		quitter = QuitIfGreaterThanOrEqual(end)
	}

	out = Generator(start, source, quitter)

	return out, nil
}
