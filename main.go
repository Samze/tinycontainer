package main

import (
	"os/exec"
	"os"
	"syscall"
)

func main(){
	if len(os.Args) < 2 {
		panic("gimme something to do")
	}

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
	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))
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
	cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
