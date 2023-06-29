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
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	Create[UserAdmin]()
}

func TestInsert(t *testing.T) {
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	Insert(&UserAdmin{Name: "dong", Age: 23, Level: true, Content: []byte("this is new tldb"), Sex: 2, Agent: 3.2, Achie: 90, City: 50})
}

func TestUpdate(t *testing.T) {
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	Update(&UserAdmin{Id: 1, Name: "dong", Age: 33, Level: false, Content: []byte("this is new tldb"), Sex: 1, Agent: 3.1, Achie: 90, City: 50})
}

func TestSelect(t *testing.T) {
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	if ua, err := SelectById[UserAdmin](1); err == nil {
		logging.Debug(ua)
	}
}

func TestSelectsByIdLimit(t *testing.T) {
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	if uas, err := SelectsByIdLimit[UserAdmin](1, 10); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}

func TestSelectByIdxLimit(t *testing.T) {
	RegisterDefaultResource(true, "192.168.2.108:7000", "mycli=123")
	if uas, err := SelectByIdxLimit[UserAdmin](0, 10, "Name", "dong"); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}
