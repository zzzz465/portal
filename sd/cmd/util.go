package cmd

import "os"

func errExit(exitCode int, format string, err error) {
    log.Errorf(format, err)
    os.Exit(exitCode)
}
