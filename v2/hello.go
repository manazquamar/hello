package hello

// SayHello is CHANGED in v2: it now takes first and last name (breaking change)
func SayHello(first, last string) string {
	return "Hello, " + first + " " + last + "!"
}

// Welcome is also available in v2 (could behave slightly differently if you want)
func Welcome(name string) string {
	return "Welcome, " + name + "!"
}
