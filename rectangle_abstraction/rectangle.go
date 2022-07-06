package main

import (
	"errors"
)

// var rectangleSlice []rectangle

type rectangle struct {
	height float64
	width  float64
}

func newRectangleWithArgumentsReturns(height float64, width float64) (*rectangle, error) {
	rekt := rectangle{}
	if width <= 0 {
		err := errors.New(WidthNotGreaterThanZero(width))
		return nil, err
	}
	if height <= 0 {
		err := errors.New(HeightNotGreaterThanZero(height))
		return nil, err
	}
	rekt.width = width
	rekt.height = height
	return &rekt, nil
}

func (r rectangle) calculateArea() float64 {
	return r.width * r.height
}
func (r rectangle) calculateCircumference() float64 {
	return 2 * (r.width + r.height)
}

//Tried some other stuff here

// func newRectangleReturns() rectangle {
// 	height, width := 0., 0.
// 	fmt.Println("Please enter height of the rectangle: ")
// 	fmt.Scanln(&height)
// 	fmt.Println("Please enter height of the rectangle: ")
// 	fmt.Scanln(&width)
// 	rekt := rectangle{width: width, height: height}
// 	fmt.Printf("Created new rectangle! \n Height: %f \n width: %f", height, width)
// 	return rekt
// }

// func createNewRectangleNoReturn(height float64, width float64){
// 	rekt := rectangle{}
// 	rekt.width = width
// 	rekt.height = height
// 	rectangleSlice = append(rectangleSlice, rekt)
// }
