package main

func safeGetError(err error) string {
	if err != nil {
		return err.Error()
	}
	return "(No details)"
}

func strEsape(input string) string {
	escaped := ""
	for _, char := range input {
		switch char {
		case '\n':
			escaped += "&#10;"
		case '\r':
			escaped += "&#13;"
		case '\t':
			escaped += "&#9;"
		case ' ':
			escaped += "&#32;"
		default:
			escaped += string(char)
		}
	}
	return escaped
}