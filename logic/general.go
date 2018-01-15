package logic


func DealCards(d deck, t *table, p1, p2 *Player)  {
	for i := 0; i < 8; i++ {
		p1.closedHand = append(p1.closedHand, d.PickCard())
		p2.closedHand = append(p2.closedHand, d.PickCard())
		t.add(d.PickCard())
	}
}