package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	//change <os.Getenv("ChromeDriver")> to your driverpath
	driver := agouti.ChromeDriver(os.Getenv("ChromeDriver"), agouti.ChromeOptions("args", []string{"--mute-audio", "--incognito", "start-maximized", "--disable-gpu", "--disable-extensions"}))

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	if err := page.Navigate("https://prayut.click/"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	sectionTitle, err := page.Title()
	if len(sectionTitle) > 0 {
		trash := page.FindByID("svelte")
		var (
			count int  = 0
			loop  bool = true
		)
		ex := exec.Command("cmd", "/c", "cls")
		ex.Stdout = os.Stdout
		ex.Run()
		start := time.Now()
		fmt.Println(page.Title())
		for loop {
			trash.Click()
			count += 1

			fmt.Printf("\rClicked!: %d", count)
			if count == 2000 {
				elapsed := time.Since(start)
				count = 0
				for i := 1; i <= 30; i++ {
					fmt.Printf("\rElapsed: %s, Cool Down!, %02d", elapsed, i)
					time.Sleep(time.Second * 1)
				}
				ex := exec.Command("cmd", "/c", "cls")
				ex.Stdout = os.Stdout
				ex.Run()
				start = time.Now()

			}
		}
	}
	ErrHandler(err)
	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
}
