package main

type IShirt interface {
	getLogo() string
	getSize() string
	setLogo(logo string)
	setSize(size string)
}

type Shirt struct {
	logo string
	size string
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size string) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() string {
	return s.size
}
