package main

import (
	"fmt"
	"syscall"
)

func memInstalled() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	return in.Totalram * uint64(in.Unit)
}

func uptime() string {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return ""
	}

	// Convert the seconds returned by in.Uptime to hours, minutes, and seconds
	h := in.Uptime / 3600 % 24
	m := in.Uptime / 60 % 60
	s := in.Uptime % 60

	uptimeSeconds := fmt.Sprintf("%vh %vm %vs", h, m, s)
	return uptimeSeconds
}
