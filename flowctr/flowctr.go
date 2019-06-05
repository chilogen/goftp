package flowctr

import (
	"fmt"
	"github.com/chilogen/goftp/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	updateTime = time.Second * 5
	UPLOAD     = 0
	DWLOAD     = 1
)

type FlowCount struct {
	count sync.Map
	maxcnt   int
	cnt int
}

type FlowCountModel struct {
	UpLoadCount   int64
	DownLoadCount int64
	StartTime     time.Time
}

var (
	flowCount *FlowCount
)

func init()  {
	flowCount=&FlowCount{}
	flowCount.cnt=0
	flowCount.maxcnt=config.GetConfig().FlowCount.MaxCon
}

func (flowCount *FlowCount) Show() {
	for {
		time.Sleep(updateTime)
		var b strings.Builder
		_, _ = fmt.Fprintf(&b, "========================---> %v\n", time.Now())
		_, _ = fmt.Fprint(&b, "User      total-upload     total-download\n")
		flowCount.count.Range(func(key, value interface{}) bool {
			counter := value.(FlowCountModel)
			_, _ = fmt.Fprintf(&b, fmt.Sprintf("%-10v%10v%10v\n", key,
				hunmanReadAble(counter.UpLoadCount),
				hunmanReadAble(counter.DownLoadCount)))
			return true
		})
		fmt.Println(b.String())
	}
}

func hunmanReadAble(bytes int64)string{
	suffix:=[]string{"B","KB","M","G","T"}
	cnt:=0
	for{
		if bytes<1000{
			break
		}
		bytes=bytes/1000
		cnt++
	}
	return strconv.FormatInt(bytes,10)+suffix[cnt]
}

func (flowCount *FlowCount) Regist(userName string)(success bool) {
	newModel := FlowCountModel{
		UpLoadCount:   0,
		DownLoadCount: 0,
		StartTime:     time.Now(),
	}
	flowCount.count.Store(userName, newModel)
	if flowCount.cnt==flowCount.maxcnt{
		return false
	}
	flowCount.cnt++
	return true
}

func (flowCount *FlowCount) Unregist(UserName string) {
	flowCount.count.Delete(UserName)
	flowCount.cnt--
}

func (flowCount *FlowCount) Add(UserName string, t int, cnt int64) {
	x, _ := flowCount.count.Load(UserName)
	newModel := x.(FlowCountModel)
	if t == UPLOAD {
		newModel.UpLoadCount = newModel.UpLoadCount + cnt
	} else if t == DWLOAD {
		newModel.DownLoadCount = newModel.DownLoadCount + cnt
	}
	flowCount.count.Store(UserName, newModel)
}

func GetFlowCount() (ret *FlowCount) {
	return flowCount
}

