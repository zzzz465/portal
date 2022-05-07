package cmd

import "os"

func errExit(exitCode int, template string, err error) {
	log.Errorf(template, err)
	os.Exit(exitCode)
}
