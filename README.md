# Hyperledger Fabric Tools

Extract proto.encode proto.decode compute.update from configtxlator to 
the rpc service.

## RPC Methods

### Proto.Encode

Converts a JSON document to protobuf.

```text
// Request params
type ProtoArgs struct {
	MsgName string // The type of protobuf structure to encode to. For example, 'common.Config'.
	Input   []byte // A file containing the JSON document.
}

// result
reply *[]byte // A file to write the output to.
```

### Proto.Decode

Converts a proto message to JSON.

```text
// Request params
type ProtoArgs struct {
	MsgName string // The type of protobuf structure to encode to. For example, 'common.Config'.
	Input   []byte // A file containing the JSON document.
}

// result
reply *[]byte
```

### Compute.Update

Takes two marshaled common.Config messages and computes the config update which transitions between the two.

```text
type ComputeUpdateArgs struct {
	ChannelName string // The name of the channel for this update.
	Origin      []byte // The original config message.
	Updated     []byte // The updated config message.
}

// result
reply *[]byte
```