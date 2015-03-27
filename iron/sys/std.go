package sys

import (
	"github.com/op/go-logging"
    "github.com/jkbbwr/iron/iron/types"
)

var log = logging.MustGetLogger("FeVM")

// A sys call always gets passed the current window.
// It then mutilates the input and output sections.
type SysMap map[string]func([]types.FeType)

type System struct {
	Table SysMap
}

func NewSystemMapper() System {
	s := &System{make(SysMap)}
	s.Add("print", SysPrint)
	log.Debug("Registered `print` as a syscall")
	return *s
}

func (s *System) Add(name string, function func([]types.FeType)) {
	s.Table[name] = function
}

func (s System) Invoke(name string, window []types.FeType) {
	s.Table[name](window)
}

