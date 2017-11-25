package eugddc

type DogDict map[string][]Dog

func (d DogDict) ToSlice() []Dog {
	slice := make([]Dog, 0, len(d))
	for _, value := range d {
		slice = append(slice, value...)
	}
	return slice
}
