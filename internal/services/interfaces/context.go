package interfaces

type Context interface {
	FormValue(string) string
	JSON(int, interface{}) error
	JSONPretty(int, interface{}, string) error
	String(int, string) error
}
