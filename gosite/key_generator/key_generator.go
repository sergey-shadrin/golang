package key_generator

import "github.com/speps/go-hashids"

func GenerateKeyById(id int64) (string, error) {
	hd := hashids.NewData()
	hd.Salt = "Shadrin"
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	return h.Encode([]int{int(id)})
}
