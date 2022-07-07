package main

import "fmt"

//check if height is >0
func invalidHeight(height float64) string {
	return fmt.Sprintf("Height of a rectangle must be a number greater than zero, given height: %g", height)
}

//check if width is >0
func invalidWidth(width float64) string {
	return fmt.Sprintf("Width of a rectangle must be a number greater than zero, given width: %g", width) //I don't know why but it also catches if you give string
}
