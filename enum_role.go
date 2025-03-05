package client

import "fmt"

type roleSet struct {
	Assistant Role
	User      Role
	System    Role
}

var Roles = roleSet{
	Assistant: newRole("assistant"),
	User:      newRole("user"),
	System:    newRole("system"),
}

func (roleSet) Parse(value string) (Role, error) {
	role, exists := roles[value]
	if !exists {
		return Role{}, fmt.Errorf("invalid role %q", value)
	}

	return role, nil
}

func (roleSet) MustParse(value string) Role {
	role, err := Roles.Parse(value)
	if err != nil {
		panic(err)
	}

	return role
}

// =============================================================================

var roles = make(map[string]Role)

type Role struct {
	value string
}

func newRole(role string) Role {
	r := Role{role}
	roles[role] = r
	return r
}

func (r Role) String() string {
	return r.value
}

func (r *Role) UnmarshalText(data []byte) error {
	role, err := Roles.Parse(string(data))
	if err != nil {
		return err
	}

	r.value = role.value
	return nil
}

func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.value), nil
}

func (r Role) Equal(r2 Role) bool {
	return r.value == r2.value
}
