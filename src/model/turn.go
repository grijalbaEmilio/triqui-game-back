package model

type Turn uint

const (
	X Turn = iota
	O
)

func (t Turn) String() string {
	return [...]string{"X", "O"}[t]
}
