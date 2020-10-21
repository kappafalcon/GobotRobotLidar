package main

//
//
//import (
//	"fmt"
//	"gobot.io/x/gobot"
//	"gobot.io/x/gobot/drivers/aio"
//	"gobot.io/x/gobot/drivers/i2c"
//	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
//	"gobot.io/x/gobot/platforms/raspi"
//	"time"
//)
//
//var isReadingObject = false
//var lowerBound = 10
//var upperBound = 50
//var measureDPS = 50
//var side1Seconds float64 = 0
//var side2Seconds float64 = 0
//
//func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,
//
//) {
//	//check to make sure lidar sensor exists / has no issues
//	err := lidarSensor.Start()
//	if err != nil {
//		fmt.Println("error starting lidarSensor")
//	}
//
//	for {
//
//		if side2Seconds != 0 { //if program has set side2seconds, kick out of forever and add the 2 sides.
//			break
//		}
//
//		// get lidar reading
//		lidarReading, err := lidarSensor.Distance()
//		if err != nil {
//			fmt.Println("Error reading lidar sensor %+v", err)
//		}
//
//		// Define message with reading
//		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)
//
//		if (upperBound > lidarReading) && (lidarReading > lowerBound) && (isReadingObject == false) { //If value suggests object && !isReading.
//			isReadingObject = true
//			fmt.Println(message)
//
//		} else if (upperBound > lidarReading) && (lidarReading > lowerBound) && (isReadingObject == true) {
//			continue
//		} else { // Move forward to seek object
//			seekForward(gopigo3)
//			fmt.Println(message)
//		}
//
//	}
//}
//func seekForward(gopigo3 *g.Driver) { // drive forward for one second
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
//	time.Sleep(time.Second)
//	gopigo3.Halt()
//}
//
//func measureForward(gopigo3 *g.Driver) {
//
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
//	time.Sleep(time.Second)
//	gopigo3.Halt()
//}
//
//func main() {
//	raspberryPi := raspi.NewAdaptor()
//	gopigo3 := g.NewDriver(raspberryPi)
//	lidarSensor := i2c.NewLIDARLiteDriver(raspberryPi)
//	lightSensor := aio.NewGroveLightSensorDriver(gopigo3, "AD_2_1")
//	workerThread := func() {
//		robotMainLoop(raspberryPi, gopigo3, lidarSensor)
//	}
//	robot := gobot.NewRobot("Gopigo Pi4 Bot",
//		[]gobot.Connection{raspberryPi},
//		[]gobot.Device{gopigo3, lidarSensor, lightSensor},
//		workerThread)
//
//	robot.Start()
//
//}
