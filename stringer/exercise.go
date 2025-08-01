package meteorology
import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (t TemperatureUnit) String() string {
    units := []string{"°C","°F"}
    return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
    return fmt.Sprintf("%v %v",t.degree,t.unit)
}

// Add a String method to SpeedUnit
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)


func (s SpeedUnit) String() string {
    units := []string{"km/h","mph"}
    return units[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}
// Add a String method to Speed
func (s Speed) String() string {
    return fmt.Sprintf("%d %s",s.magnitude,s.unit)
}



// Add a String method to MeteorologyData
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",m.location,m.temperature,m.windDirection,m.windSpeed,m.humidity)
}