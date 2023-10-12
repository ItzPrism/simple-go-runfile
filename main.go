package main
import (
	"fmt"
	"log"
	"os/exec"
	"os"
)
func runFile(dir string, args string) string {
	cmnd := exec.Command(dir, args)
	cmnd.Start()
	log.Println("log")
	fmt.Println("Logging started")
	return dir
	return args
}
func main() {
	var enterDir string
	var enterArgs string
	var yesNoDir string
	var yesNoArgs string
	fmt.Println("Please enter directory")
	fmt.Println("***")
	fmt.Println("The Directory must be enclosed with double Quotes.")
	fmt.Scanln(&enterDir)
	fmt.Println("Is", enterDir, "Correct? Yes/No")
	fmt.Scanln(&yesNoDir)
	if yesNoDir == "Yes" {
		fmt.Println("Please enter Arguments")
		fmt.Println("***")
		fmt.Println("The Arguments must be enclosed with double Quotes")
		fmt.Scanln(&enterArgs)
		fmt.Println("Are these arguments correct?:", enterArgs, "Yes/No")
		fmt.Scanln(&yesNoArgs)
		if yesNoArgs == "Yes" {
			runFile(enterDir, enterArgs)
			log.Println("Current executed Command: runfile(", enterDir, ", ", enterArgs, ")")
		} else {
			os.Exit(3)
		}

	} else {
		os.Exit(2)
	}

}
