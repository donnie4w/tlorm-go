/*
*
* Copyright 2023 tldb Author. All Rights Reserved.
* email: donnie4w@gmail.com
* github.com/donnie4w/tldb
* github.com/donnie4w/tlorm-go
*
 */
package orm

import (
	"testing"

	"github.com/donnie4w/simplelog/logging"
)

type UserAdmin struct {
	Id      int64
	Name    string `idx`
	Age     int
	Level   bool
	Content []byte
	Sex     byte
	Agent   float32 `idx`
	Achie   uint16
	City    int8
}

func TestCreat(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	Create[UserAdmin]()
}

func TestInsert(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	Insert(&UserAdmin{Name: "dong", Age: 23, Level: true, Content: nil, Sex: 2, Agent: 3.2, Achie: 90, City: 50})
}

func TestUpdate(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	err := Update(&UserAdmin{Id: 1, Name: "dong3", Content: []byte("this is new tldb2")})
	logging.Info(err)
}

func TestSelect(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	if ua, err := SelectById[UserAdmin](1); err == nil {
		logging.Debug(ua)
	}
}

func TestSelectIdx(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	if ua, err := SelectByIdx[UserAdmin]("Name", "dong"); err == nil {
		logging.Debug(ua)
	}
}

func TestSelectsByIdLimit(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	if uas, err := SelectsByIdLimit[UserAdmin](1, 10); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}

func TestSelectByIdxLimit(t *testing.T) {
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	if uas, err := SelectByIdxLimit[UserAdmin](0, 10, "Name", "dong", "dong2"); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}

func Benchmark_Select(b *testing.B) {
	b.StopTimer()
	RegisterDefaultResource(false, "192.168.2.108:7100", "mycli=123")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SelectByIdx[UserAdmin]("Name", "dong")
		// SelectByIdxLimit[UserAdmin](0, 10, "Name", "dong", "dong2")
		// SelectById[UserAdmin](1)
		// SelectId[UserAdmin]()
	}

}
