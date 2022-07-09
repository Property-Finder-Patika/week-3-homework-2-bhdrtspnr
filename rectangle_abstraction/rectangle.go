package main

import (
	"errors"
)

// var rectangleSlice []rectangle

type rectangle struct { //define rectangle struct
	height float64
	width  float64
}

func newRectangleWithArgumentsReturns(height float64, width float64) (*rectangle, error) { //main new rectangle function, takes height and width, returns rectangle and an error
	rekt := rectangle{} //init empty rectangle
	if width <= 0 {     //if width is less than zero throw error and return nil, error I don't know why but it also catches if you give string
		err := errors.New(invalidWidth(width))
		return nil, err
	}
	if height <= 0 { //if height is less than zero throw error and return nil, error
		err := errors.New(invalidHeight(height))
		return nil, err
	}
	//if there are no errors create a struct with given parameters and return it.
	rekt.width = width
	rekt.height = height
	return &rekt, nil
}

func (r rectangle) calculateArea() float64 { //calculate area of a rectangle
	return r.width * r.height
}
func (r rectangle) calculateCircumference() float64 { //calculate circumference of a rectangle
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
