package wait

type Boat struct {
	Starting uint
	Increase uint
}

func (b *Boat) GetDistance(delay uint, remaining uint) uint {
	return b.Starting + b.Increase*delay*remaining
}

type Race struct {
	Boat
	Time     uint
	Distance uint
}

func (r *Race) GetWinningTimes() (counter uint) {
	for i := uint(0); i < r.Time; i++ {
		if r.GetDistance(i, r.Time-i) > r.Distance {
			counter++
		}
	}
	return
}
