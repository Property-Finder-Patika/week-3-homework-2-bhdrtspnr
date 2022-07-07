package main

import "fmt"

func main() {

	menu()

}

func menu() {

	var height, width float64 //define height and width
	fmt.Println("Please enter the height of the rectangle: ")
	fmt.Scanln(&height) //get height

	fmt.Println("Please enter the width of the rectangle: ")
	fmt.Scanln(&width) //get width

	newRectangle, err := newRectangleWithArgumentsReturns(height, width) //try to create a new rectangle if given arguments are not valid throw error
	if err != nil {                                                      //if error is not nil, it means there's an error
		fmt.Println("Failed to create a new rectangle")
		fmt.Println(err)
		return
	}

	//print rectangle's height, width, area and circumference
	fmt.Println("Successfully created  a new rectangle")
	fmt.Printf("Rectangle's height is %.2f \n", newRectangle.height)
	fmt.Printf("Rectangle's width is %.2f \n", newRectangle.width)

	fmt.Printf("Area of the rectangle is: %.2f.", newRectangle.calculateArea())
	fmt.Printf("Circumference of the rectangle is %.2f", newRectangle.calculateCircumference())

}
