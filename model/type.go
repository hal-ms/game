package model

type State struct {
	IsStandby bool
	IsWearing bool
	IsHit     bool
}

type Hit struct {
	Point int
}

type Job struct {
	Jobs []string
	Job  string
}
