package op

import (
	"github.com/fatih/color"
)

type state int

const (
	stateComplete state = iota
	stateIncomplete
	stateError
	stateUnknown
)

// State represents possible target or source state
type State struct {
	value state
	label string
}

func (st *State) String() string {
	if st.IsComplete() {
		return color.New(color.FgGreen).Sprint(st.label)
	}
	if st.IsError() {
		return color.New(color.FgRed).Sprintf("err: %s", st.label)
	}
	return st.label
}

// IsComplete checks whether state is `stateComplete`
func (st *State) IsComplete() bool {
	return st.value == stateComplete
}

// IsIncomplete checks whether state is `stateIncomplete`
func (st *State) IsIncomplete() bool {
	return st.value == stateIncomplete
}

// IsError checks whether state is `stateError`
func (st *State) IsError() bool {
	return st.value == stateError
}

// IsUnknown checks whether state is `stateUnknown`
func (st *State) IsUnknown() bool {
	return st.value == stateUnknown
}

// NewCompleteState returns State of `stateComplete`
func NewCompleteState(l string) *State {
	return &State{value: stateComplete, label: l}
}

// NewIncompleteState returns State of `stateIncomplete`
func NewIncompleteState(l string) *State {
	return &State{value: stateIncomplete, label: l}
}

// NewErrorState returns State of `stateError`
func NewErrorState(l string) *State {
	return &State{value: stateError, label: l}
}

// NewUnknownState returns State of `stateUnknown`
func NewUnknownState(l string) *State {
	return &State{value: stateUnknown, label: l}
}
