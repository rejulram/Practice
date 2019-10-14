package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/ProtocolBuffer_Practice/Section7/ProtocolBuf-Example-GO/src/complex_example"
	enumpb "github.com/ProtocolBuffer_Practice/Section7/ProtocolBuf-Example-GO/src/enum_example"
	simplepb "github.com/ProtocolBuffer_Practice/Section7/ProtocolBuf-Example-GO/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	//readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	cx := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
		},
	}
	fmt.Println("Complex Message", cx)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           1234,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)
	fmt.Println("Json Value : ", smAsString)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully coverts json to Pb ", sm2)
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Can't covert json into pb", err)
	}
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to json")
		return ""
	}
	return out
}

func readAndWriteDemo(pb proto.Message) {
	//writeToFile
	writeToFile("simple.bin", pb)
	//readFromFile
	sm2 := &simplepb.SimpleMessage{}
	err := readFromFile("simple.bin", sm2)
	if err != nil {
		log.Fatalln("Read File failed", err)
	}
	fmt.Println("Read from file : ", sm2)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read file", err)
		return err
	}
	proto.Unmarshal(in, pb)
	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0664); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Write to file success")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SimpleList: []int32{1, 4, 7, 9},
	}
	fmt.Println(sm)
	sm.Name = "I renamed you"
	fmt.Println(sm)
	fmt.Println("The Id value is :", sm.GetId())
	return &sm
}
