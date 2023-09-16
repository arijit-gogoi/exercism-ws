package elon

import "fmt"

// Drive drives the car.
func (c *Car) Drive() {
	if c.batteryDrain <= c.battery {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance displays the distance.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery displays the battery.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish tells you if car can finish the track.
func (c *Car) CanFinish(trackDistance int) bool {
	chunks := c.battery / c.batteryDrain
	travelled := c.speed * chunks
	return trackDistance <= travelled
}
