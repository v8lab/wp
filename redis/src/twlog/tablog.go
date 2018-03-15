package twlog

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"runtime"
	"text/tabwriter"
)

func NewTwLogOptStu() *TwLogOptStu {
	var TwLogOpt TwLogOptStu
	return &TwLogOpt
}

type TwLogOptStu struct {
	Format string
	Time   int64
	Level  int
	Tab    string
	Tag    string
}

func NewTwLogStu(Opt *TwLogOptStu) *TwLogStu {
	var PrnLog TwLogStu
	PrnLog.Init(Opt)
	return &PrnLog
}

type TwLogStu struct {
	Tag    string
	Level  int
	Format string
	Time   int64
	PwdPre string
	Buf    *bytes.Buffer
	Tw     *tabwriter.Writer
}

func (r *TwLogStu) GetPlaceTime() (Rst string) {
	funcName, file, line, ok := runtime.Caller(2)
	if ok {
		file = strings.Replace(file, r.PwdPre, "/", -1)
		Rst = file + "|" + "L" + strconv.Itoa(line) + "|" + runtime.FuncForPC(funcName).Name()

	}
	return
}
func (r *TwLogStu) Init(Opt *TwLogOptStu) {
	r.Format = Opt.Format
	r.Level = Opt.Level
	r.Time = Opt.Time
	r.Tag = Opt.Tag
	r.Buf = new(bytes.Buffer)
	r.Tw = new(tabwriter.Writer).Init(r.Buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(r.Tw, r.Format, "start", r.Tag)
	r.PwdPre = "E:/wp/redis/src/"
}
func (r *TwLogStu) Write(v ...interface{}) {
	fmt.Fprintf(r.Tw, r.Format, "", fmt.Sprint(v))
}
func (r *TwLogStu) Enter() *TwLogStu {

	fmt.Fprintf(r.Tw, r.Format, "enter", r.GetPlaceTime())
	return r
}
func (r *TwLogStu) Exit() *TwLogStu {
	fmt.Fprintf(r.Tw, r.Format, "exit", r.GetPlaceTime())

	return r
}

func (r *TwLogStu) GetPwdFre() (Rst string) {
	_, Rst, _, _ = runtime.Caller(2)

	return
}
func (r *TwLogStu) GetRst() string {
	fmt.Fprintf(r.Tw, r.Format, "finish", r.Tag)
	r.Tw.Flush()
	return r.Buf.String()

}
