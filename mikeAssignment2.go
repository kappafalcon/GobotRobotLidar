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

var isReadingObject = false
var lowerBound = 10
var upperBound = 25
var checker = 0
var seconds = 0
var length = 0.00

func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,

) {

	for { //check to make sure lidar sensor exists / has no issues
		err := lidarSensor.Start()
		if err != nil {
			fmt.Println("error starting lidarSensor")
		}

		// get lidar reading
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		// Define message with reading
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

		if (upperBound > lidarReading) && (lidarReading > lowerBound) { //If value suggests object, get reading and continue to drive.
			isReadingObject = true
			fmt.Println(message)
			measureForward(gopigo3)
			seconds += 1
			length = float64(seconds) * 50 * .05803
			checker = 1

		} else if (lidarReading < upperBound) && (checker > 0) {
			seconds += 1
			fmt.Print("Length equals:", length)
			gopigo3.SetMotorDps(g.MOTOR_RIGHT, 90) // ROTATE 90 DEGREES
			gopigo3.SetMotorDps(g.MOTOR_LEFT, 90)
			if seconds > 1 { // IF TIME LONGER THAN 1 SECOND STOP MOTORS
				gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 0)
			}
			gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100) // AFTER ROTATION DRIVE STRAIGHT
			seconds = 0

		} else { // Move forward to seek object
			seekForward(gopigo3)
			fmt.Println(message)
			fmt.Println("Seeking...")
		}

	}
}

func seekForward(gopigo3 *g.Driver) { // drive forward for one second
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
	time.Sleep(time.Second)
	gopigo3.Halt()
}

func measureForward(gopigo3 *g.Driver) { // drive forward for one second
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 50)
	time.Sleep(time.Second)

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
