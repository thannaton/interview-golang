package logUtils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	Info    *LogUtils
	Warning *LogUtils
	Debug   *LogUtils
	Error   *LogUtils
)

type LogUtils struct {
	enable bool
	logger *log.Logger
}

var filename = "logs-{date}.log"

func InitLogs(enableInfo bool, enableWarning bool, enableDebug bool, enableError bool) {
	dirPath := filepath.Join("log-files")
	if _, staterr := os.Stat(dirPath); os.IsNotExist(staterr) {
		os.Mkdir(dirPath, 0755)
	}

	logFilename := strings.ReplaceAll(filename, "{date}", time.Now().Format("2006-01-02"))

	writeFile, err := os.OpenFile(filepath.Join(dirPath, logFilename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("%v - %v %v ", color.RedString("ERROR"), "Cannot open log file: ", err)
	}

	logWriter := io.MultiWriter(os.Stdout, writeFile)
	Info = NewLogs(enableInfo, color.CyanString("[INFO] "), logWriter)
	Warning = NewLogs(enableWarning, color.YellowString("[WARNING] "), logWriter)
	Debug = NewLogs(enableDebug, color.MagentaString("[DEBUG] "), logWriter)
	Error = NewLogs(enableError, color.RedString("[ERROR] "), logWriter)
}

func NewLogs(enable bool, prefix string, writer io.Writer) *LogUtils {
	return &LogUtils{
		enable: enable,
		logger: log.New(writer, prefix, log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

func (l *LogUtils) Println(v ...interface{}) {
	if l.enable || l.logger != nil {
		l.logger.Println(v...)
	}
}

func (l *LogUtils) Printf(format string, v ...interface{}) {
	if l.enable || l.logger != nil {
		l.logger.Printf(format, v...)
	}
}

func (l *LogUtils) Fatal(v ...interface{}) {
	if l.enable || l.logger != nil {
		l.logger.Fatal(v...)
	}
}
