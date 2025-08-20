package hello

// SayHello returns a greeting message (kept same to avoid breaking v1 users).
func SayHello(name string) string {
	return "Hello, " + name
}

// Welcome is NEW in v1.1.0 (non-breaking addition).
// It keeps the old API working and just adds a new feature.
func Welcome(name string) string {
	return "Welcome, " + name
}
