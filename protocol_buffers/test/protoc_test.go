package test

import (
	"fmt"
	proto_test "github.com/ByronLeeLee/go/study/protocol_buffers"
	 "google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"testing"
)

func write() {
	p1 := &proto_test.Person{
		Id:   1,
		Name: "小李",
		Phones: []*proto_test.Phone{
			{Type: proto_test.PhoneType_HOME, Number: "123456"},
			{Type: proto_test.PhoneType_WORK, Number: "345678"},
		},
	}

	p2 := &proto_test.Person{
		Id:   2,
		Name: "小刘",
		Phones: []*proto_test.Phone{
			{Type: proto_test.PhoneType_WORK, Number: "567890"},
			{Type: proto_test.PhoneType_HOME, Number: "678901"},
		},
	}

	book := &proto_test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	data, _ := proto.Marshal(book)

	ioutil.WriteFile("./book.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("./book.txt")
	book := &proto_test.ContactBook{}

	proto.Unmarshal(data, book)
	for _, person := range book.Persons {
		for _, number := range person.Phones {
			fmt.Printf("id:%d;name:%s;phoneNumber:%s", person.Id, person.Name, number.Number)
			fmt.Println()
		}
	}
}

func TestProtoFile(t *testing.T) {
	write()
	read()
}
