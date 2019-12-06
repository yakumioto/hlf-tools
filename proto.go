package main

import (
	"bytes"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/tools/protolator"
	_ "github.com/hyperledger/fabric/protos/common"
	_ "github.com/hyperledger/fabric/protos/msp"
	_ "github.com/hyperledger/fabric/protos/orderer"
	_ "github.com/hyperledger/fabric/protos/orderer/etcdraft"
	_ "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

type Proto int

type ProtoArgs struct {
	MsgName string
	Input   []byte
}

func (*Proto) Decode(args *ProtoArgs, reply *[]byte) error {
	msgType := proto.MessageType(args.MsgName)
	if msgType == nil {
		return errors.Errorf("message of type %s unknown", msgType)
	}
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)

	err := proto.Unmarshal(args.Input, msg)
	if err != nil {
		return errors.Wrapf(err, "error unmarshaling")
	}

	buf := bytes.NewBuffer(nil)
	err = protolator.DeepMarshalJSON(buf, msg)
	if err != nil {
		return errors.Wrapf(err, "error encoding output")
	}

	*reply = buf.Bytes()

	return nil
}

func (*Proto) Encode(args *ProtoArgs, reply *[]byte) error {
	msgType := proto.MessageType(args.MsgName)
	if msgType == nil {
		return errors.Errorf("message of type %s unknown", msgType)
	}
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)

	input := bytes.NewBuffer(args.Input)

	err := protolator.DeepUnmarshalJSON(input, msg)
	if err != nil {
		return errors.Wrapf(err, "error decoding input")
	}

	*reply, err = proto.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "error marshaling")
	}

	return nil
}
