package adapter

import "github.com/zcong1993/badge"

// BadgeInput is input type for badge generator
type BadgeInput badge.Input

// ErrorInput is default input when error occurred
var ErrorInput = BadgeInput{
	Subject: "error",
	Status:  "api error",
	Color:   "red",
}

// ApiFunc is function type api provider should implement
type ApiFunc func(args ...string) BadgeInput
