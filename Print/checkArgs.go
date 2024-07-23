package box

// cheack args and transform file
func CheckArgs(font string) string {
	switch font {
	case "thinkertoy":
		return "fs/thinkertoy.txt"

	case "shadow":
		return "fs/shadow.txt"

	case "standard":
		return "fs/standard.txt"

	default:
		return ""
	}
}
