package discovery

type IgnoreRules struct {
	Names map[string]struct{}
}
func (r *IgnoreRules) Contains(name string) bool
