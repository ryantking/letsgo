package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must specify a service to run")
		os.Exit(1)
	}
	user, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get current user: %s\n", err.Error())
		os.Exit(1)
	}
	cfg, err := LoadConfig(user.HomeDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load config: %s\n", err.Error())
		os.Exit(1)
	}
	svc := os.Args[1]
	svcCfg, ok := cfg[svc]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown service: %s\n", svc)
		os.Exit(1)
	}
	dir := os.ExpandEnv(svcCfg.Dir)
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "could not cd to dir '%s': %s\n", dir, svc)
		os.Exit(1)
	}
	cmd := exec.Command(svcCfg.Command, svcCfg.Args...)
	cmd.Env = append(os.Environ(), svcCfg.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %s\n", err.Error())
		os.Exit(1)
	}
}
