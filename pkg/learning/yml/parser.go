package yml

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

// TODO check yaml.v3 too and see how it is different.
// TODO use decoder for fine grained parsing.

type BType struct {
	RenamedC int `yaml: "c"`
	D        []int
}

type T struct {
	A string
	B BType
}

func ParseYaml() error {
	data := `
a: 'name with space'
b:
  c: 10
  d: [2,5]
`
	t := T{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Printf("error in unmarshal: %v", err)
		return err
	}

	fmt.Printf("unmarsaled struct is %+v", t)

	t1 := &T{
		A: "some value with space ",
		B: BType {
			RenamedC: 123,
			D:        []int{4, 6},
		},
	}

	d1, err := yaml.Marshal(t1)
	if err != nil {
		log.Printf("error in marshal: %v", err)
		return err
	}

	log.Printf("marshalled yaml: \n%s\n", string(d1))


	unst := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &unst)
	if err != nil {
		log.Printf("error in unmarshal with unstructured: %v", err)
		return err
	}

	log.Printf("unmarshal unstructued: \n%+v", unst)

	return nil
}
