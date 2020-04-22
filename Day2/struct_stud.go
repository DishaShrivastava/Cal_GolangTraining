package main
 
import "fmt"

type contactinfo struct{
    phone_no string
    email string
}
type Student struct { 
    Name string
    contact contactinfo
}
 
func main( )  {
	stud := Student{Name: "st. Joshep",contact: contactinfo{phone_no:"70896587", email:"joshep.20@gmail.com"}}
	//fmt.Printf("%+v", stud)
	stud.Name="st.Joshep"
	stud.contact = contactinfo{
		phone_no:"70896587",
		email:"joshep@gmail.com",
}
    fmt.Printf("%+v", stud)
}