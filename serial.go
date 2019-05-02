package main

import (
    "github.com/tarm/serial"
    "log"
    "fmt"
    "os/exec"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "time"
)

type properties struct {
    First      string `yaml:"first"`
    Second     string `yaml:"second"`
    Duration   int    `yaml:"duration"`
    Name       string `yaml:"name"`
}

type mode struct {
    Mode []properties `yaml:"mode"`
}


func (c *mode) getConf() *mode {
    yamlFile, err := ioutil.ReadFile("./serial.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c

}

func exec_shell(cmd string) string {
    out, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        log.Fatal("EXEC_SHELL:", err)
    }
    //fmt.Println(out[:len(out)-1])
    return string(out[:len(out)-1])
}

func send_to_tty(first string, second string) {
    con := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
    ser, err := serial.OpenPort(con)
    if err != nil {
       log.Fatal("SEND TO TTY:",err)
    }
    line1 := exec_shell(first)
    fmt.Println(line1)
    line2 := exec_shell(second)
    fmt.Println(line2)
    ser.Write([]byte(line1))
    ser.Write([]byte("\x01"))
    ser.Write([]byte(line2))
}

func main() {
    var c mode
    c.getConf()

    for {
        for i := range c.Mode {
            send_to_tty(c.Mode[i].First, c.Mode[i].Second)
            //fmt.Println(c.Mode[i])
            time.Sleep(time.Duration(c.Mode[i].Duration) * time.Second)
        }

    }
}