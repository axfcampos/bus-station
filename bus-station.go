package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main(){

	var max int = 0
	var solution []int = make([]int, 0)  

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	input, _ := reader.ReadString('\n') 
	input = strings.TrimRight(input, "\r\n")

	array := strings.Split(input, " ")

	cum_sum := make([]int, len(array))

	for i, element := range array {
		ele, _ := strconv.Atoi(element)

		if ele > max {
			max = ele
		}

		if i == 0 {
			cum_sum[i] = ele
		} else {
			cum_sum[i] = cum_sum[i-1] + ele
		}
		//fmt.Printf("%d - ", cum_sum[i])
	}

	total_people := cum_sum[len(cum_sum)-1]
	//fmt.Println()

	for i, ele := range cum_sum {
		
		if ele < max { continue }
		if ele > (total_people / 2) { continue }
		if total_people % ele != 0 { continue }	

		n_mults := total_people / ele
		c := 0
		for _, elem := range cum_sum[i:] {
			if elem % ele == 0 { c++ }
		}

		if c == n_mults {
			solution = append(solution, ele)
		}
	}

	solution = append(solution, total_people)

	//Print solution
	for i := range solution {
		fmt.Printf("%d ", solution[i])
	}
}
