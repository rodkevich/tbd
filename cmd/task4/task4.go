package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func main() {
	arr1 := [6]int{17, 18, 5, 4, 6, 1}
	arr2 := [6]int{17, 18, 0, -10, 0, -5}
	arr3 := [1]int{3222333}
	err := someFunction(arr1)
	if err != nil {
		return
	}
	err = someFunction(arr2)
	if err != nil {
		return
	}
	err = someFunction(arr3)
	if err != nil {
		return
	}
}

func someFunction(i interface{}) error {
	val := reflect.ValueOf(i)
	fmt.Println("Got new: ", val.Kind())
	if val.Kind() == reflect.Array {
		arrLength := val.Len()
		arr := make([]int64, arrLength)
		for i := range arr {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				arr[i] = e.Int()
			default:
				return errors.New("invalid argument type: %v")
			}
		}
		printSlice(arr[:])
		for i := 0; i < arrLength; i++ {
			tempSlice := arr[i:]
			var greatest int64 = -math.MaxInt64
			for _, value := range tempSlice[1:] {
				if value > greatest {
					greatest = value
					arr[i] = greatest
				}
			}
		}
		arr[arrLength-1] = -1
		printSlice(arr[:])
	}
	return nil
}

func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
