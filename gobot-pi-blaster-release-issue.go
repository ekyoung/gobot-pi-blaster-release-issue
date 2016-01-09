package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor("raspi")
	led := gpio.NewLedDriver(r, "led", "12")

	work := func() {
		current := uint8(0)
		change := uint8(5)

		gobot.Every(30*time.Millisecond, func() {
			led.Brightness(current)
			current += change
			if current == 0 || current == 255 {
				change = -change
			}
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
