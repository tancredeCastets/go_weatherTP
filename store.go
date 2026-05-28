package weather

type Store struct {
	stations map[string]Station
}

func NewStore() *Store {
	return &Store{make(map[string]Station)}
}

func (s *Store) Put(st Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok

}

func (s *Store) Get(id string) (Station, bool) {
	st, ok := s.stations[id]
	return st, ok

}

func (s *Store) Delete(id string) bool {
	_, ok := s.stations[id]
	delete(s.stations, id)
	return ok

}

func (s *Store) All() []Station {
	sts := make([]Station, 0, len(s.stations))
	for _, st := range s.stations {
		sts = append(sts, st)
	}
	return sts

}
