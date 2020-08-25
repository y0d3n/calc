package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func slot(s string) {
	var str = "\t"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		for j := 0; j < rand.Intn(150); j++ {
			fmt.Printf("\r%s%s", str, string(abc[rand.Intn(len(abc)-1)]))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
		str += string(s[i])
	}
	fmt.Printf("\r%s\n", str)
}

func main() {
	fmt.Printf("\x1b[32m\n\n")

	fmt.Printf("\t")
	slot("Initializing...")
	fmt.Printf("\n\n")
	slot("837E834E81755982B382F182CC82B182C682CD")
	slot("82E082A4965982EA82BD82D982A482AA")
	slot("82A282A282C68E7682A282DC82B78176...")
	fmt.Printf("\n")
	slot("838B834A8175918A95CF82ED82E782B8")
	slot("96A297FB82AA82DC82B582A282C88176...")
	fmt.Printf("\n")
	slot("8357837E817592FA82DF82AA88AB82A282E68176...")

	fmt.Printf("\n\n")

	time.Sleep(time.Second)
	fmt.Println("\t　■■■■　　　　　　　■")
	time.Sleep(time.Second)
	fmt.Println("\t■　　　　■　　■■　　■　　■■■")
	time.Sleep(time.Second)
	fmt.Println("\t■　　　　　　　　　■　■　■")
	time.Sleep(time.Second)
	fmt.Println("\t■　　　　　　　■■■　■　■")
	time.Sleep(time.Second)
	fmt.Println("\t■　　　　■　■　　■　■　■")
	time.Sleep(time.Second)
	fmt.Println("\t　■■■■　　　■■■　■　　■■■　■")
	time.Sleep(time.Second)

	fmt.Printf("\n\n")

	slot("8357837E815B817581638163")
	slot("28814C8145")
	slot("81CD8145814D29")
	slot("816381638176...")
	fmt.Printf("\x1b[0m")
	slot("hogefugapiyo")
}
