package track

type ProcessorSorter struct {
	Data []ProcessorInterface
}

func (sorter ProcessorSorter) Len() int {
	return len(sorter.Data)
}

func (sorter ProcessorSorter) Swap(i int, j int) {
	sorter.Data[i], sorter.Data[j] = sorter.Data[j], sorter.Data[i]
}

func (sorter ProcessorSorter) Less(i int, j int) bool {
	return sorter.Data[i].GetPriority() < sorter.Data[j].GetPriority()
}
