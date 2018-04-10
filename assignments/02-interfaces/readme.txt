Write a program that creates two custom struct types called 'triangle' and 'square'.

The 'square' type should be a struct with a field called 'sideLength' of type float64.

The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64.

Both types should have function called 'getArea' that returns the calculated area of the square or triangle:
Area of a triangle = 0.5 * base * height.
Area of a square = sideLength * sideLength.

Add a 'shape' interface that defines a function called 'printArea'.
This function should calculate the area of the given shape and print it out to the terminal.
Design the interface so that the 'printArea' function can be called with either a triangle or a square.


Questions for this Assignment

1. Did you run into any error messages as you worked on the program?  If so, write the error message out here and a short explanation on what you think it might mean. 

Initially I declared the printArea() function like this:
    func (s shape) printArea() {...}
with the intention of calling it through a square and triangle instance like this:
    s.printArea()
    t.printArea()
But the linter (or whatever is checking the code in VSCode) gave me this error:
    "invalid receiver type shape (shape is an interface type)"
I then realized that the correct way to declare printArea() is like this:
    func printArea(s shape) {...}
And call it with a square and triangle instance like this:
    printArea(s)
    printArea(t)
Still I wonder if there is a way to declare the printArea() function so I can call it the way I intially intended, i.e.:
    s.printArea()
    t.printArea()
