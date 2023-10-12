package main
import (
  "fmt", "log", "os/exec"  
)
func runFile(dir string, args string) string {
	cmnd := exec.Command(dir, args)
	cmnd.Start()
	log.Println("log")
	fmt.Println("Logging started")
	return dir
	return args
