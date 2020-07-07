package main

import (
	"fmt"
	"math"
	"sync"
)

func getMinimum(sli []int) int {
	min := sli[0]
	minIndex := 0
	for i := 0; i < len(sli); i++ {

		if sli[i] < min {
			min = sli[i]
			minIndex = i
		}
	}
	return minIndex
}
func sortPartition(sli []int, wg *sync.WaitGroup) {
	for i := 0; i < len(sli)-1; i++ {
		min := getMinimum(sli[i:len(sli)])
		sli[i], sli[i+min] = sli[i+min], sli[i]
	}
	fmt.Println("Sorted Partition:", sli)
	wg.Done()
}
func merge(sli []int, wg *sync.WaitGroup) []int {
	n := len(sli)
	a := sli[0:r1]
	b := sli[r1:r2]
	c := sli[r2:r3]
	d := sli[r3:n]
	sortedSli := make([]int, 0, n)
	ai, bi, ci, di := 0, 0, 0, 0
	for ai+bi+ci+di < len(sli) {
		elem := make([]int, 0, 4)
		if ai < len(a) {
			elem = append(elem, a[ai])
		} else {
			elem = append(elem, 2147483647)
		}
		if bi < len(b) {
			elem = append(elem, b[bi])
		} else {
			elem = append(elem, 2147483647)
		}
		if ci < len(c) {
			elem = append(elem, c[ci])
		} else {
			elem = append(elem, 2147483647)
		}
		if di < len(d) {
			elem = append(elem, d[di])
		} else {
			elem = append(elem, 2147483647)
		}
		switch getMinimum(elem) {
		case 0:
			{
				sortedSli = append(sortedSli, a[ai])
				ai++
			}
		case 1:
			{
				sortedSli = append(sortedSli, b[bi])
				bi++
			}
		case 2:
			{
				sortedSli = append(sortedSli, c[ci])
				ci++
			}
		case 3:
			{
				sortedSli = append(sortedSli, d[di])
				di++
			}
		default:
			{
				break
			}
		}
	}
	return (sortedSli)
}

//partitions parameter
var r1, r2, r3 = 0, 0, 0

func main() {
	var n int
	fmt.Println("Enter total number you want to enter: ")
	fmt.Scan(&n)
	sli := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sli[i])
	}
	fmt.Println("Unsorted list: ", sli, "\n")
	//Partition parameters
	r1 = int(math.Round(float64(n)/4.0)) * 1
	r2 = int(math.Round(float64(n)/4.0)) * 2
	r3 = int(math.Round(float64(n)/4.0)) * 3
	if r3 > n {
		r3--
	}
	var wg sync.WaitGroup
	wg.Add(4)
	go sortPartition(sli[0:r1], &wg)
	go sortPartition(sli[r1:r2], &wg)
	go sortPartition(sli[r2:r3], &wg)
	go sortPartition(sli[r3:n], &wg)
	wg.Wait()
	SortedSlice := merge(sli, &wg)
	fmt.Println("\nSorted list:", SortedSlice)
}
