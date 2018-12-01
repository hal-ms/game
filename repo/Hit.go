package repo

var Hit = hitRepo{0}

// TODO 必要であれば排他処理を入れる
type hitRepo struct {
	p int
}

func init() {
	go func() {

	}()
}
func (h *hitRepo) Get() int {
	return h.p
}

func (h *hitRepo) Reset() {
	h.p = 0
}

func (h *hitRepo) Add(p int) int {
	if p > 130 {
		return h.Get()
	}
	h.p += p
	return h.Get()
}
