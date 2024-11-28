package iteration


const count = 5

func Repeat(char string) string{
	var repeated string
	for i := 0; i < count; i++ {
		repeated = repeated + char
	}
	return repeated
}