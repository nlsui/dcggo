package main

import card "github.com/DecentralCardGame/cardobject"
import "fmt"
import "reflect"

func main() {
	//z card.Zone
	rs := []card.Ressource{}
	rs = append(rs, card.FOOD)

	//zs := []card.Zone{card.DECK, card.DUSTPILE, card.FIELD, card.HAND, card.EXILE, card.STACK}

	cost := card.Cost{}
	costVars := reflect.ValueOf(&cost).Elem()
	//s := costVars.Type()

	fmt.Println(costVars)

	for i := 0; i < costVars.NumField(); i++ {

		fmt.Println(reflect.ValueOf(costVars.Field(i).Type()).Type())
	}
	d := card.Damage{}

	cci := card.CardConditionInt{d, card.EQUAL}


	//plc := card.PlayerLifeCondition{}
	//ev := card.ZoneChangeEvent{}

	fmt.Println(cci.GetCardProperty())
	//fmt.Println(plc.Comparator())
	//fmt.Println(*ev.Card)

	card.EXILE.ProtectedZoneID()
	//z := card.Zone{}
	//z := card.NewDustpileZone()
}