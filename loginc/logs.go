package loginc

import (
	"fmt"
	"logs_go/conf"
	"time"
)

type log_chan struct {
	type_nume  int
	type_value string
	times      string
	time_int   int64
}

var log_channle chan *log_chan

func DebugInfo(prompt string, value ...interface{}) {
	values := fmt.Sprintf(prompt, value)
	time_format := time.Now().Format("2006-01-02 15:04:05")
	time_unix := time.Now().Unix()
	fmt.Println("Debug Info ", values, " time_format ", time_format, " time_uinx ", time_unix)
	log_channle = make(chan *log_chan, 1)
	log_channle <- &log_chan{
		type_nume:  Log_DeBug,
		type_value: values,
		times:      time_format,
		time_int:   time_unix,
	}
}

func Infos(prompt string, value ...interface{}) {
	values := fmt.Sprintf(prompt, value)
	time_format := time.Now().Format("2006-01-02 15:04:05")
	time_unix := time.Now().Unix()
	fmt.Println("Infos ", values, " time_format ", time_format, " time_uinx ", time_unix)
	log_channle = make(chan *log_chan, 1)
	log_channle <- &log_chan{
		type_nume:  Log_Info,
		type_value: values,
		times:      time_format,
		time_int:   time_unix,
	}
}

func Dangerous(prompt string, value ...interface{}) {
	values := fmt.Sprintf(prompt, value)
	time_format := time.Now().Format("2006-01-02 15:04:05")
	time_unix := time.Now().Unix()
	fmt.Println("Dangerous ", values, " time_format ", time_format, " time_uinx ", time_unix)
	log_channle = make(chan *log_chan, 1)
	log_channle <- &log_chan{
		type_nume:  Log_Dangerous,
		type_value: values,
		times:      time_format,
		time_int:   time_unix,
	}
}

func Errors(prompt string, value ...interface{}) {
	values := fmt.Sprintf(prompt, value)
	time_format := time.Now().Format("2006-01-02 15:04:05")
	time_unix := time.Now().Unix()
	fmt.Println("Errors ", values, " time_format ", time_format, " time_uinx ", time_unix)
	log_channle = make(chan *log_chan, 1)
	log_channle <- &log_chan{
		type_nume:  Log_Error,
		type_value: values,
		times:      time_format,
		time_int:   time_unix,
	}
}

func Crashs(prompt string, value ...interface{}) {
	values := fmt.Sprintf(prompt, value)
	time_format := time.Now().Format("2006-01-02 15:04:05")
	time_unix := time.Now().Unix()
	fmt.Println("Crashs ", values, " time_format ", time_format, " time_uinx ", time_unix)
	log_channle = make(chan *log_chan, 1)
	log_channle <- &log_chan{
		type_nume:  Log_Crash,
		type_value: values,
		times:      time_format,
		time_int:   time_unix,
	}
}

func Withlog() {
	for {
		select {
		case info := <-log_channle:
			close(log_channle)
			switch info.type_nume {
			case Log_DeBug:
				if Log_DeBug <= conf.GetIntoFileLevel() {

				}
			case Log_Info:
				if Log_Info >= conf.GetIntoFileLevel() {

				}
			case Log_Dangerous:
				if Log_Dangerous >= conf.GetIntoFileLevel() {

				}
			case Log_Error:
				if Log_Error >= conf.GetIntoFileLevel() {

				}
			case Log_Crash:
				if Log_Crash >= conf.GetIntoFileLevel() {

				}
			default:
				fmt.Println("this log level error : level " , info.type_nume)
			}
		}
	}
}
