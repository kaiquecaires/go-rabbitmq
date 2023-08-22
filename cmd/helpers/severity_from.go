package helpers

func SeverityFrom(args []string) string {
	var s string
	if (len(args) < 2) || args[1] == "" {
		s = "info"
	} else {
		s = args[1]
	}
	return s
}
