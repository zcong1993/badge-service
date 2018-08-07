package adapter

import "github.com/zcong1993/badge"

type BadgeInput badge.Input

var ErrorInput = BadgeInput{
	Subject: "error",
	Status:  "api error",
	Color:   "red",
}
