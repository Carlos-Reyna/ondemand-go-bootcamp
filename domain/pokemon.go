package domain

type Pokemon struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (p *Pokemon) Init(id int, name string) {
	p.Id = id
	p.Name = name
}
