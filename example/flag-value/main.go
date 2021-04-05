package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type idsFlag []string
type Person struct {
	name string
	born time.Time
}

func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

func (p Person) String() string {
	return fmt.Sprintf("name: %s born on: %s", p.name, p.born.String())
}

func (p *Person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var ids idsFlag
	var p Person

	flag.Var(&ids, "id", "The id will be added to the list.")
	flag.Var(&p, "name", "The name of a person.")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
	allIds := strings.Split(ids[0], ",")
	fmt.Println(allIds[len(allIds)-1])

}
