package generalG

// check if variable is an empty variable
func isEmpty(i interface{}) bool {
	switch i.(type) {
	default:
		break
	case string:
		if i.(string) != "" {
			return false
		}
		break
	case float64:
		if i.(float64) != 0 {
			return false
		}
		break
	case int:
		if i.(int) != 0 {
			return false
		}
		break
	}
	return true
}
