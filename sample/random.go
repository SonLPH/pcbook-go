package sample

import (
	"math/rand"
	"time"

	"github.com/SonLPH/pcbook-go/pb"
	"github.com/google/uuid"
)

func init() {
	rand.New(rand.NewSource((time.Now().UnixNano())))
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E", "Core i9", "Core i7", "Core i5", "Core i3",
		)
	}
	return randomStringFromSet(
		"Ryzen 7", "Ryzen 5", "Ryzen 3",
	)
}

func randomStringFromSet(set ...string) string {
	n := len(set)
	if n == 0 {
		return ""
	}
	return set[rand.Intn(n)]
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NIVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NIVIDIA" {
		return randomStringFromSet(
			"RTX 2060", "RTX 2070", "RTX 2080", "GTX 1660Ti",
		)
	}
	return randomStringFromSet(
		"RX 590", "RX 580", "RX 570", "RX 560",
	)
}

func randomScreenSolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Width:  int32(width),
		Height: int32(height),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1")
	}
}
