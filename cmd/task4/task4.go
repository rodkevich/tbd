package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
)

var (
	maxLength = math.Pow(10, 4)
	maxValue  = math.Pow(10, 5) + 1
)

var usageError = errors.New("usage error")

func main() {
	// Пример 1: замена большим элементом справа от текущего,
	// а самый последний в массиве на -1
	arr1 := [6]int{17, 18, 5, 4, 6, 1}
	_, err := someFunction(arr1)
	if err != nil {
		log.Fatal(err)
	}

	// Пример 2: There are no elements to the right of index 0.
	arr2 := [1]int{400}
	_, err = someFunction(arr2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 &lt;= arr.length &lt;= 10в4
	var maxLengthFailure = [10001]int{}
	_, err = someFunction(maxLengthFailure)
	if err != nil {
		log.Fatal(err) // NOTE: should fail
	}

	// 1 &lt;= arr[i] &lt;= 10в5
	maxValueFailure := 100001
	arr3 := [12]int{17, 1, 0, -10, 18, -5, 100000, maxValueFailure, 0, -10, 0, -5}
	_, err = someFunction(arr3)
	if err != nil {
		log.Fatal(err) // NOTE: should fail
	}

	wrongTypeFailure := []int{0, 001, 00001, 0000000001}
	_, err = someFunction(wrongTypeFailure)
	if err != nil {
		log.Fatal(err) // NOTE: should fail
	}
}

func someFunction(arr interface{}) (interface{}, error) {
	val := reflect.ValueOf(arr)

	if val.Kind() == reflect.Array && val.Len() <= int(maxLength) {
		arrLength := val.Len()
		entities := make([]int64, arrLength, arrLength)
		printSlice(entities)

		for i := range entities {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				if e.Int() >= int64(int(maxValue)) {
					return nil, usageError
				}
				entities[i] = e.Int()
			default:
				return nil, errors.New("not implemented")
			}
		}

		printSlice(entities)
		for i := 0; i < arrLength; i++ {
			tempSlice := entities[i:]
			var greatest int64 = -math.MaxInt64
			for _, value := range tempSlice[1:] {
				if value > greatest {
					greatest = value
					entities[i] = greatest
				}
			}
		}
		entities[arrLength-1] = -1
		printSlice(entities)

		// "После чего, верните массив arr."
		// NOTE: я не знаю как мне мутировать и возвращать arrays разной длины
		// на данный момент нет ...
		// поэтому пока возвращаю что умею. :)
		// for ind, el := range entities {
		// 	val[ind] = el
		// }
		// (type 'Value' does not support indexing)

		return entities, nil
	}
	return nil, usageError
}

func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
