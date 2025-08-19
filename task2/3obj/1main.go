package main

import (
	"fmt"
	"math"
)

func main() {

	// 创建矩形实例
	rect := Rectangle{Width: 5, Height: 3}
	// 创建圆形实例
	circle := Circle{Radius: 4}

	// 打印矩形信息
	fmt.Println("矩形:")
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())

	// 打印圆形信息
	fmt.Println("\n圆形:")
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Perimeter())

	// 演示多态性：使用Shape接口数组
	shapes := []Shape{rect, circle}
	fmt.Println("\n通过接口数组访问:")
	for i, shape := range shapes {
		fmt.Printf("形状 %d:\n", i+1)
		fmt.Printf("  面积: %.2f\n", shape.Area())
		fmt.Printf("  周长: %.2f\n", shape.Perimeter())
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius

}
