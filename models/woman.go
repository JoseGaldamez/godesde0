package models

type Woman struct {
	Man
}

func (woman *Woman) Breath()        { woman.Breathing = true }
func (woman *Woman) Eat()           { woman.Eating = true }
func (woman *Woman) Think()         { woman.Thinking = true }
func (woman *Woman) Gender() string { return "Mujer" }
