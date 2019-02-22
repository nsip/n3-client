package xjy

// Node is
type Node struct {
	/* from yaml */
	tag        string
	value      string
	path       string
	level      int
	levelXPath []int /* index is the level, value is the line number */
	aevalue    bool
	id         string
}
