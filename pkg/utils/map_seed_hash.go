package utils

import "math"

const mapHashDivisor = 1 << 16

func GetMapSeed(initHashSeed, endHashSeed uint) (uint, bool) {
	var gameSeedXor uint = 0
	seed, found := reverseMapSeedHash(endHashSeed)
	if found {
		gameSeedXor = initHashSeed ^ seed
	}

	if gameSeedXor == 0 {
		return 0, false
	}

	return seed, true
}

func reverseMapSeedHash(hash uint) (uint, bool) {
	incrementalValue := uint(1)

	for startValue := uint(0); startValue < math.MaxUint; startValue += incrementalValue {
		seedResult := (uint(startValue)*0x6AC690C5 + 666) & 0xFFFFFFFF

		if seedResult == hash {
			return startValue, true
		}

		if incrementalValue == 1 && (seedResult%mapHashDivisor) == (hash%mapHashDivisor) {
			incrementalValue = mapHashDivisor
		}
	}

	return 0, false
}
