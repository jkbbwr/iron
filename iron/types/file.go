package types

import (
	"fmt"
	"github.com/op/go-logging"
	"net"
	"os"
)

const (
	FILE      = "file"
	SOCKET    = "socket"
	STREAM    = "stream"
	LISTEN    = "listen"
	READ      = 0
	WRITE     = 1
	READWRITE = 2
	APPEND    = 3
	TCP       = 4
	UDP       = 5
)

var log = logging.MustGetLogger("FeVM")

type FeFile struct {
	Path     string
	Mode     string
	FileType string
	GoFile   *os.File
	GoConn   net.Conn
}

func NewFile(ftype, path, mode string) FeFile {
	file := FeFile{
		Path: path,
		Mode: mode,
	}
	switch ftype {
	case FILE:
		var fi *os.File
		var err error
		switch mode {
		case "r":
			fi, err = os.OpenFile(path, os.O_RDONLY, 0666)
		case "w":
			fi, err = os.OpenFile(path, os.O_RDWR, 0666)
		case "a":
			fi, err = os.OpenFile(path, os.O_APPEND, 0666)
		}
		if err != nil {
			log.Fatal(err)
		}
		file.GoFile = fi
		file.FileType = FILE
	case SOCKET:
		conn, err := net.Dial(mode, path)
		if err != nil {
			log.Fatal(err)
		}
		file.GoConn = conn
		file.FileType = SOCKET
	case STREAM:
		log.Fatal("Stream not implemented yet")
	default:
		log.Fatal("Unknown type. Expected one of file, socket, stream got %s", ftype)
	}
	return file
}

func (this *FeFile) ReadFile(size int) string {
	buf := make([]byte, size)
	_, err := this.GoFile.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func (this *FeFile) ReadSocket(size int) string {
	buf := make([]byte, size)
	_, err := this.GoConn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func (this *FeFile) Read(size int) string {
	switch this.FileType {
	case FILE:
		return this.ReadFile(size)
	case SOCKET:
		return this.ReadSocket(size)
	default:
		log.Fatal("Unknown file type")
	}
	return ""
}

func (this FeFile) String() string {
	return fmt.Sprintf("FeFile{%s with path %s in mode %s}", this.FileType, this.Path, this.Mode)
}

func (this FeFile) Cmp(that FeType) int {
	panic("FeFile: Not implemented")
}
