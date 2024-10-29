package geometry

type Segment struct {
	A, B Point
}

func (s Segment) Length() float64 {
	return s.A.Distance(s.B)
}

func (s Segment) Eq(t Segment) bool {
	return s.A.Eq(t.A) && s.B.Eq(t.B) || s.A.Eq(t.B) && s.B.Eq(t.A)
}

func (s Segment) Contains(p Point) bool {
	return float64Eq(
		s.A.Distance(p)+s.B.Distance(p),
		s.Length(),
	)
}
