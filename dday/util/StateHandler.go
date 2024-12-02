package util

type StateHandler[T comparable] struct {
	states  []T
	current int
}

// NewStateHandler initializes the state handler with a slice of states.
func NewStateHandler[T comparable](states []T, cur int) *StateHandler[T] {
	if len(states) == 0 {
		panic("states cannot be empty")
	}
	return &StateHandler[T]{states: states, current: cur}
}

// CurrentState returns the current state.
func (sh *StateHandler[T]) CurrentState() T {
	return sh.states[sh.current]
}

// Returns current state's index
func (sh *StateHandler[T]) Index() int {
	return sh.current
}

// NextState transitions to the next state, wrapping around if necessary.
func (sh *StateHandler[T]) NextState() T {
	sh.current = (sh.current + 1) % len(sh.states)
	return sh.CurrentState()
}

// PrevState transitions to the previous state, wrapping around if necessary.
func (sh *StateHandler[T]) PrevState() T {
	if sh.current == 0 {
		sh.current = len(sh.states) - 1
	} else {
		sh.current--
	}
	return sh.CurrentState()
}
