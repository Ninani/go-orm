package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name" gorm:"text;not null;default:null`
	Quizzes []Quiz `gorm:"foreignKey:UserRefer"`
}

type Quiz struct {
	gorm.Model
	UserRefer uint
	Facts     []Fact `gorm:"many2many:quiz_facts"`
	Result    int
}

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}
