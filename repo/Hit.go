package repo

import "github.com/hal-ms/game/model"

var Hit = hitRepo{}

// TODO 必要であれば排他処理を入れる
type hitRepo struct {
	hit model.Hit
}

func (h *hitRepo) Get() model.Hit {
	return h.hit
}

func (h *hitRepo) Reset() {
	h.hit.Point = 0
}

func (h *hitRepo) Add(p int) int {
	h.hit.Point += p
	return h.hit.Point
}
