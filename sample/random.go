package sample

import (
	"math/rand"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/google/uuid"
)

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomKeyboardLayout() message.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return message.Keyboard_AZERTY
	case 2:
		return message.Keyboard_QWERTY
	default:
		return message.Keyboard_QWERTZ
	}
}

func randomCpuBrand() string {
	return randomString("Intel", "AMD")
}

func randomString(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomModernCPUName(brand string) string {
	if brand == "Intel" {
		return randomString(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomString(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min-1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomGPUBrand() string {
	return randomString("Nvidia", "AMD")
}

func randomScreenResolution() *message.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &message.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() message.Screen_Panel {
	if rand.Intn(2) == 1 {
		return message.Screen_IPS
	}
	return message.Screen_OLED
}

func randomId() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomString("Apple", "Asus", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomString("MacBook Pro 16", "MacBook Pro 14", "MacBook Air M2")
	case "Asus":
		return randomString("ROG Zephyrus G14", "ZenBook Pro Duo", "ROG Strix Scar 15")
	default:
		return randomString("ThinkPad X1 Carbon", "ThinkPad X1 Extreme", "Legion 7i")
	}
}
