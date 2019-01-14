package generalG

// check if variable is an empty variable
func isEmpty(i interface{}) bool {
	switch i.(type) {
	default:
		break
	case string:
		if i.(string) == "" {
			return true
		}
		break
	case int:
		if i.(int) == 0 {
			return true
		}
		break
	}
	return false
}
