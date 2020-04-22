package main
import "fmt"
func map1() {
dict := map[int]string{1:"Rose"}
dict[2]="Lotus"
fmt.Println(dict)
dict[1]="Sunflower"
fmt.Println(dict)
delete(dict,1)
fmt.Println(dict)
}
func main() {
	map1()
}