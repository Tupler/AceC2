package Logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"time"
)

const (
	IsDebug = true
)

type Logger struct {
	lg        *log.Logger
	logfile   string
	isLogFile bool
}

var (
	BlueText   = color.New(color.FgBlue).SprintfFunc()
	RedText    = color.New(color.FgHiRed).SprintfFunc()
	GreenText  = color.New(color.FgGreen).SprintfFunc()
	YellowText = color.New(color.FgHiYellow).SprintfFunc()
)

func NewLogger(islogfile bool) *Logger {

	return &Logger{
		lg:        log.New(os.Stdout, "", 0),
		logfile:   "./log.txt",
		isLogFile: false,
	}
}
func (l *Logger) Error(data ...string) {

	prefix := fmt.Sprintf("[%s]:", RedText("ERROR"))
	l.lg.SetPrefix(prefix)
	l.lg.Println(data)

}

func (l *Logger) Info(data ...string) {
	prefix := fmt.Sprintf("[%s]:", BlueText("INFO"))
	l.lg.SetPrefix(prefix)
	l.lg.Println(data)
}

func (l *Logger) Warn(data ...string) {
	prefix := fmt.Sprintf("[%s]:", YellowText("WARN"))
	l.lg.SetPrefix(prefix)
	l.lg.Println(data)
}

func (l *Logger) Debug(data ...string) {
	prefix := fmt.Sprintf("[%s]:", "DEBUG")
	l.lg.SetPrefix(prefix)
	l.lg.Println(data)
	
}

func (l *Logger) Success(data ...string) {
	if IsDebug {
		prefix := fmt.Sprintf("[%s]:", GreenText("SUCCESS"))
		l.lg.SetPrefix(prefix)
		l.lg.Println(data)
	}
}

func (l *Logger) LogToFile(data ...string) {
	fileHandle, err := os.OpenFile(time.Now().Format("06年01月02日")+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log fail ,err: %v\n", err)
		return
	}
	_, err = fmt.Fprintf(fileHandle, "[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), data)
	if err != nil {
		return
	}
}

var ALogger *Logger
