package speed

import "math"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}

	return car
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}

	return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if (car.distance > car.speed) || (car.battery < car.batteryDrain) {
		return car
	}
	car = Car{
		battery:      car.battery - car.batteryDrain,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     car.distance + car.speed,
	}

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	totalDistance := float64(track.distance)
	carTimes := int(math.Ceil(totalDistance / float64(car.speed)))
	carDrain := car.batteryDrain * carTimes
	return carDrain <= car.battery
}
