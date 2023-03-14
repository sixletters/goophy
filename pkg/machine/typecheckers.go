package machine

func isNumber(value interface{}) bool {
    switch value.(type) {
    case int, int32, int64, float32, float64:
        return true
    default:
        return false
    }
}

func isString(value interface{}) bool {
    _, ok := value.(string)
    return ok
}
func isBool(value interface{}) bool {
    _, ok := value.(bool)
    return ok
}