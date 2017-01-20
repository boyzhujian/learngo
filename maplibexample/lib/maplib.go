package lib

var language map[string]string

func init() {
	language = make(map[string]string)
	language["cs"] = "C#"
	language["js"] = "javascript"
}
func Get(key string) string {
	return language[key]
}
func Add(key string, value string) string {
	language[key] = value
	return ""
}

func GetAll() map[string]string {
	return language
}
