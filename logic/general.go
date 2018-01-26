package logic

type cardHolder struct {
	cards map[*card]bool
}

func (t *cardHolder) add(c *card) {
	t.cards[c] = true
}

func (t cardHolder) hasCard(c *card) bool {
	return t.cards[c]
}

type cardHolderInterface interface {
	hasCard(*card) bool
}

func DealCards(d deck, t *table, p1, p2 *Player) {
	for i := 0; i < 8; i++ {
		p1.closedHand.add(d.PickCard())
		p2.closedHand.add(d.PickCard())
		t.add(d.PickCard())
	}
}

func CheckCards(c *card, ch cardHolderInterface) bool {
	if !ch.hasCard(c) {
		return false
	}
	return true
}