package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// we can rest noon for 75 minutes, and work for 8 hours
const workTime = 8 * 60 + 75

func getCMDResult(today string) (string, error) {
    // wevtutil qe system /q:Event[EventData[Data[@Name='StartTime']>'2017-09-14T00:00:00']]
    subcmd := fmt.Sprintf("wevtutil qe system /q:Event[EventData[Data[@Name='StartTime']^>'%s']]", today)
    cmd := exec.Command("cmd", "/C", subcmd)

    out, err := cmd.Output()
    if err != nil {
        return "", err
    }

    return string(out), nil
}

func getStartTime(log string) string {
    pattern := "'StartTime'>"

    i := strings.Index(log, pattern)
    if i >= 0 {
        start := i + len(pattern)
        out := log[start:start+30]
        return out
    }

    return ""
}

func getHumanTime(t time.Time) string {
    return fmt.Sprintf(
        "%04d-%02d-%02d %02d:%02d",
        t.Year(),
        t.Month(),
        t.Day(),
        t.Hour(),
        t.Minute(),
    )
}

func printWorkTime(startTime string) {
    t1, err := time.Parse(time.RFC3339, startTime)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("")
    fmt.Printf("clock in  : %s \n", getHumanTime(t1.Local()))
    fmt.Printf("clock out : %s \n", getHumanTime(t1.Add(time.Duration(workTime)*time.Minute).Local()))
}

func main() {
    t1 := time.Now()
    t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.FixedZone("UTC", 0))

    out, err := getCMDResult(t2.Format(time.RFC3339))
    if err != nil {
        log.Fatal(err)
    }

    startTime := getStartTime(out)

    if startTime != "" {
        printWorkTime(startTime)
        return
    }

    fmt.Println("Can't get start time in Windows event log")
}
