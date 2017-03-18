package font

import (
	"errors"
	"fmt"
)

type Font struct {
	fontFamily string
	size int
}

func (f *Font) SetFamily(family string) error {
	if family == "" {
		return errors.New("Font family can not be empty")
	}
	f.fontFamily = family
	return nil
}

func (f *Font) SetSize(size int) error {
	if !(size > 4 && size < 145) {
		return errors.New("Font size must be in range [5 .. 144]")
	}
	f.size = size
	return nil
}

func NewFont(family string, size int) *Font {
	defaultFont := &Font{"Arial", 12}
	defaultFont.SetFamily(family)
	defaultFont.SetSize(size)
	return defaultFont
}

func (f *Font) Family() string {
	return f.fontFamily
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) String() string {
	return fmt.Sprintf("{font-family: \"%v\"; font-size: %vpt;}", f.fontFamily, f.size)
}