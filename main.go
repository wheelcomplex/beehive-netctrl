package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	bh "github.com/kandoo/beehive"
	"github.com/kandoo/beehive-netctrl/controller"
	"github.com/kandoo/beehive-netctrl/nom"
	"github.com/kandoo/beehive-netctrl/openflow"
	"github.com/kandoo/beehive-netctrl/switching"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	h := bh.NewHive()
	openflow.StartOpenFlow(h)
	controller.RegisterNOMController(h)
	discovery.RegisterDiscovery(h)

	//app := h.NewApp("Hub")
	//app.Handle(nom.PacketIn{}, &switching.Hub{})
	//app.SetFlags(0)

	h.Start()
}
