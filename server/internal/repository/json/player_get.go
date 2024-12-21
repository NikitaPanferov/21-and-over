package json

func (r *Repo) GetPlayer(name string) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	chips, ok := r.players[name]
	if ok {
		return chips
	}

	r.players[name] = r.initialChips

	return r.initialChips
}
