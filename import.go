package testimport

type Type1 struct{
	ID int
	Name String
	text String
}

func (*t Type1) ChangeID(int i){
	t.ID := i
}