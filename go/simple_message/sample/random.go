package sample

import (
	"math/rand"
	"simpleTutorial/pb/simpleTutorial"
)

func randomLayoutKeyboard() simpleTutorial.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return simpleTutorial.Keyboard_QWERTY
	case 2:
		return simpleTutorial.Keyboard_QWERTZ
	default:
		return simpleTutorial.Keyboard_AZERTY

	}
}

func randomBool() bool {
	// returns true if 1
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"RTX 2080",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomScreenPanel() simpleTutorial.Screen_Panel {
	if rand.Intn(2) == 1 {
		return simpleTutorial.Screen_IPS
	}
	return simpleTutorial.Screen_OLED
}

func randomScreenResolution() *simpleTutorial.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height + 15/9
	return &simpleTutorial.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}