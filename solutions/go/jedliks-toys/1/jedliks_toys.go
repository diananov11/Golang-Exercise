package jedlik
import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
    if car.battery < car.batteryDrain {
        return
    }
    car.distance += car.speed
    car.battery -= car.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %v meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %v%%",c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
    BatteryUse := float64(trackDistance)/float64(c.speed) * float64(c.batteryDrain)
    return float64(c.battery) >= BatteryUse
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
