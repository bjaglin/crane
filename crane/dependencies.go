package crane

// Dependencies contains 4 fields:
// all: contains all dependencies
// link: containers linked to
// volumesFrom: containers that provide volumes
// net: container the net stack is shared with
type Dependencies struct {
	All         []string
	Link        []string
	VolumesFrom []string
	Net         string
}

// includes checks whether the given needle is
// included in the dependency list
func (d *Dependencies) includes(needle string) bool {
	return d.includesAsKind(needle, "all")
}

// includes checks whether the given needle is
// included in the dependency list as the given kind
func (d *Dependencies) includesAsKind(needle string, kind string) bool {
	for _, name := range d.forKind(kind) {
		if name == needle {
			return true
		}
	}
	return false
}

// returns the list of dependencies for a certain
// kind of dependency
func (d *Dependencies) forKind(kind string) []string {
	switch kind {
	case "all":
		return d.All
	case "link":
		return d.Link
	case "volumesFrom":
		return d.VolumesFrom
	case "net":
		return []string{d.Net}
	default:
		return []string{}
	}
}

// mustRun checks whether the given needle needs
// to be running
func (d *Dependencies) mustRun(needle string) bool {
	if needle == d.Net {
		return true
	}
	for _, name := range d.Link {
		if name == needle {
			return true
		}
	}
	return false
}

// satisfied is true when there are no
// dependencies
func (d *Dependencies) satisfied() bool {
	return len(d.All) == 0
}

// remove removes the given name from all
func (d *Dependencies) remove(resolved string) {
	for i, name := range d.All {
		if name == resolved {
			d.All = append(d.All[:i], d.All[i+1:]...)
		}
	}
}
