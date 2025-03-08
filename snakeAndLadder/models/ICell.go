package models

type ICell interface {
	SetJump(j *Jump)
	GetJump() *Jump
}
