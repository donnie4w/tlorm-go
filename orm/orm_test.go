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
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"testing"

	"github.com/donnie4w/simplelog/logging"
)

func init() {
	go func() {
		if err := http.ListenAndServe(":8001", nil); err != nil {
			panic("debug failed:" + err.Error())
		}
	}()
}

type UserAdmin struct {
	Id      int64
	Name    string `idx`
	Age     int
	Level   bool
	Content []byte
	Gender  byte
	Agent   float32 `idx`
	Achie   *uint16
	DestID  *int
	City    int8
}

func TestCreat(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	Create[UserAdmin]()
}

func TestInsert(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	achie := uint16(90)
	Insert(&UserAdmin{Age: 23, Level: true, Content: nil, Gender: 2, Achie: &achie, City: 49})
}

func TestUpdate(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	achie := uint16(90)
	err := Update(&UserAdmin{Id: 1, Name: "tom2", Age: 23, Level: true, Content: []byte("this is new tldb"), Gender: 2, Agent: 3.2, Achie: &achie, City: 50})
	logging.Info(err)
}

func TestUpdateNonzero(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	err := UpdateNonzero(&UserAdmin{Id: 1, Name: "tom3", Content: []byte("this is new tldb2")})
	logging.Info(err)
}

func TestSelect(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	if ua, err := SelectById[UserAdmin](1); err == nil {
		logging.Debug(ua)
	}
}

func TestSelectIdx(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	if ua, err := SelectByIdx[UserAdmin]("Name", "tom"); err == nil {
		logging.Debug(ua)
	}
}

func TestSelectsByIdLimit(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	if uas, err := SelectsByIdLimit[UserAdmin](1, 10); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}

func TestSelectByIdxLimit(t *testing.T) {
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	if uas, err := SelectByIdxLimit[UserAdmin](0, 10, "Name", "dong", "dong2"); err == nil {
		for _, ua := range uas {
			logging.Debug(ua)
		}
	}
}

func Benchmark_Select(b *testing.B) {
	b.StopTimer()
	RegisterDefaultResource(true, "127.0.0.1:7100", "mycli=123")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SelectByIdx[UserAdmin]("Name", "tom")
		// SelectByIdxLimit[UserAdmin](0, 10, "Name", "dong", "dong2")
		// SelectById[UserAdmin](1)
		// SelectId[UserAdmin]()
	}

}

func TestIntToBytes(t *testing.T) {
	var i int16 = 50
	bs := IntToBytes(i)
	fmt.Println(bs)
	fmt.Println(BytesToInt[int16](bs))
}

type timmucmember struct {
	Id       int64
	RelateId uint64 `idx:"1"`
	Guuid    uint64 `idx:"1"`
	Uuuid    uint64 `idx:"1"`
	Status   uint8
}

type timmessage struct {
	Id     int64
	ChatId uint64 `idx:"1"`
	Stanza []byte
}

func Benchmark_SelectByIdxAscLimit(b *testing.B) {
	RegisterDefaultResource(true, "127.0.0.1:3336", "mycli=123")
	as, err := SelectByIdxAscLimit[timmucmember]("Uuuid", uint64(15534333452477200259), 2, 1)
	fmt.Println(err)
	for _, a := range as {
		fmt.Println(a)
	}
}

func Benchmark_SelectByIdxDescLimit(b *testing.B) {
	RegisterDefaultResource(true, "127.0.0.1:3336", "mycli=123")
	as, err := SelectByIdxDescLimit[timmucmember]("Uuuid", uint64(15534333452477200259), 9, 4)
	fmt.Println(err)
	for _, a := range as {
		fmt.Println(a)
	}
}

func Benchmark_DeleteBatch(b *testing.B) {
	RegisterDefaultResource(true, "127.0.0.1:3336", "mycli=123")
	err := DeleteBatch[timmessage](3404, 3405, 3406)
	fmt.Println(err)
}

func Benchmark_SelectIdByIdxSeq(b *testing.B) {
	RegisterDefaultResource(true, "127.0.0.1:3336", "mycli=123")
	id, err := SelectIdByIdxSeq[timmucmember]("Uuuid", uint64(15534333452477200259),49)
	fmt.Println(id, err)
}
