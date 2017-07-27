package main

import "fmt"

type Vec2 struct {
	X, Y int32
}

func (self Vec2) Equal(other Vec2) bool {

	return other.X == self.X && other.Y == self.Y

}

func (self Vec2) Add(other Vec2) Vec2 {

	return Vec2{
		self.X + other.X,
		self.Y + other.Y,
	}
}

func (self *Vec2) SetX(v int32) {

	self.X = v
}

func main() {

	v1 := Vec2{1, 1}

	if v1.Add(Vec2{2, 3}).Equal(Vec2{3, 4}) {
		fmt.Println("vec{1, 1} + vec{2, 3} == vec{3, 4}")
	}

	v1.SetX(2)

	fmt.Println("vec{1, 1}  set vec.x = 1,  vec==", v1)

}
