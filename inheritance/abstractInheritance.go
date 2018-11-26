package inheritance

type creature struct {}
type human struct {*creature}
type animal struct {*creature}

type Man struct {*human}
type Woman struct {*human}
type Horse struct {*animal}
type Dog struct {*animal}

type Creature interface {

}
type Human interface {
	Creature
}
type Animal interface {
	Creature
}