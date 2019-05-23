package entities

type Money struct {
	Value int64
}

func MoneyEquals(m1, m2 *Money) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	return m1.Value == m2.Value
}
