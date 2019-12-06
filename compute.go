package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/tools/configtxlator/update"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

type Compute int

type ComputeUpdateArgs struct {
	ChannelName string
	Origin      []byte
	Updated     []byte
}

func (*Compute) Update(args *ComputeUpdateArgs, reply *[]byte) error {
	origConf := &cb.Config{}
	err := proto.Unmarshal(args.Origin, origConf)
	if err != nil {
		return errors.Wrapf(err, "error unmarshaling original config")
	}

	updtConf := &cb.Config{}
	err = proto.Unmarshal(args.Updated, updtConf)
	if err != nil {
		return errors.Wrapf(err, "error unmarshaling updated config")
	}

	cu, err := update.Compute(origConf, updtConf)
	if err != nil {
		return errors.Wrapf(err, "error computing config update")
	}

	cu.ChannelId = args.ChannelName

	*reply, err = proto.Marshal(cu)
	if err != nil {
		return errors.Wrapf(err, "error marshaling computed config update")
	}

	return nil
}
