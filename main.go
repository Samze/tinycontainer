package main

import (
	"os/exec"
	"os"
)

func main(){
	switch os.Args[1] {
	case "run" :
		parent(os.Args[2:]...)
	case "child":
		child(os.Args[2:]...)
	default:
		panic(os.Args[1])
	}
}

func child(args ...string){
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func parent(args ...string) {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

