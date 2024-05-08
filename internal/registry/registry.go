package registry

type TokenRegistry struct{
	Map map[string]int64
}

func (tr TokenRegistry) DistributeTokens(list []string) {
	for _, v := range list {
		tr.Map[v] = int64(len(tr.Map))
	}
}
