package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/global"
	"log"
	"time"
)

type BaseModel struct {
	ID        string        `json:"id" gorm:"primary_key"`
	CreatedAt FormatterTime `json:"createdAt"`
	UpdatedAt FormatterTime `json:"updatedAt"`
}

type FormatterTime time.Time

func (t *FormatterTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	formattedTime := tTime.Format("2006-01-02 15:04:05") // 按需格式化
	return json.Marshal(formattedTime)
}
func (t *FormatterTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	err := json.Unmarshal(data, &timeStr)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("解析时间时出错:%v", err))
		return err
	}
	tTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Println("解析时间时出错", err)
		return err
	}
	*t = FormatterTime(tTime)
	return nil
}
func (t FormatterTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *FormatterTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = FormatterTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
