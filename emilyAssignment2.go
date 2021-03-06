package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func robotMainLoopTest(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,

) {

	err := lidarSensor.Start()
	if err != nil {
		fmt.Println("error starting lidarSensor")
	}

	for { //loop forever
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)
		speed(gopigo3)
		fmt.Println(lidarReading)
		fmt.Println(message)
		time.Sleep(time.Second * 3)
	}
}

func speed(gopigo3 *g.Driver) {
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 1000)
	time.Sleep(time.Second)
	gopigo3.Halt()

}

func main() {
	raspberryPi := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspberryPi)
	lidarSensor := i2c.NewLIDARLiteDriver(raspberryPi)
	lightSensor := aio.NewGroveLightSensorDriver(gopigo3, "AD_2_1")
	workerThread := func() {
		robotMainLoopTest(raspberryPi, gopigo3, lidarSensor)
	}
	robot := gobot.NewRobot("Gopigo Pi4 Bot",
		[]gobot.Connection{raspberryPi},
		[]gobot.Device{gopigo3, lidarSensor, lightSensor},
		workerThread)

	robot.Start()

}
