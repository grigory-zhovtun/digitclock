package main

var (
	Zero 	= [...]string{"███", 
						  "█ █", 
						  "█ █", 
						  "█ █", 
						  "███"}

	One 	= [...]string{"██ ", 
	                      " █ ", 
	                      " █ ", 
	                      " █ ", 
	                      "███"}

	Two 	= [...]string{"███", 
	                      "  █", 
	                      "███", 
	                      "█  ", 
	                      "███"}

	Three 	= [...]string{
						  "███", 
						  "  █", 
					      "███", 
						  "  █", 
						  "███"}

	Four 	= [...]string{"█ █", 
	                      "█ █", 
	                      "███", 
	                      "  █", 
	                      "  █"}

	Five 	= [...]string{"███", 
	                      "█  ", 
	                      "███", 
	                      "  █", 
	                      "███"}

	Six 	= [...]string{"███", 
	                      "█  ", 
	                      "███", 
	                      "█ █", 
	                      "███"}

	Seven 	= [...]string{"███", 
	                      "  █", 
	                      "  █", 
	                      "  █", 
	                      "  █"}

	Eight 	= [...]string{"███", 
	                      "█ █", 
	                      "███", 
	                      "█ █", 
	                      "███"}

	Nine 	= [...]string{"███", 
	                      "█ █", 
	                      "███", 
	                      "  █", 
	                      "███"}

	Colon	= [...]string{"   ", " ▒ ", "   ", " ▒ ", "   "}
	Empty 	= [...]string{"   ", "   ", "   ", "   ", "   "}
)

var Digits 	= [...][5]string{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine}

var Display = [8][5]string{Zero, Zero, Colon, Zero, Zero, Colon, Zero, Zero}