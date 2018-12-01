package repo

import "fmt"

var Hit = hitRepo{}

// TODO 必要であれば排他処理を入れる
type hitRepo struct {
	p int
}

func (h *hitRepo) Get() int {
	return h.p
}

func (h *hitRepo) Reset() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!RESET!!!!!!!!!!!!!!!!!!!!!")
	h.p = 0
}

func (h *hitRepo) Add(p int) int {
	h.p += p
	fmt.Println(h.Get())
	return h.Get()
}
