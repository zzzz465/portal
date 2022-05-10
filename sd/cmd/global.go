package cmd

import "go.uber.org/zap"

var log *zap.SugaredLogger

func init() {
    l, _ := zap.NewDevelopment()
    log = l.Sugar()
}
