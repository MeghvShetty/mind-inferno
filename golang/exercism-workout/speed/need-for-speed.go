package speed

import "fmt"

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

// TODO: define the 'Drive' type struct

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	RemoteControlled := Car{
		speed:        speed,
		battery:      100,
		batteryDrain: batteryDrain,
	}
	return RemoteControlled
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	Track := Track{distance: distance}
	return Track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if NewBatteryStatus := car.battery - car.batteryDrain; NewBatteryStatus >= 0 {
		car.battery = NewBatteryStatus
		car.distance += car.speed
	}
	return car

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	BatteryLeft := (track.distance/car.speed)*car.batteryDrain <= car.battery
	return BatteryLeft
}

// Struct calls
func main() {
	fmt.Println(NewCar(5, 2))
	fmt.Println(NewTrack(800))
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(Drive(car))
}
