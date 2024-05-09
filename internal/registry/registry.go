package registry

type TokenRegistry struct{
	Map map[string]int64
}

func (tr TokenRegistry) DistributeTokens(list [][]string) {
	for _, file := range list {
		for _, value := range file {
			tr.Map[value] = int64(len(tr.Map))
		}
	}
}
