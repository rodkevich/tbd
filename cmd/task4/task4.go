package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
)

func main() {
	arr1 := [6]int{17, 18, 5, 4, 6, 1}
	err := someFunction(arr1)
	if err != nil {
		return
	}

	arr2 := [1]int{400}
	err = someFunction(arr2)
	if err != nil {
		return
	}

	arr3 := [12]int{17, 1, 0, -10, 18, -5, 17, 12, 0, -10, 0, -5}
	err = someFunction(arr3)
	if err != nil {
		return
	}

	slice := []int{0, 001, 00001, 0000000001}
	err = someFunction(slice)
	if err != nil {
		log.Fatal(err) // NOTE: should fail
	}
}

func someFunction(i interface{}) error {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Array {
		arrLength := val.Len()
		arr := make([]int64, arrLength)
		for i := range arr {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				arr[i] = e.Int()
			default:
				return errors.New("not implemented")
			}
		}
		printSlice(arr)
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
		printSlice(arr)
		return nil
	}
	return errors.New("invalid argument type: %v")
}

func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
