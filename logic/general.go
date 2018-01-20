package logic

type cardHolder struct {
	cards []*card
}

func (t *cardHolder) add(c *card) {
	t.cards = append(t.cards, c)
}

func (t cardHolder) hasCard(c *card) bool {
	for i := range t.cards {
		if t.cards[i].suit == c.suit && t.cards[i].group == c.group {
			return true
		}
	}
	return false
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