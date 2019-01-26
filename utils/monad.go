package utils

//TODO: POC, refactor it
func Monad(result ...interface{}) (interface{}, error) {
	var f []interface{}

	for _, r := range result {
		switch r.(type) {
		case error:
			return nil, r.(error)
		default:
			f = append(f, r)
		}
	}

	return f, nil
}
