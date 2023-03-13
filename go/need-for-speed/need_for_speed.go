package speed

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	fmt.Printf("car is battery:%d batteryDrain: %d, speed:%d\n", car.battery, car.batteryDrain, car.speed)
	if car.battery-car.batteryDrain < 0 {
		fmt.Printf("not enough battery.\n")
		return car
	} else {
		var newBattery = car.battery - car.batteryDrain
		fmt.Printf("newBattery is: %d\n", newBattery)
		return Car{battery: newBattery, speed: car.speed, batteryDrain: car.batteryDrain,
			distance: car.distance + car.speed}
	}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var distance = track.distance
	var speed = car.speed
	var drain = car.batteryDrain
	var consume_rate = distance / speed
	var battconsumed = consume_rate * drain
	fmt.Printf("dist: %d, speed:%d, drain: %d, rate: %d, battConsume: %d\n",
		distance, speed, drain, consume_rate, battconsumed)
	if battconsumed <= car.battery {
		return true
	} else {
		return false
	}
}
