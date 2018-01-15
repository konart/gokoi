package logic


//var (
//	newline = []byte{'\n'}
//	space   = []byte{' '}
//)

type Hand []*card

type Player struct {
	openHand   Hand
	closedHand Hand
	//yakus      []Yaku
	name       string
	id         int
}





