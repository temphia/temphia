package singleton

type Singleton struct {
	BFiles    map[string][]byte
	Agents    map[string]any
	Resources map[string]any
}
