package oop

type IBuilder interface {
	setName(name string)
	setAge(age int)
	GetName() string
	GetAge() int
	ToString() string
}
