package main

import "fmt"

func main() {
	//notice that you canot use var and shorthand also the shorthand meaning := is function scoped you cannot use := with constants
	speakers, notes := "evanso", "tell not"

	//var barack go will need you to reduce the chances of errrors or make your code more robust by actually explicity declaring a type ...
	var conferenceNmae = "GO Conference"
	var remainingtickets int = 20
	//declarations with multiple initializations
	var events string
	//notice this happpens that a const can only be initialized by a literal or another const
	fmt.Println("welcome to", conferenceNmae, "glad yor are here with us todays speakers include:", speakers, "who will actually talk about", notes)
	fmt.Printf("here are the remaining %T tickets event : %v", remainingtickets, remainingtickets)
	fmt.Scan(&events)
}
