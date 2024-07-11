package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main() {

	// we want 3 arguments : path time msg, other than that throw error
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm> <text message\n>", os.Args[0])
		// Args[0] is path of project
		os.Exit(1) // if not pass argumemts the exit from programme
	}
	now := time.Now() // current time

	w := when.New(nil)
	w.Add(en.All...)     // add english language parsing rules
	w.Add(common.All...) // add common parsing rules

	t, err := w.Parse(os.Args[1], now) // parse os.Args[1] i.e. time string to time based on time.Now() format

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Println("Unable to parse time!")
		os.Exit(2)
	}

	if now.After(t.Time) {
		fmt.Println("Set a Future time!") // time given has to be after current time to give reminder in the future
		os.Exit(3)
	}

	diff := t.Time.Sub(now)

	if os.Getenv(markName) == markValue {

		time.Sleep(diff) // we want programme to sleep for the remianing time

		//"Reminder": Title of the notification.
		// strings.Join(os.Args[2:], " "): Joins the command-line arguments starting from index 2 (which is the reminder message) into a single string separated by spaces.
		// "assets/information.png": Path to an image file to display alongside the notification (optional).
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		// If GOLANG_CLI_REMINDER is not "1", it sets up a new process (cmd) to execute the same command (os.Args[0], which is the current program) with the same arguments (os.Args[1:])
		// runs same program with path and then expand all slice items as arguments
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue)) //Sprintf formats according to a format specifier and returns the resulting string
		// sets markname to markvalue i.e. 1 and set this key value pair to env variables
		// It adds GOLANG_CLI_REMINDER=1 to the environment variables for the new process

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed after : ", diff.Round(time.Second))
		os.Exit(0)
	}

}
