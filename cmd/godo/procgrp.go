// +build darwin dragonfly freebsd linux netbsd openbsd

package main

import (
	"os/exec"
	"syscall"
)

// killGodo kills the spawned godo process.
func killGodo(cmd *exec.Cmd, killProcessGroup bool) {
	cmd.Process.Kill()
	// process group may not be cross platform but on Darwin and Linux, this
	// is the only way to kill child processes
	if killProcessGroup {
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		if err != nil {
			panic(err)
		}
		syscall.Kill(-pgid, syscall.SIGKILL)
	}
}
