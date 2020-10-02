package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/emersion/go-ical"
)

func main() {
	exitCode := 0
	defer func() {
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage:\n%s <filename>\n", os.Args[0])
		exitCode = 1
		return
	}

	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file: %v\n", err)
		exitCode = 1
		return
	}
	defer fd.Close()

	calendar, err := ical.NewDecoder(fd).Decode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse file: %v\n", err)
		exitCode = 1
		return
	}

	for _, ev := range calendar.Events() {
		PrintIndentedString(ev.Props, ical.PropSummary, "    Summary:")

		if startdate, err := ev.DateTimeStart(time.Local); err == nil {
			fmt.Println("      Start:", startdate)

		}
		if enddate, err := ev.DateTimeEnd(time.Local); err == nil {
			fmt.Println("        End:", enddate)
		}

		PrintPersonInfo(ev.Props, ical.PropOrganizer, "  Organizer:")
		PrintPersonInfo(ev.Props, ical.PropAttendee, "   Attendee:")
		PrintIndentedString(ev.Props, ical.PropLocation, "   Location:")
		PrintIndentedString(ev.Props, ical.PropDescription, "Description:")
	}
}

func PrintPersonInfo(props ical.Props, propName string, displayName string) {
	x := props[propName]
	for _, prop := range x {
		val := prop.Value
		if cn, ok := prop.Params[ical.ParamCommonName]; ok && len(cn) > 0 {
			val = cn[0] + " " + val
		}

		if role, ok := prop.Params[ical.ParamRole]; ok && len(role) > 0 {
			for _, r := range role {
				if r == "REQ-PARTICIPANT" {
					val += " (required)"
				}
			}
		}
		fmt.Println(displayName, val)
	}
}

func PrintIndentedString(props ical.Props, propName string, displayName string) {
	if desc, err := props.Text(propName); err == nil {
		replacer := strings.NewReplacer("\\n", "\n", "\\r", "")
		desc := strings.Split(replacer.Replace(desc), "\n")
		prefix := displayName
		for _, l := range desc {
			if l == "" {
				continue
			}
			fmt.Println(prefix, l)
			prefix = strings.Repeat(" ", len(displayName))
		}
	}
}
