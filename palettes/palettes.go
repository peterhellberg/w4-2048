package palettes

type Palette [4]uint32

var (
	BluesGB     = Palette{0xe5f1f3, 0x7ba8b8, 0x30617b, 0x08263b} // https://lospec.com/palette-list/blues-gb (reversed)
	EN4         = Palette{0xfbf7f3, 0xe5b083, 0x426e5d, 0x20283d} // https://lospec.com/palette-list/en4
	GBChocolate = Palette{0xffe4c2, 0xdca456, 0xa9604c, 0x422936} // https://lospec.com/palette-list/gb-chocolate
	Grapefruit  = Palette{0xfff5dd, 0xf4b26b, 0xb76591, 0x65296c} // https://lospec.com/palette-list/grapefruit (reversed)
	GreyMist    = Palette{0xf1ffe0, 0x988171, 0x463534, 0x1e1721} // https://lospec.com/palette-list/grey-mist
	IceCreamGB  = Palette{0xfff6d3, 0xf9a875, 0xeb6b6f, 0x7c3f58} // https://lospec.com/palette-list/ice-cream-gb
	Keeby       = Palette{0xc5ccb8, 0x899b98, 0x5d6872, 0x333343} // https://lospec.com/palette-list/keeby
	Lightgreen  = Palette{0xf4fbd0, 0x68cf68, 0x1e9178, 0x05241f} // https://lospec.com/palette-list/lightgreen
	Platinum    = Palette{0xe0f0e8, 0xa8c0b0, 0x507868, 0x183030} // https://lospec.com/palette-list/platinum
	Warmlight   = Palette{0xffd191, 0xff924f, 0x66605c, 0x211e20} // https://lospec.com/palette-list/warmlight (reversed)

	All = []Palette{
		BluesGB,
		EN4,
		GBChocolate,
		Grapefruit,
		GreyMist,
		IceCreamGB,
		Keeby,
		Lightgreen,
		Platinum,
		Warmlight,
	}
)
