package main

/*
type person struct {
	First string                       //<-- we need to put the first letter of the property bigger
}                                      //   as it will be called by other classes also while marshling-unmarshling
                                       // if we do-not use first letter big then this property will not be processed
func main() {
	p1 := person{
		First: "sunny",
	}
	p2 := person{
		First: "Bunny",
	}

	xp := []person{p1, p2}
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bs, "++++++++++")
	fmt.Println(string(bs), "------------")

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

    http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person {
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
      var p2 person
	  err := json.NewDecoder(r.Body).Decode(&p2)
	  if err != nil {
		  log.Println("Decoded bad data ", err)
	  }

	  log.Println(p2)
}

*/
