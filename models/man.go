package models

type Man struct {
	Age       int
	Height    float32
	Weight    float32
	Breathing bool
	Thinking  bool
	Eating    bool
	IsAlive   bool
}

func (man *Man) Breath()        { man.Breathing = true }
func (man *Man) Eat()           { man.Eating = true }
func (man *Man) Think()         { man.Thinking = true }
func (man *Man) Gender() string { return "Hombre" }
