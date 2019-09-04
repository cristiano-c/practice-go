// Package triangle provides simple gemetry calculations on triangles
package triangle

// Kind represents one of the four const triangle types
type Kind string

const (
	// NaT means not a triangle
	NaT string = "NaT"
	// Equ means equilateral
	Equ string = "Equ"
	// Iso means isosceles
	Iso string = "Iso"
	// Sca means scalene
	Sca string = "Sca" // scalene
)

// KindFromSides returns the kind of a triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a+b <= c || a+c <= b || b+c <= a {
		if a == b && b == c {
			k = Kind(Equ)
		} else if a == b || b == c || a == c {
			k = Kind(Iso)
		} else {
			k = Kind(Sca)
		}
	} else {
		k = Kind(NaT)
	}
	return k
}
