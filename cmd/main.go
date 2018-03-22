package main

import (
	"fmt"
	"os/exec"
	"time"
)

var (
	currentTurn = 1
	totalTurns  = 5
	shortBreak  = 5 * time.Minute
	longBreak   = 15 * time.Minute
)

func pomodoroTurn(chanPomodoro chan bool) {
	say("Work work")
	time.Sleep(time.Minute * 25)
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
	exec.Command("say", message).Output()
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
	turn := make(chan bool)
	smallBreak := make(chan bool)
	longBreak := make(chan bool)
	done := make(chan bool)

	go pomodoroTurn(turn)
	go pomodoroService(turn, smallBreak, longBreak, done)

	<-done
}
