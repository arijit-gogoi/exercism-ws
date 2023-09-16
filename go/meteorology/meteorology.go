package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String method to the TemperatureUnit type.
func (t TemperatureUnit) String() (s string) {
	// switch t {
	// case Fahrenheit:
	// 	s = "°F"
	// case Celsius:
	// 	s = "°C"
	// }
	// return s
	units := [2]string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String method to the Temperature type.
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String method to SpeedUnit.
func (su SpeedUnit) String() (s string) {
	// switch sp {
	// case KmPerHour:
	// 	s = "km/h"
	// case MilesPerHour:
	// 	s = "mph"
	// }
	// return s
	units := [2]string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String method to Speed.
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String method to MeteorologyData.
func (s MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		s.location, s.temperature, s.windDirection, s.windSpeed, s.humidity)
}
