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
var upperBound = 50
var measureDPS = 50
var side1Set = false
var side2Set = false
var side1Seconds float64 = 0
var side2Seconds float64 = 0

func setIsReading(lidarSensor *i2c.LIDARLiteDriver) {
	//check to make sure lidar sensor exists / has no issues; Kick out otherwise
	err := lidarSensor.Start()
	if err != nil {
		fmt.Println("error starting lidarSensor")
		fmt.Println("FATAL ERROR! \nExiting...")
		os.Exit(1)
	}

	for { //While true
		if (upperBound > lidarReading) && (lidarReading > lowerBound) { // if lidar suggests object, isReading = true
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

func timer(){
	start := time.Now()
	for{
		if *&isReadingObject{
			continue
		}else if !*&isReadingObject{

		}
	}
}


func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver) {
	go setIsReading(lidarSensor)

	for { // while true
		if side2Set {
			break
		} //both sides set. Time to end the program.

		if *&isReadingObject {
			continue
		}
	}
	// Add lengths of sides
	//side := float64(halfSeconds) * float64(DPS/2) * .05803
	//fmt.Println("The combined side length is: ", side)

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
