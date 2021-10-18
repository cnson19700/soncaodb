package model

type Student struct {
	ID        int     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string  `gorm:"size:50;not null" json:"name"`
	MathSco   float64 `gorm:"math_sco" json:"math_sco"`
	PhysicSco float64 `gorm:"physic_sco" json:"physic_sco"`
	ChemiSco  float64 `gorm:"chemi_sco" json:"chemi_sco"`
}
