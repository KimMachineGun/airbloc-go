package common

import (
	"bufio"
	"context"
	"io"
	"reflect"

	p2p "github.com/airbloc/airbloc-go/proto/p2p"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/pkg/errors"
)

type ProtoMessage struct {
	p2p.Message
}

func (message ProtoMessage) ID() peer.ID {
	return peer.ID(message.GetFrom())
}

func (message ProtoMessage) Pid() (Pid, error) {
	return ParsePid(string(message.GetProtocol()))
}

func ReadMessage(stream net.Stream) (ProtoMessage, error) {
	var raw []byte
	reader := bufio.NewReader(stream)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		raw = append(raw, p[:n]...)
	}

	msg := ProtoMessage{}
	err := proto.Unmarshal(raw, &msg)
	return msg, err
}

func (message ProtoMessage) WriteMessage(stream net.Stream) error {
	raw, err := proto.Marshal(&message)
	if err != nil {
		return errors.Wrap(err, "proto error : failed to marshal proto message")
	}

	if _, err := stream.Write(raw); err != nil {
		return errors.Wrap(err, "proto error : failed to write data to stream")
	}
	return nil
}

func (message ProtoMessage) MakeMessage(ctx context.Context, typ reflect.Type) (Message, error) {
	msg, ok := reflect.New(typ).Interface().(proto.Message)
	if !ok {
		return Message{}, errors.New("message is not protobuf message")
	}

	if err := proto.Unmarshal(message.GetData(), msg); err != nil {
		return Message{}, errors.Wrap(err, "failed to unmarshal data")
	}

	pid, err := message.Pid()
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to parse protocol")
	}
	return NewMessage(msg, peerstore.PeerInfo{ID: message.ID()}, pid), nil
}

func MessageType(msg proto.Message) reflect.Type {
	return reflect.ValueOf(msg).Elem().Type()
}

type Message struct {
	Data     proto.Message
	Info     peerstore.PeerInfo
	Protocol Pid
}

func NewMessage(data proto.Message, info peerstore.PeerInfo, protocol Pid) Message {
	return Message{
		Data:     data,
		Info:     info,
		Protocol: protocol,
	}
}