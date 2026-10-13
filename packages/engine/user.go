package engine

// import (
// 	"fmt"
// )

// defines a single user action which is
// used to build a UserProfile
// data is an optional field which is only required if
// the user wishes to stress test their database
type UserAction struct {
	Name     string         `json:"name"`
	Endpoint string         `json:"endpoint"`
	Data     map[string]any `json:"data,omitempty"`
}

// a UserProfile is a collection of UserActions which
// can be defined to have a stastical odd of happening
type UserProfile struct {
	Name        string       `json:"name"`
	UserActions []UserAction `json:"userActions"`
	ActionOdds  []float32    `json:"actionOdds"`
}
