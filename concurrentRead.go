/*
#	ConcurrentRead.go
#	Written by Kyle S. && Mike D.
#	Last edit 10/20/20
*/
package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"os"
	"time"
)

var isReadingObject = false
var lowerBound = 10
var upperBound = 25
var measureDPS = 50
var lidarReading = 0
var length = 0.00
var width = 0.00
var finished = false

func setIsReading(lidarSensor *i2c.LIDARLiteDriver) {
	//check to make sure lidar sensor exists / has no issues; Kick out otherwise
	err := lidarSensor.Start()
	if err != nil {
		fmt.Println("error starting lidarSensor")
		fmt.Println("FATAL ERROR! \nExiting...")
		os.Exit(1)
	}
	for { //While true

		//get lidar reading
		*&lidarReading, err = lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}

		if (upperBound > *&lidarReading) && (*&lidarReading > lowerBound) { // if lidar suggests object, isReading = true
			if *&isReadingObject {
				continue
			} else {
				*&isReadingObject = true
			}
		} else { // if lidar suggests no object, isReading = false
			*&isReadingObject = false
		}
	}
}

func seekForward(gopigo3 *g.Driver) { // drive forward for one second
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
	time.Sleep(time.Second)
	gopigo3.Halt()
}

func measureForward(gopigo3 *g.Driver) float64 {
	side = 0.00
	// set indicator light
	gopigo3.SetLED(g.LED_EYE_LEFT+g.LED_EYE_RIGHT, 0, 0, 255)
	start := time.Now()
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, measureDPS)
	for {
		//wait until not reading object
		if !*&isReadingObject {
			duration := time.Since(start)
			side = duration.Seconds() * float64(measureDPS) * .05803
			return side
		}
	}
}

func stepAndRotate(gopigo3 *g.Driver) {
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, DPS*2)
	time.Sleep(time.Second * 3)

	//90 degree rotation
	gopigo3.SetMotorDps(g.MOTOR_LEFT, -110)
	gopigo3.SetMotorDps(g.MOTOR_RIGHT, 110)
	time.Sleep(time.Second * 2)
	gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, DPS)
	time.Sleep(time.Second / 2)
	gopigo3.Halt()
}

func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver) {
	go setIsReading(lidarSensor)

	for { // while true
		if finished {
			break
		} //both sides set. Time to end the program.

		if *&isReadingObject {
			if length == 0 {
				length = measureForward(gopigo3)
				stepAndRotate(gopigo3)
			} else if (length > 0) && width == 0 {
				width = measureForward(gopigo3)
				finished = true
			}

		} else {
			seekForward(gopigo3)
		}
	}
	fmt.Print("Deez dem sidez, boah")
	fmt.Print("length: ", length)
	fmt.Print("width: ", width)
}

func main() {
	raspberryPi := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspberryPi)
	lidarSensor := i2c.NewLIDARLiteDriver(raspberryPi)
	lightSensor := aio.NewGroveLightSensorDriver(gopigo3, "AD_2_1")
	workerThread := func() {
		robotMainLoop(raspberryPi, gopigo3, lidarSensor)
	}
	robot := gobot.NewRobot("Gopigo Pi4 Bot",
		[]gobot.Connection{raspberryPi},
		[]gobot.Device{gopigo3, lidarSensor, lightSensor},
		workerThread)

	robot.Start()
}
