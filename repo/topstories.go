package repo

func (host Host) GetTopStoryIds() ([]int, error) {
	ids := []int{}
	err := host.LoadJson("/v0/topstories.json", &ids)
	return ids, err
}
