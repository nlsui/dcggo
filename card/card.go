package card

import basics "github.com/nlsui/dccgo/basics"

type Card struct{
	//ability
	speedmodifier basics.Speedmodifier
	//tag
	//text
}

type Entity struct{
	*Card
	//ability
	//damage
	//life
}

type Location struct{
	*Card
	//ability
	//life
}

type Action struct{
	*Card
	//effect
}