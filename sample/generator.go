package sample

import (
	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/golang/protobuf/ptypes"
)

func NewLaptop() *message.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &message.Laptop{
		Id:       randomId(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCpu(),
		Memory:   NewRAM(),
		Gpu:      []*message.GPU{NewGpu()},
		Storage:  []*message.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &message.Laptop_Kg{
			Kg: randomFloat32(1.0, 3.0),
		},
		PriceRs:   randomFloat32(1500, 3500),
		Year:      uint32(randomInt(2015, 2019)),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}

func NewKeyboard() *message.Keyboard {
	keyboard := &message.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCpu() *message.CPU {
	brand := randomCpuBrand()
	name := randomModernCPUName(brand)

	cores := randomInt(6, 196)
	threads := randomInt(cores, cores*2)

	min_freq := randomFloat64(2.0, 4.0)
	max_freq := randomFloat64(min_freq, 6.0)

	cpu := &message.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MinGhz:        min_freq,
		MaxGhz:        max_freq,
	}
	return cpu

}

func NewGpu() *message.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	min_freq := randomFloat64(1.0, 1.5)
	max_freq := randomFloat64(min_freq, 3.0)
	memGb := randomInt(6, 24)

	gpu := &message.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: min_freq,
		MaxGhz: max_freq,
		Memory: &message.Memory{
			Value: float64(memGb),
			Unit:  message.Memory_GIGABYTE,
		},
	}
	return gpu
}

func NewRAM() *message.Memory {
	memGB := randomInt(4, 64)

	ram := &message.Memory{
		Value: float64(memGB),
		Unit:  message.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *message.Storage {
	memGB := randomInt(128, 1024)

	ssd := &message.Storage{
		Driver: message.Storage_SSD,
		Memory: &message.Memory{
			Value: float64(memGB),
			Unit:  message.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *message.Storage {
	memGb := randomInt(512, 8000)

	hdd := &message.Storage{
		Driver: message.Storage_HDD,
		Memory: &message.Memory{
			Value: float64(memGb),
			Unit:  message.Memory_GIGABYTE,
		},
	}

	return hdd
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomString(
			"RTX 4090",
			"RTX 4080",
			"RTX 4070",
			"RTX 4060",
		)
	}

	return randomString(
		"RX 7900 XT",
		"RX 7800 XT",
		"RX 7700 XT",
		"RX 7600 XT",
	)
}

func NewScreen() *message.Screen {
	screen := &message.Screen{
		Size:       randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Touch:      randomBool(),
	}

	return screen
}
