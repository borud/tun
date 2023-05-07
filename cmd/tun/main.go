// Package main contains the executable main.
package main

import (
	"fmt"
	"log"

	"github.com/borud/tun/pkg/global"
	"github.com/borud/tun/pkg/tun"
	"github.com/borud/tun/pkg/util"
)

var opt struct {
	KeyFile          string   `long:"key" description:"SSH key file"`
	Target           string   `long:"target" default:"localhost:22" description:"target SSH server for login on local end" required:"yes"`
	RemoteListenAddr string   `long:"remote-listen-addr" default:"localhost:2222" description:"remote listener address" required:"yes"`
	Via              []string `long:"via" description:"hosts we jump via on format user@host:port"`

	Version versionCmd `command:"version" description:"show version"`
}

type versionCmd struct{}

func (v *versionCmd) Execute([]string) error {
	return fmt.Errorf("%s", global.Version)
}

func main() {
	util.FlagParse(&opt)

	t, err := tun.New(tun.Config{
		KeyFile:          opt.KeyFile,
		Target:           opt.Target,
		RemoteListenAddr: opt.RemoteListenAddr,
		Chain:            opt.Via,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = t.Run()
	if err != nil {
		log.Fatal(err)
	}
}
