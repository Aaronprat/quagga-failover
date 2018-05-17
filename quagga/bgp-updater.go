package main

import (
        "io/ioutil"
        "log"
        "strings"
	"os/exec"
	"flag"
	"fmt"
)

func main() {
	var state = flag.String("state", "unknown", "Required State - up/down")
	flag.Parse()
	
	bool := 0	

        input, err := ioutil.ReadFile("/etc/quagga/bgpd.conf")
        if err != nil {
                log.Fatalln(err)
        }

        lines := strings.Split(string(input), "\n")
        
	if *state == "up" {
		for i, line := range lines {
			if strings.Contains(line, "#ip prefix-list anycast-ip seq 10") {
				lines[i] = strings.Replace(line, "#", " ", 1)
				bool = 1
			} else if strings.Contains(line, " ip prefix-list anycast-ip seq 10") {
				fmt.Println("bgp already up and advertising")
			}
		}
	} else if *state == "down" {
		for i, line := range lines {	
			if strings.Contains(line, " ip prefix-list anycast-ip seq 10") {
				lines[i] = strings.Replace(line, " ", "#", 1)
				bool = 1
			} else if strings.Contains(line, "#ip prefix-list anycast-ip seq 10") {
				fmt.Println("bgp already down and not advertisin")
			}
		}
	} else {
		fmt.Println("wrong state presented:", *state)
		}

        output := strings.Join(lines, "\n")
        err = ioutil.WriteFile("/etc/quagga/bgpd.conf", []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
        }
	if bool == 1 {
		 bgpRestart()
	}
}

func bgpRestart() {
	exec.Command("service", "bgpd", "restart").CombinedOutput()
}
