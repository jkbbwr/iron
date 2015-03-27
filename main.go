package main

import (
	"flag"
	"github.com/op/go-logging"
	"github.com/jkbbwr/iron/iron"
	"os"
	"runtime/pprof"
)

var step bool
var debug bool
var useLogging bool
var profile bool

var log = logging.MustGetLogger("FeVM")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{id:04d}%{color:reset} %{message}",
)

func init() {
    flag.BoolVar(&debug, "debug", false, "Enable debugging")
    flag.BoolVar(&step, "step", false, "Enable step through")
	flag.BoolVar(&profile, "profile", false, "Enable write profile to file")
	flag.BoolVar(&useLogging, "logging", false, "Enable logging")
}

func main() {
	flag.Parse()

	stdout := logging.NewLogBackend(os.Stderr, "", 0)
	stdoutWithFormatting := logging.NewBackendFormatter(stdout, format)
	logging.SetBackend(stdoutWithFormatting)

	if !useLogging {
		logging.SetLevel(logging.ERROR, "FeVM")
	}

    script := flag.Arg(0)
    if script == "" {
        log.Critical("No program given.")
        os.Exit(-1)
    }

	if profile {
		f, err := os.Create("./github.com.jkbbwr.iron.prof")
		if err != nil {
			log.Critical("%s", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
        // Profiling doesn't allow for any other settings.
        log.Info("Starting and running FeVM 1000 times.")
        for i := 1; i <= 1000; i++ {
            log.Info("Starting a new FeVM")
            vm := iron.NewVM()
            vm.Load(script)
            vm.Run()
        }
        return
	}

	log.Info("Starting FeVM")
	vm := iron.NewVM()

	log.Debug("Running script %s", script)
	vm.Load(script)

	if step {
		log.Debug("Enabling step")
		vm.RunStep()
        return
	}

    if debug {
        log.Debug("Enabling debug")
        vm.RunDebug()
        return
    }

	vm.Run()
}
