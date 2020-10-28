/*
#	ConcurrentRead.go
#	Written by Kyle S. && Mike D.
#	Last edit 10/26/20
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
var upperBound = 40

const targetValue = 25

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
		fmt.Println(lidarReading)
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
	gopigo3.SetLED(g.LED_EYE_LEFT+g.LED_EYE_RIGHT, 255, 0, 0)
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
	time.Sleep(time.Second)
	gopigo3.Halt()
}

func measureForward(gopigo3 *g.Driver) float64 { //Measures along the side of a box until reaching end
	side := 0.00
	// set indicator light
	gopigo3.SetLED(g.LED_EYE_LEFT+g.LED_EYE_RIGHT, 0, 0, 255)
	start := time.Now()
	feedbackControl(gopigo3)
	//gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, measureDPS)
	for {
		//wait until not reading object
		if !*&isReadingObject {
			duration := time.Since(start)
			side = duration.Seconds() * float64(measureDPS) * .05803
			gopigo3.SetLED(g.LED_EYE_LEFT+g.LED_EYE_RIGHT, 255, 0, 0)
			return side
		}
	}
}

func stepAndRotate(gopigo3 *g.Driver) { //Drives forward, then makes a 90 degree turn
	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, measureDPS*2)
	time.Sleep(time.Second * 3)

	//90 degree rotation
	gopigo3.SetMotorDps(g.MOTOR_LEFT, -113)
	gopigo3.SetMotorDps(g.MOTOR_RIGHT, 113)
	time.Sleep(time.Second * 2)
	gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, measureDPS)
	time.Sleep(time.Second / 2)
	gopigo3.Halt()

}

func feedbackControl(gopigo3 *g.Driver) { //Adjustment based on lidar readings
	output := targetValue - lidarReading
	if output > 5 {
		gopigo3.SetMotorDps(g.MOTOR_LEFT, measureDPS+output)
		gopigo3.SetMotorDps(g.MOTOR_RIGHT, measureDPS)
	} else if output < -5 {
		gopigo3.SetMotorDps(g.MOTOR_LEFT, measureDPS)
		gopigo3.SetMotorDps(g.MOTOR_RIGHT, measureDPS-output)
	} else {
		gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, measureDPS)
	}

}

func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver) {
	go setIsReading(lidarSensor)

	for { // while true
		if finished {
			gopigo3.Halt()
			break
			os.Exit(1)
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
	fmt.Println("The Length of the box is: ", length, "cm.")
	fmt.Print("The Width of the box is: ", width, "cm.")
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
