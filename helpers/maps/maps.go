package maps

type Numbers interface {
	byte
	int8
	int16
	int32
	int64
	int
	uint8
	uint16
	uint32
	uint64
	uint
	float32
}

func GetSmallestValue(m map[Numbers]int) (smallest int) {
	for _, smallest = range m {
		break
	}

	for _, v := range m {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}
