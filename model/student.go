package model

type Student struct {
	ID        int64   `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Name      string  `gorm:"column:name;size:50" json:"name"`
	MathSco   float64 `gorm:"column:math_sco" json:"math_sco"`
	PhysicSco float64 `gorm:"column:physic_sco" json:"physic_sco"`
	ChemiSco  float64 `gorm:"column:chemi_sco" json:"chemi_sco"`
}
