package twlog

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func NewTwLogStu(Id string) *TwLogStu {
	var PrnLog TwLogStu
	PrnLog.Init(Id)
	return &PrnLog
}

var TIME_FORMAT string = "15:04:05"

const (
	LOG_DBG = iota
	LOG_INFO
	LOG_ERR
)

type TwLogStu struct {
	Tag      string
	Format   string
	LogLevel int
	LogTag   string
	Depth    int
	DepthMax int
	Tin      time.Time
	Tins     []time.Time
	Tlst     int64
	Tnow     int64
	PwdPre   string
	Tw       *tabwriter.Writer
	Buf      *bytes.Buffer
}

func (r *TwLogStu) Init(Id string) {
	r.Format = "%v\t%v\t%v\t%v\t\n"
	r.Tin = time.Now()
	r.Tag = Id
	r.LogLevel = LOG_DBG
	r.LogTag = "DEBUG"
	r.Buf = new(bytes.Buffer)
	r.Tw = new(tabwriter.Writer).Init(r.Buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(r.Tw, r.Format, "["+time.Now().Format(TIME_FORMAT)+"]", "START", r.Tag, "")
}

func (r *TwLogStu) SetLogLevel(Level int) {
	r.LogLevel = Level
	if r.LogLevel == LOG_DBG {
		r.LogTag = "DEBUG"
	}
	if r.LogLevel == LOG_INFO {
		r.LogTag = "INFO"
	}
	if r.LogLevel == LOG_ERR {
		r.LogTag = "ERROR"
	}
}
func (r *TwLogStu) SetPwdPre(PwdPre string) {
	r.PwdPre = PwdPre
}

func (r *TwLogStu) Debug(v ...interface{}) {
	if r.LogLevel <= LOG_DBG {
		LogLevel := r.LogLevel
		r.SetLogLevel(LOG_DBG)
		r.write(v)
		r.SetLogLevel(LogLevel)
	}
}
func (r *TwLogStu) Info(v ...interface{}) {
	if r.LogLevel <= LOG_INFO {
		LogLevel := r.LogLevel
		r.SetLogLevel(LOG_INFO)
		r.write(v)
		r.SetLogLevel(LogLevel)
	}
}
func (r *TwLogStu) Error(v ...interface{}) {
	if r.LogLevel <= LOG_ERR {
		LogLevel := r.LogLevel
		r.SetLogLevel(LOG_ERR)
		r.write(v)
		r.SetLogLevel(LogLevel)
	}
}

func (r *TwLogStu) write(v ...interface{}) {
	r.Nop()
	fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), r.LogTag, r.GetPlaceTime(3))
	fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), "", "---------DETAIL-----------")
	Str := fmt.Sprint(v)
	row := 0
	for row < len(Str)/100 {
		fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), "", Str[row*100:100*(row+1)])
		row++
	}
	fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), "", Str[row*100:len(Str)])
}
func (r *TwLogStu) Snap() {
	Tin := time.Now()
	if len(r.Tins) <= r.Depth {
		r.Tins = append(r.Tins, Tin)
	} else {
		r.Tins[r.Depth] = Tin
	}
	return
}
func (r *TwLogStu) DepthIncr() {
	r.Depth++
	if r.Depth > r.DepthMax {
		r.DepthMax = r.Depth
	}
}
func (r *TwLogStu) From() *TwLogStu {
	r.Snap()
	r.DepthIncr()
	if r.LogLevel == LOG_DBG {
		fmt.Fprintf(r.Tw, r.Format, "", r.TreeArg(-1), "", "")
		fmt.Fprintf(r.Tw, r.Format, time.Since(r.Tin), r.Tree(), "FROM", r.GetPlaceTime(2))
	}
	return r
}
func (r *TwLogStu) To() *TwLogStu {
	if r.LogLevel == LOG_DBG {
		fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), "TO", r.GetPlaceTime(2))
	}
	return r
}
func (r *TwLogStu) Exit() {
	if r.LogLevel == LOG_DBG {
		r.Nop()
		fmt.Fprintf(r.Tw, r.Format, time.Since(r.Tin), r.TreeEnd(), "EXIT", r.GetPlaceTime(2))
		fmt.Fprintf(r.Tw, r.Format, "", r.TreeArg(-1), time.Since(r.Tins[r.Depth-1]), "")
	}
	r.Depth--

}
func (r *TwLogStu) GetPlaceTime(depth int) (Rst string) {
	funcName, file, line, ok := runtime.Caller(depth)
	if ok {
		file = strings.Replace(file, r.PwdPre, "/", -1)
		Rst = file + "|" + "L" + strconv.Itoa(line) + "|" + runtime.FuncForPC(funcName).Name()
	}
	return
}

func (r *TwLogStu) Tree() string {
	return string(bytes.Repeat([]byte("|"), r.Depth))
}
func (r *TwLogStu) TreeArg(a int) string {
	return string(bytes.Repeat([]byte("|"), r.Depth+a))
}

func (r *TwLogStu) TreeEnd() string {
	Str01 := string(bytes.Repeat([]byte("|"), r.Depth))
	Max := (r.DepthMax / 8) + 1
	Str02 := string(bytes.Repeat([]byte("-"), 8*Max-r.Depth))
	return Str01 + Str02
}
func (r *TwLogStu) Flush() {
	r.Nop()
	fmt.Fprintf(r.Tw, r.Format, "["+time.Now().Format(TIME_FORMAT)+"]", "FINISH", r.Tag, "")
	r.Tw.Flush()
}
func (r *TwLogStu) Nop() {
	fmt.Fprintf(r.Tw, r.Format, "", r.Tree(), "", "")
}
