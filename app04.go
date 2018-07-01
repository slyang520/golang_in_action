package main

import "fmt"

// map
//  var map1 map[keytype]valuetype

// var map1 = make(map[keytype]valuetype)
// map1 := make(map[keytype]valuetype)

func main() {

	var mapLit map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}

	var mapAssigned map[string]int
	mapAssigned = mapLit
	mapAssigned["two"] = 3

	//  map1 := make(map[keytype]valuetype)  <=> mapCreated := map[string]float32{}
	mapCreated := make(map[string]float32)
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	// 删除KEY
	// delete(mapLit, "one")
	test_v, ok := mapLit["one"] // 如果key1存在则ok == true，否则ok为false
	if ok{
		fmt.Println(test_v);
	}else{
		fmt.Println("not present");
	}

	for key, value := range mapLit {
		fmt.Printf("Map key = %s  value = %d \n", key,value)
	}



}
