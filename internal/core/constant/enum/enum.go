package enum

type Enum int

const (
	// product Id
	WipingCloth Enum = iota

	// texture id
	Clear
	Matte
	Privacy
)

func (e Enum) String() string {
	return [...]string{
		"WIPING-CLOTH",
		"CLEAR",
		"MATTE",
		"PRIVACY",
	}[e]
}
