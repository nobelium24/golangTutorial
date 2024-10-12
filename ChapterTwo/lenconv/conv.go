package lenconv

func FToM(f Feet) Meters {
	return Meters(f / 3.28084)
}

func MToF(m Meters) Feet {
	return Feet(m * 3.28084)
}
