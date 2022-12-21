package cmw

type Type uint

const (
	JSONArray = Type(iota)
	CBORArray
	CBORTag
	Unknown
)

func identify(b []byte) Type {
	if len(b) == 0 {
		return Unknown
	}

	switch b[0] {
	case 0x82:
		return CBORArray
	case 0xda:
		return CBORTag
	case 0x5b:
		return JSONArray
	}

	return Unknown
}
