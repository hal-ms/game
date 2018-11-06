package repo

import (
	"github.com/hal-ms/game/model"
)

var State = stateRepo{state: model.State{false, false}}

type stateRepo struct {
	state model.State
}

func (s *stateRepo) Get() model.State {
	return s.state
}

func (s *stateRepo) Set(state model.State) {
	s.state = state
}

func (s *stateRepo) IsWearing(state bool) {
	s.state.IsWearing = state
}
func (s *stateRepo) IsStandby(state bool) {
	s.state.IsStandby = state
}
