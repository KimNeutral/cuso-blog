package model

// Primary : primary type for basic gorm model
type Primary uint

func NewPrimary(i uint) Primary {
	var t = Primary(i)

	return t
}
