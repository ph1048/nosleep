package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	k32 := windows.NewLazyDLL("kernel32")
	pfnSetThreadExecutionState := k32.NewProc("SetThreadExecutionState")
	const ES_CONTINUOUS = 0x80000000
	const ES_SYSTEM_REQUIRED = 0x00000001
	pfnSetThreadExecutionState.Call(ES_CONTINUOUS | ES_SYSTEM_REQUIRED)
	println("System is prevented from sleep")
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	println("Exiting")
	pfnSetThreadExecutionState.Call(ES_CONTINUOUS)
}
