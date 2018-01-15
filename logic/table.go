package logic

type table struct {
	cards map[*card]bool
}

func (t *table) add(c *card) {
	t.cards[c] = true
}