// Package twofer should have a package comment that summarizes what it's about.
// And it does.
package twofer

// ShareWith should have a comment documenting it. And it does.
func ShareWith(name string) string {
	who := ""
	if name == "" {
		who = "you"
	} else {
		who = name
	}

	return "One for " + who + ", one for me."
}
