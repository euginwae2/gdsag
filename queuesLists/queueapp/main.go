package main

import (
	"fmt"
	"math"
	"math/rand"

	"example.com/nodequeue"
)

const (
	arrivalRate =  0.25 //average customer arrival per minute
	lowerBoundServiceTime = 0.5 //arrivalRate
	upperBoundServicceTime = 2.0 //arrivalRate
	quitTime = 480 //minutes in an 8 hour day
)

func InterArrivalInterval(arrivalRate float64) float64 {
	// models a poisson process and returns
	// P(Wait time > t) = e^(-gamma * t)
	rn := rand.Float64()
	return -math.Log(1.0 -rn)
}

func ServiceTime() float64 {
	// Uniform distribution
	rn := rand.Float64()
	return lowerBoundServiceTime + (upperBoundServicceTime - lowerBoundServiceTime) * rn
}

type Customer struct {
	arrivalTime float64
	serviveDuration float64
}

// ADT for statistics
type Statistics struct {
	waitTimes []float64
	queueTime float64
	longestQueue int
	longestWaitTime float64
}

func (s *Statistics) AddWaitTime(wait float64) {
	s.waitTimes = append(s.waitTimes, wait)
	if wait > s.longestWaitTime {
		s.longestWaitTime = wait
	}
}

func (s *Statistics) AddQueueSizeTime(queueSize int, timeAtSize float64) {
	s.queueTime += float64(queueSize) * timeAtSize
}

func (s *Statistics) AddLength(length int) {
	if length > s.longestQueue {
		s.longestQueue = length
	}
}

var lastArrivalTime, departureTime, lastEventTime float64

func main() {
	lastEventTime = 0.0
	line := nodequeue.Queue[Customer]{}
	statistics := Statistics{}
	// Start simulation
	for {
		lastArrivalTime =  lastArrivalTime + InterArrivalInterval(arrivalRate)
		if lastArrivalTime > quitTime {
			break
		}
		if line.Size() == 0 {
			lastEventTime =  lastArrivalTime
			fmt.Printf("\nline no longer empty at time; %0.2f. line size is 1",lastEventTime)
			serviceTime := ServiceTime()
			customer := Customer{lastArrivalTime, serviceTime}
			line.Insert(customer)
			statistics.AddLength(line.Size())
			departureTime = lastArrivalTime + serviceTime
		} else {
			if lastArrivalTime < departureTime {
				// next event is an arrival
				customer := Customer{lastArrivalTime, ServiceTime()}
				statistics.AddQueueSizeTime(line.Size(), lastArrivalTime - lastEventTime)
				lastEventTime =  lastArrivalTime
				line.Insert(customer)
				fmt.Printf("\nArriival event at %0.2f - line size is: %d ", lastEventTime, line.Size())
				statistics.AddLength(line.Size())
			} else {
				// next event is a departure
				statistics.AddQueueSizeTime(line.Size(), departureTime - lastEventTime)
				departingCustomer := line.Remove()
				statistics.AddWaitTime(departureTime - departingCustomer.arrivalTime)
				lastEventTime =  departureTime
				fmt.Printf("\nDeparture event at %0.2f - line size is : %d", lastEventTime, line.Size())
				if line.Size() > 0 {
					departureTime =  lastEventTime + line.First().serviveDuration
				}
			}
		}
	}

	totalWaitTime := 0.0
	for i := 0; i < len(statistics.waitTimes); i++ {
		totalWaitTime += statistics.waitTimes[i]
	}

	averageWaitTime := totalWaitTime / float64(len(statistics.waitTimes))
	fmt.Printf("\naverage Time from Arrival to Departure : %0.2f minutes", averageWaitTime)
	fmt.Printf("\nAverage Size of waiting line : %0.2f", statistics.queueTime / lastEventTime)
	fmt.Printf("\nLo9ngest queue during the day: %d: ", statistics.longestQueue)
	fmt.Printf("\nLongest wait time during the day: %0.2f minutes", statistics.longestWaitTime)
}