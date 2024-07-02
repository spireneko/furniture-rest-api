package model

type Furniture struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Fabricator string `json:"fabricator"`
	Height     uint32 `json:"height"`
	Width      uint32 `json:"width"`
	Length     uint32 `json:"length"`
}

func (f *Furniture) HasEmptyFields() bool {
	if len(f.Name) == 0 {
		return true
	}
	if len(f.Fabricator) == 0 {
		return true
	}
	if f.Height == 0 {
		return true
	}
	if f.Width == 0 {
		return true
	}
	if f.Length == 0 {
		return true
	}

	return false
}
