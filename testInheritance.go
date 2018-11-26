package main

import inh "github.com/nlsui/dcggo/inheritance"
import prop "github.com/DecentralCardGame/cardobject"
import "fmt"
import "reflect"
import "encoding/json"

func main() {
	man := inh.Man{}
	woman := inh.Woman{}
	horse := inh.Horse{}
	dog := inh.Dog{}

	fmt.Println(reflect.ValueOf(man))
	fmt.Println(reflect.ValueOf(woman))
	fmt.Println(reflect.ValueOf(horse))
	fmt.Println(reflect.ValueOf(dog))

	damage := prop.NewAttack(5)
	text := prop.NewText("text")
	handsize := prop.NewHandSize(2)

	fmt.Println(damage.ExtractIntProp())
	fmt.Println(text.ExtractStringProp())

	fmt.Println(damage.IsCardProperty())

	fmt.Println(damage.IsProperty())

	cond, err := prop.NewIntCondition(damage, prop.GREATER, 7)

	fmt.Println(cond.GetComparator())
	fmt.Println(cond.ExtractIntCond())

	cond2, err := prop.NewIntCondition(handsize, prop.GREATER, 7)
	fmt.Println(cond2.GetComparator())

	if (err != nil) {
		fmt.Println(err.Error())
	}

    bytes, err := json.Marshal(text)
    if err != nil {
        fmt.Println("Can't serislize", text)
    }
 
    fmt.Printf("%v => %v, '%v'\n", text, bytes, string(bytes))
}