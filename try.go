package main

import "fmt"


func main() {
	var masaStudi int
	var publikasi bool
	var cumlaude bool
	var sgMemuaskan bool
	var memuaskan bool

	fmt.Scanln(&masaStudi, &publikasi)
	if( masaStudi <= 8 and publikasi == true ) then 
		cumlaude = true
		sgMemuaskan = false 
		memuaskan = false
	else if(masaStudi <= 14 and publikasi == false ) then 
		cumlaude = false
		sgMemuaskan = true
		memuaskan = false
	else then 
		cumlaude = false
		sgMemuaskan = false 
		memuaskan = true 
	fmt.Println("Cumlaude : ",cumlaude)
	fmt.Println("sgMemuaskan : ",sgMemuaskan)
	fmt.Println("memuaskan : ",memuaskan)

}
