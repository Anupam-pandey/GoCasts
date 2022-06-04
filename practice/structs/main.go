package main

import ("fmt"
)



type contactinfo struct{
	email string
	zipcode int

}

type person struct{
	firstname string
	lastname  string
	contact  contactinfo
}

func main(){
	p := person{firstname:"anupam",lastname:"pandey"}
	fmt.Println(p.firstname)
	var pp person
	fmt.Printf("%+v",pp)
	pp.firstname  = "yo"
	pp.contact.email = "abc@sd"
	pp.contact.zipcode = 1103553
	fmt.Printf("%+v",pp)

	ppp := person{
		firstname:"abc",
		lastname:"def",
		contact: contactinfo{
			email: "ghi@jkl",
			zipcode: 224454,
		},
	}
	ppp.print()
	pointertoppp := &ppp
	pointertoppp.updatename("yo")
	ppp.print()
}

func (p person) print(){
	fmt.Printf("%+v",p)
}



func (pointertoperson *person) updatename(name string){
	(*pointertoperson).firstname = name
}