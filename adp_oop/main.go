package main

import "fmt"

type FixedPriceJob struct {
	description string
	fixedprice  float64
}

type HourlyJob struct {
	description string
	hourlyRate  float64
	numberHours int
}

type JobInterface interface {
	Cost() float64
	GetDescription() string
}

func (job FixedPriceJob) Cost() float64 {
	return job.fixedprice
}

func (job FixedPriceJob) GetDescription() string {
	return job.description
}

func (job HourlyJob) Cost() float64 {
	return job.hourlyRate * float64(job.numberHours)
}

func (job HourlyJob) GetDescription() string {
	return job.description
}

func TotalJobCost(jobs []JobInterface) float64 {
	result := 0.0
	for _, job := range jobs {
		result += job.Cost()
	}
	return result
}

func main() {
	job1 := FixedPriceJob{"Stucco House", 34760.0}
	job2 := HourlyJob{"Landscaping", 40.0, 50}
	jobs := []JobInterface{job1, job2}
	totalCost := TotalJobCost(jobs)
	fmt.Printf("Total job cost: $ %0.2f", totalCost)
}
