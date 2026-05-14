package main

import "fmt"

type cost struct{
	day int
	value float64
}

func getCostByDay(costs []cost) []float64{

	var costsByDay []float64
	

	for i:=0;i<len(costs);i++{
		cost := costs[i]
		for cost.day >= len(costsByDay){
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

func test(costs []cost){
	fmt.Println("creating daily buckets for %v costs ...\n", len(costs))

	costsByDay := getCostByDay(costs)
	fmt.Printf("costs by day: %v\n", costsByDay)
}


func main(){

	costs := []cost{
		{day: 0, value: 10.0},
		{day: 1, value: 20.0},
		{day: 0, value: 5.0},
		{day: 2, value: 15.0},
	}
	test(costs)
}