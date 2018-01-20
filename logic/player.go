package logic

type hand struct {
	cardHolder
}

type Player struct {
	openHand   hand
	closedHand hand
	//yakus      []Yaku
	name       string
	id         int
}