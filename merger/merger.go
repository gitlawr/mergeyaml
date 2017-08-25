package merger

import (
	// Base packages.
	"fmt"
	"reflect"

	// Third party packages.
	"gopkg.in/yaml.v2"
)

//MergeYaml merges yml1 into yml2 and return merged result
func MergeYaml(yml1 []byte, yml2 []byte) ([]byte, error) {
	var obj1 interface{}
	var obj2 interface{}
	if err := yaml.Unmarshal(yml1, &obj1); err != nil {
		fmt.Printf("get error:%v", err)
		return nil, err
	}
	if err := yaml.Unmarshal(yml2, &obj2); err != nil {
		fmt.Printf("get error:%v", err)
		return nil, err
	}
	res := MergeMap(obj1.(map[interface{}]interface{}), obj2.(map[interface{}]interface{}))
	fmt.Printf("get map:%v", res)
	out, err := yaml.Marshal(res)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//MergeMap maps from first to second
func MergeMap(first map[interface{}]interface{}, second map[interface{}]interface{}) map[interface{}]interface{} {
	if first == nil {
		return second
	}
	if second == nil {
		second = make(map[interface{}]interface{})
	}

	for k, v := range first {
		fmt.Printf("first type:%v", reflect.TypeOf(first[k]))
		fmt.Printf("second type:%v", reflect.TypeOf(second[k]))
		if reflect.TypeOf(second[k]) != reflect.TypeOf(first[k]) {
			second[k] = v
		} else if reflect.TypeOf(first[k]) == reflect.TypeOf(map[interface{}]interface{}{}) {
			//merge maps
			second[k] = MergeMap(first[k].(map[interface{}]interface{}), second[k].(map[interface{}]interface{}))
		} else if reflect.TypeOf(first[k]) == reflect.TypeOf([]interface{}{}) {
			//for other types,replace it with value in first map.
			second[k] = v
		}
	}

	return second
}
