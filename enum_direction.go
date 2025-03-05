package client

import "fmt"

type directionSet struct {
	Right Direction
	Left  Direction
}

var Directions = directionSet{
	Right: newDirection("Right"),
	Left:  newDirection("Left"),
}

func (directionSet) Parse(value string) (Direction, error) {
	server, exists := directions[value]
	if !exists {
		return Direction{}, fmt.Errorf("invalid direction %q", value)
	}

	return server, nil
}

func (directionSet) MustParse(value string) Direction {
	server, err := Directions.Parse(value)
	if err != nil {
		panic(err)
	}

	return server
}

// =============================================================================

var directions = make(map[string]Direction)

type Direction struct {
	value string
}

func newDirection(direction string) Direction {
	d := Direction{direction}
	directions[direction] = d
	return d
}

func (d Direction) String() string {
	return d.value
}

func (d *Direction) UnmarshalText(data []byte) error {
	direction, err := Directions.Parse(string(data))
	if err != nil {
		return err
	}

	d.value = direction.value
	return nil
}

func (d Direction) MarshalText() ([]byte, error) {
	return []byte(d.value), nil
}

func (d Direction) Equal(d2 Direction) bool {
	return d.value == d2.value
}
