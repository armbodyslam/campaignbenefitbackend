package common

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//Applog strucure
type Applog struct {
	Timestamp       string   `json:"@timestamp"`
	Tags            []string `json:"tags"`
	OrderNo         string   `json:"orderno"`
	TrackingNo      string   `json:"trackingno"`
	ApplicationName string   `json:"applicationname"`
	FunctionName    string   `json:"functionname"`
	OrderDate       string   `json:"orderdate"`
	TVSNo           string   `json:"tvsno"`
	MobileNo        string   `json:"mobileno"`
	SerialNo        string   `json:"serialno"`
	Reference1      string   `json:"reference1"`
	Reference2      string   `json:"reference2"`
	Reference3      string   `json:"reference3"`
	Reference4      string   `json:"reference4"`
	Reference5      string   `json:"reference5"`
	Request         string   `json:"request"`
	Response        string   `json:"response"`
	Start           string   `json:"start"`
	End             string   `json:"end"`
	Duration        string   `json:"duration"`
}

// NewApplog Obj
func NewApplog() *Applog {
	t0 := time.Now()
	return &Applog{
		Timestamp: t0.Format(time.RFC3339Nano),
	}
}

// NewApploginfo Obj
func NewApploginfo(trackingno string, applicationname string,
	functionname string, tagappname string,
	taglogtype string) Applog {
	var applog Applog
	var tags []string
	env := os.Getenv("ENVAPP")
	tags = append(tags, env)
	tags = append(tags, tagappname)
	tags = append(tags, taglogtype)
	applog.TrackingNo = trackingno
	applog.ApplicationName = applicationname
	applog.FunctionName = functionname
	applog.Tags = tags
	applog.Timestamp = time.Now().Format(time.RFC3339Nano)
	return applog
}

//Processconfig struct
type Processconfig struct {
	CallFunction string `json:"callfunction"`
	Start        string `json:"start"`
	End          string `json:"end"`
	Duration     string `json:"duration"`
	ResultCode   string `json:"resultcode"`
	ResultDesc   string `json:"resultdesc"`
}

//Workflowlog strucure
type Workflowlog struct {
	OrderNo         string          `json:"orderno"`
	TrackingNo      string          `json:"trackingno"`
	ApplicationName string          `json:"applicationname"`
	FunctionName    string          `json:"functionname"`
	OrderDate       string          `json:"orderdate"`
	TVSNo           string          `json:"tvsno"`
	MobileNo        string          `json:"mobileno"`
	SerialNo        string          `json:"serialno"`
	Reference1      string          `json:"reference1"`
	Reference2      string          `json:"reference2"`
	Reference3      string          `json:"reference3"`
	Reference4      string          `json:"reference4"`
	Reference5      string          `json:"reference5"`
	InputData       string          `json:"InputData"`
	Start           string          `json:"start"`
	End             string          `json:"end"`
	Duration        string          `json:"duration"`
	ProcessConfig   []Processconfig `json:"processconfig"`
}

//PrintJSONLog func
func (a *Applog) PrintJSONLog() error {
	logJSON, _ := json.Marshal(a)
	fmt.Println(string(logJSON))
	return nil
}
