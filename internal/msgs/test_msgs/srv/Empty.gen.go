/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/srv/empty.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("test_msgs/Empty", EmptyTypeSupport)
	typemap.RegisterService("test_msgs/srv/Empty", EmptyTypeSupport)
}

type _EmptyTypeSupport struct {}

func (s _EmptyTypeSupport) Request() types.MessageTypeSupport {
	return Empty_RequestTypeSupport
}

func (s _EmptyTypeSupport) Response() types.MessageTypeSupport {
	return Empty_ResponseTypeSupport
}

func (s _EmptyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__test_msgs__srv__Empty())
}

// Modifying this variable is undefined behavior.
var EmptyTypeSupport types.ServiceTypeSupport = _EmptyTypeSupport{}

// EmptyClient wraps rclgo.Client to provide type safe helper
// functions
type EmptyClient struct {
	*rclgo.Client
}

// NewEmptyClient creates and returns a new client for the
// Empty
func NewEmptyClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*EmptyClient, error) {
	client, err := node.NewClient(serviceName, EmptyTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &EmptyClient{client}, nil
}

func (s *EmptyClient) Send(ctx context.Context, req *Empty_Request) (*Empty_Response, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Empty_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type EmptyServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s EmptyServiceResponseSender) SendResponse(resp *Empty_Response) error {
	return s.sender.SendResponse(resp)
}

type EmptyServiceRequestHandler func(*rclgo.RmwServiceInfo, *Empty_Request, EmptyServiceResponseSender)

// EmptyService wraps rclgo.Service to provide type safe helper
// functions
type EmptyService struct {
	*rclgo.Service
}

// NewEmptyService creates and returns a new service for the
// Empty
func NewEmptyService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler EmptyServiceRequestHandler) (*EmptyService, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Empty_Request)
		responseSender := EmptyServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, EmptyTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &EmptyService{service}, nil
}