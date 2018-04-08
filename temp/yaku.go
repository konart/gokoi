package temp

type Rule struct{}

//returns after the first successful match
func (r *Rule) toSatisfy(card *Card, criterias []map[string]string, passed *bool) {
	for _, criteria := range criterias {
		*passed = card.checkCriteria(criteria)
		if *passed {return}
	}
	return
}

//returns after the first successful match
func (r *Rule) toExclude(card *Card, criterias []map[string]string, passed *bool) {
	for _, criteria := range criterias {
		*passed = !card.checkCriteria(criteria)
		if !(*passed) {return}
	}
	return
}

var r = &Rule{}

type Yaku interface {
	Rules(c *Card)
}

type BasicYaku struct {
	cardsCollected int
	totalPoints    int
	cardsToFinish  int
	defPoints      int
	canOverDraft   bool
	finished       bool
}

type Sanko struct {
	BasicYaku
}

func (s *Sanko) Rules(c *Card) {
	passed := false
	criterias := []map[string]string{{"group": "bright"}}
	toExclude := []map[string]string{{"suit": "Willow", "group": "bright"}}
	r.toSatisfy(c, criterias, &passed)
	r.toExclude(c, toExclude, &passed)
	if passed {
		s.cardsCollected += 1
	}
	if s.cardsToFinish == s.cardsCollected {
		s.finished = true
		s.totalPoints = s.defPoints
	}
	if s.canOverDraft && s.cardsCollected > s.cardsToFinish {
		s.totalPoints += 1
	}
}

///same yakus

func setYakus() (Yakus []Yaku) {
	Yakus = []Yaku{
		&Sanko{BasicYaku{0, 0, 3, 6, false, false}},
	}
	return
}
