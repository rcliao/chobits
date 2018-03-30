package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/rcliao/tachikoma/views"
)

var (
	currentTurn    = 1
	totalTurns     = 5
	shortBreak     = 5 * time.Minute
	longBreak      = 15 * time.Minute
	pomodoroPeriod = 25 * time.Minute
)

func pomodoroTurn(chanPomodoro chan bool) {
	say("Work work")
	time.Sleep(pomodoroPeriod)
	say("Stop work")
	chanPomodoro <- true
}

func pomodoroBreak(chanBreak chan bool, t time.Duration) {
	say("Small break start")
	time.Sleep(t)
	say("Small break stop")
	chanBreak <- true
}

func say(message string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(message)
}

func pomodoroService(chanPomodoro, chanBreak, chanLongBreak, chanDone chan bool) {
	fmt.Println("Pomodoro started")
	for {
		select {
		case <-chanPomodoro:
			if currentTurn >= totalTurns {
				go pomodoroBreak(chanLongBreak, longBreak)
				currentTurn = 1
			} else {
				currentTurn++
				go pomodoroBreak(chanBreak, shortBreak)
			}
		case <-chanBreak:
			go pomodoroTurn(chanPomodoro)
		case <-chanLongBreak:
			input := askAnotherSession()
			for input != "Y" && input != "N" {
				input = askAnotherSession()
			}
			if input == "Y" {
				go pomodoroTurn(chanPomodoro)
			} else {
				chanDone <- true
			}
		}
	}
}

func askAnotherSession() string {
	fmt.Println("Ready for another pomodoro session? (Y/N)")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	drawer, _ := views.NewTerminalView()

	drawer.DrawMain(views.ConvertClockToMain("20:55"))
}
