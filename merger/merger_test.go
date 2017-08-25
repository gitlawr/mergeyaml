package merger

import (
	"fmt"
	"testing"
)

const str1 = `---
foo-map:
  something: 
    thing: one
    thing: two
  newstruct:
  - foo-bar
  - 1234
  bar: "bar"`

const str2 = `---
foo-map:
  anotherthing: 
  - thing: three
  bar: newbar`

func TestMerge(t *testing.T) {
	res, err := MergeYaml([]byte(str1), []byte(str2))
	if err != nil {
		fmt.Printf("got error merging yaml: %v", err)
	}
	fmt.Printf("got merged yaml:\n%v\n", string(res))
}
