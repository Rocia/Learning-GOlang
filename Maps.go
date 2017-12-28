package main
import "fmt"
func main(){
	m := make(map[string]int)
	m["Key 1"], m["Key 2"] = 7, 12
	fmt.Println("Map (Dictionary): ",m)
	value1 , value2 := m["Key 1"], m["Key 2"]
	fmt.Println("values fetched by key:", value1, value2)
	fmt.Println("Length of Map:", len(m))
	delete (m, "Key 1")
	fmt.Println("map after deletion: ", m)
	_, prs := m["Key 1"]
	fmt.Println("Check if key 'Key1' is present:", prs)

	//single line declaration
	n := map[string]int {"spam":1, "eggs": 2}
	fmt.Println("Map:", n)

}

/*
Map (Dictionary):  map[Key 1:7 Key 2:12]
values fetched by key: 7 12
Length of Map: 2
map after deletion:  map[Key 2:12]
Check if key 'Key1' is present: false
Map: map[eggs:2 spam:1]
 */