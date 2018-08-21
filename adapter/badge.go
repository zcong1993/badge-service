package adapter

// BadgeApi is api for badgen directly
func BadgeApi(args ...string) BadgeInput {
	if len(args) != 3 || args[0] == "" || args[1] == "" {
		return ErrorInput
	}

	return BadgeInput{
		Subject: args[0],
		Status:  args[1],
		Color:   args[2],
	}
}
