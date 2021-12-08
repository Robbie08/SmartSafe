package piUtils

import (
	"fmt"
	"log"
	"time"

	"github.com/cgxeiji/servo"
)

func UnlockSafe() {
	PIN := 14 // PIN number of where the sevo signal is connected
	fmt.Println("Unlocking Safe...")

	// rotate servo motor 4 times
	for i := 0; i < 1; i++ {
		rotateServo(PIN)
	}
}

func rotateServo(pin int) {
	oServo := servo.New(pin) // create new servo struct for pin 14

	/* Declare servo struct information */

	oServo.MinPulse = 0.05    // sets the Min PWM pulse of the servo (this is the default)
	oServo.MaxPulse = 0.25    // sets the Max PWN pulse of the servo (this is the default)
	oServo.Name = "ServoLock" // set the name of the servo
	oServo.SetPosition(90)    // the starting poition of the servo hand
	oServo.SetSpeed(0.8)      // set our rotation speed of 20%

	fmt.Println(".")

	err := oServo.Connect() // Connect Servo instance to pi-blaster daemon

	// handle any errors while connecting
	if err != nil {
		log.Fatal(err)
	}

	// close any connection to the pin
	defer oServo.Close()

	oServo.SetSpeed(0.8)
	oServo.MoveTo(90) // no-blocking will rotate 180 degrees from start
	oServo.Wait()     // will allow sync with servo
	time.Sleep(20 * time.Second)
	oServo.MoveTo(0).Wait()

}
