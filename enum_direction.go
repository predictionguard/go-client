package client

import "fmt"

type directionSet struct {
	Right Direction
	Left  Direction
}

// Directions represents the set of directions.
var Directions = directionSet{
	Right: newDirection("Right"),
	Left:  newDirection("Left"),
}

// Parse parses the string value and returns a direction if one exists.
func (directionSet) Parse(value string) (Direction, error) {
	server, exists := directions[value]
	if !exists {
		return Direction{}, fmt.Errorf("invalid direction %q", value)
	}

	return server, nil
}

// MustParse parses the string value and returns a direction if one exists.
// If an error occurs the function panics.
func (directionSet) MustParse(value string) Direction {
	server, err := Directions.Parse(value)
	if err != nil {
		panic(err)
	}

	return server
}

// =============================================================================

// Set of known directions.
var directions = make(map[string]Direction)

// Direction represents a direction in the system.
type Direction struct {
	value string
}

func newDirection(direction string) Direction {
	d := Direction{direction}
	directions[direction] = d
	return d
}

// String returns the name of the direction.
func (d Direction) String() string {
	return d.value
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (d *Direction) UnmarshalText(data []byte) error {
	direction, err := Directions.Parse(string(data))
	if err != nil {
		return err
	}

	d.value = direction.value
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (d Direction) MarshalText() ([]byte, error) {
	return []byte(d.value), nil
}

// Equal provides support for the go-cmp package and testing.
func (d Direction) Equal(d2 Direction) bool {
	return d.value == d2.value
}
