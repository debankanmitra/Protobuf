// in protocol buffer here we are doing conversion to json , read from file , write to file (the binary )
// we are not sending the message format through wire , for that we need grpc (which sends proto messages)

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	pkg "proto/go_proto/src/simple"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := dosimple()
	fmt.Println(sm)

	WritetoFile("simple.txt", sm)

	sm2 := &pkg.Simplemessage{}
	readFromfile("simple.txt", sm2)
	fmt.Println(sm2)

	toJSON(sm)

}

// creating a object in golang
func dosimple() *pkg.Simplemessage {
	sm := pkg.Simplemessage{ // HERE WE ARE CREATING OBJECT
		Name:  "DEBANKAN MITRA",
		Class: 14,
		Roll:  12,
	}
	return &sm
}

func WritetoFile(fname string, pb proto.Message) {
	// This function is used to encode a protocol buffer message pb into its wire-format encoding.
	// The wire-format encoding is a binary representation of the protocol buffer message that can be transmitted or stored.
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("saved to file")
	}

}

func readFromfile(fname string, pb proto.Message) {
	data, _ := ioutil.ReadFile(fname)
	fmt.Println("THE BINARY SERIALIZATION DATA IS ", data)

	err := proto.Unmarshal(data, pb)
	if err != nil {
		log.Fatal(err)
	}

}

// protobuf format TO JSON
func toJSON(pb proto.Message) {
	// marshaler := jsonpb.Marshaler{}
	// data, err := marshaler.MarshalToString()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data, _ := protojson.Marshal(pb) // unmarshal for json to protobuf message

	fmt.Println("to json", string(data))
}
