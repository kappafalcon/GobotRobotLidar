package main

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
//var upperBound = 25
//var checker = 0
//var halfSeconds = 0
//var length = 0.00
//var width = 0.00
//var DPS = 50
//var measuredSides = 0
//
//func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,
//
//) {
//
//	for { //check to make sure lidar sensor exists / has no issues
//		err := lidarSensor.Start()
//		if err != nil {
//			fmt.Println("error starting lidarSensor")
//		}
//
//		// get lidar reading
//		lidarReading, err := lidarSensor.Distance()
//		if err != nil {
//			fmt.Println("Error reading lidar sensor %+v", err)
//		}
//		// Define message with reading
//		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)
//
//		if (upperBound > lidarReading) && (lidarReading > lowerBound) { //If value suggests object, get reading and continue to drive.
//			isReadingObject = true
//			fmt.Println(message)
//			measureForward(gopigo3)
//			halfSeconds += 1
//			if checker == 0 {
//				length = float64(halfSeconds) * float64(DPS/2) * .05803
//			} else if checker == 1 {
//				width = float64(halfSeconds) * float64(DPS/2) * .05803
//
//			} else if (upperBound - lidarReading) < 0 { //Adjustments to turning if not in bounds
//				gopigo3.SetMotorDps(g.MOTOR_RIGHT, 5)
//			} else if (lowerBound - lidarReading) > 5 {
//				gopigo3.SetMotorDps(g.MOTOR_LEFT, 5)
//			}
//			measuredSides = 1
//
//		} else if (lidarReading > upperBound) && (measuredSides > 0) {
//			halfSeconds = 0
//			checker = 1
//			gopigo3.Halt()
//			stepAndRotate(gopigo3)
//			seekForward(gopigo3)
//
//		} else { // Move forward to seek object
//			seekForward(gopigo3)
//			fmt.Println(message)
//			fmt.Println("Seeking...")
//		}
//	}
//	fmt.Print("Length of side 1 equals: ", length, "cm")
//	fmt.Print("Width of side 2 equals: ", width, "cm")
//}
//
//func seekForward(gopigo3 *g.Driver) { // drive forward for one second
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
//	time.Sleep(time.Second)
//	gopigo3.Halt()
//}
//
//func measureForward(gopigo3 *g.Driver) { // drive forward for one second
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, 100)
//	gopigo3.SetLED(g.LED_EYE_LEFT+g.LED_EYE_RIGHT, 0, 0, 255)
//	time.Sleep(time.Second)
//	gopigo3.Halt()
//}
//
//func stepAndRotate(gopigo3 *g.Driver) {
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT+g.MOTOR_LEFT, DPS*2)
//	time.Sleep(time.Second * 3)
//	//90 degree rotation
//	gopigo3.SetMotorDps(g.MOTOR_LEFT, -110)
//	gopigo3.SetMotorDps(g.MOTOR_RIGHT, 110)
//	time.Sleep(time.Second * 2)
//	gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, DPS)
//	time.Sleep(time.Second / 2)
//	gopigo3.Halt()
//
//}
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
