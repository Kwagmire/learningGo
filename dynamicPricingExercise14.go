package main

import (
	"fmt"
)

type rideRequest struct {
	baseFare          float64
	distanceFare      float64
	trafficMultiplier float64
	demandMultiplier  float64
	totalFare         float64
}

func calculateFare(distance int, trafficLevel string, demandLevel string) *rideRequest {
	p := &rideRequest{}

	p.baseFare = 2.5

	rpd := 1.0
	p.distanceFare = float64(distance) * rpd

	switch trafficLevel {
	case "high":
		p.trafficMultiplier = 1.5
	default:
		p.trafficMultiplier = 1.0
	}

	switch demandLevel {
	case "peak":
		p.demandMultiplier = 1.8
	default:
		p.demandMultiplier = 1.0
	}

	p.totalFare = p.baseFare + (p.distanceFare * p.trafficMultiplier * p.demandMultiplier)

	return p
}

func main() {
	p := calculateFare(15, "normal", "peak")
	fmt.Printf("A ride which racks up $%v due to distance, "+
		"with a traffic multiplier of %v and a demand multiplier "+
		"of %v, will cost $%v in total.\n", p.distanceFare, p.trafficMultiplier, p.demandMultiplier, p.totalFare)
}
