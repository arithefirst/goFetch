package main

import (
	"fmt"
	"github.com/zcalusic/sysinfo"
)

// Function for generating bolded text
func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func main() {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	// Get the logo and title for the OS, and save the OS color esc
	esc := osReply(si)

	// CPU
	fmt.Printf("   %v %v\n", bold(fmt.Sprintf("%vCPU:", esc)), si.CPU.Model)

	// Drives/Storage
	fmt.Println(bold(fmt.Sprintf("%v   Drives:", esc)))
	for _, drive := range si.Storage {
		fmt.Printf("    ╰─%v %v (%vGB)", bold(drive.Name+":"), drive.Model, drive.Size)
		fmt.Println()
	}

	// Kernel
	fmt.Printf("   %v %v\n", bold(fmt.Sprintf("%vKernel:", esc)), si.Kernel.Release)

	// Memory
	fmt.Printf("   %v %vMiB\n", bold(fmt.Sprintf("%vMemory:", esc)), memInstalled()/1024/1024)

	//	Uptime
	fmt.Printf("   %v %v\n", bold(fmt.Sprintf("%vUptime:", esc)), uptime())

}
