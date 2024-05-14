package splitter

type RegistryEntry struct {
	Name string
	ChildEntrys []RegistryEntry
	Tokens []string
}

func (re RegistryEntry) CreateMatrix() [][]string {
	if re.ChildEntrys != nil {
		mat := [][]string{}
		for _, child := range re.ChildEntrys {
			mat = addToMatrix(mat, child.CreateMatrix())
		}
		return mat
	} else {
		return [][]string{re.Tokens}
	}
}

