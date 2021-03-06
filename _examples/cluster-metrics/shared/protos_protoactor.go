// Package shared is generated by protoactor-go/protoc-gen-gograin@0.1.0
package shared

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/gogo/protobuf/proto"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf


var xHelloFactory func() Hello

// HelloFactory produces a Hello
func HelloFactory(factory func() Hello) {
	xHelloFactory = factory
}

// GetHelloGrainClient instantiates a new HelloGrainClient with given ID
func GetHelloGrainClient(c *cluster.Cluster, id string) *HelloGrainClient {
	if c == nil {
		panic(fmt.Errorf("nil cluster instance"))
	}
	if id == "" {
		panic(fmt.Errorf("empty id"))
	}
	return &HelloGrainClient{ID: id, cluster: c}
}

// Hello interfaces the services available to the Hello
type Hello interface {
	Init(id string)
	Terminate()
	ReceiveDefault(ctx actor.Context)
	SayHello(*HelloRequest, cluster.GrainContext) (*HelloResponse, error)
	Add(*AddRequest, cluster.GrainContext) (*AddResponse, error)
	VoidFunc(*AddRequest, cluster.GrainContext) (*Unit, error)
	
}

// HelloGrainClient holds the base data for the HelloGrain
type HelloGrainClient struct {
	ID      string
	cluster *cluster.Cluster
}

// SayHello requests the execution on to the cluster with CallOptions
func (g *HelloGrainClient) SayHello(r *HelloRequest, opts ...*cluster.GrainCallOptions) (*HelloResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 0, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Hello", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &HelloResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// Add requests the execution on to the cluster with CallOptions
func (g *HelloGrainClient) Add(r *AddRequest, opts ...*cluster.GrainCallOptions) (*AddResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 1, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Hello", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &AddResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// VoidFunc requests the execution on to the cluster with CallOptions
func (g *HelloGrainClient) VoidFunc(r *AddRequest, opts ...*cluster.GrainCallOptions) (*Unit, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 2, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Hello", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &Unit{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}


// HelloActor represents the actor structure
type HelloActor struct {
	inner Hello
	Timeout time.Duration
}

// Receive ensures the lifecycle of the actor for the received message
func (a *HelloActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.inner = xHelloFactory()
		id := ctx.Self().Id[17:] // skip "activator/Remote$"
		a.inner.Init(id)
		if a.Timeout > 0 {
			ctx.SetReceiveTimeout(a.Timeout)
		}
	case *actor.ReceiveTimeout:
		a.inner.Terminate()
		ctx.Poison(ctx.Self())

	case actor.AutoReceiveMessage: // pass
	case actor.SystemMessage: // pass

	case *cluster.GrainRequest:
		switch msg.MethodIndex {
			
		case 0:
			req := &HelloRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.SayHello(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 1:
			req := &AddRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.Add(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
			
		case 2:
			req := &AddRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.VoidFunc(req, ctx)
			if err == nil {
				bytes, errMarshal := proto.Marshal(r0)
				if errMarshal != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", errMarshal)
				}
				resp := &cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
		
		}
	default:
		a.inner.ReceiveDefault(ctx)
	}
}
