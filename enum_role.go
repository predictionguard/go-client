package client

import "fmt"

type roleSet struct {
	Assistant Role
	User      Role
	System    Role
}

// Roles represents the set of roles that can be used.
var Roles = roleSet{
	Assistant: newRole("assistant"),
	User:      newRole("user"),
	System:    newRole("system"),
}

// Parse parses the string value and returns a role if one exists.
func (roleSet) Parse(value string) (Role, error) {
	role, exists := roles[value]
	if !exists {
		return Role{}, fmt.Errorf("invalid role %q", value)
	}

	return role, nil
}

// MustParse parses the string value and returns a role if one exists. If
// an error occurs the function panics.
func (roleSet) MustParse(value string) Role {
	role, err := Roles.Parse(value)
	if err != nil {
		panic(err)
	}

	return role
}

// =============================================================================

// Set of known roles.
var roles = make(map[string]Role)

// Role represents a role in the system.
type Role struct {
	value string
}

func newRole(role string) Role {
	r := Role{role}
	roles[role] = r
	return r
}

// String returns the name of the role.
func (r Role) String() string {
	return r.value
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (r *Role) UnmarshalText(data []byte) error {
	role, err := Roles.Parse(string(data))
	if err != nil {
		return err
	}

	r.value = role.value
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.value), nil
}

// Equal provides support for the go-cmp package and testing.
func (r Role) Equal(r2 Role) bool {
	return r.value == r2.value
}
