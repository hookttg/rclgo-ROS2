/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <sensor_msgs/srv/set_camera_info.h>
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
	typemap.RegisterService("sensor_msgs/SetCameraInfo", SetCameraInfoTypeSupport)
	typemap.RegisterService("sensor_msgs/srv/SetCameraInfo", SetCameraInfoTypeSupport)
}

type _SetCameraInfoTypeSupport struct {}

func (s _SetCameraInfoTypeSupport) Request() types.MessageTypeSupport {
	return SetCameraInfo_RequestTypeSupport
}

func (s _SetCameraInfoTypeSupport) Response() types.MessageTypeSupport {
	return SetCameraInfo_ResponseTypeSupport
}

func (s _SetCameraInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__sensor_msgs__srv__SetCameraInfo())
}

// Modifying this variable is undefined behavior.
var SetCameraInfoTypeSupport types.ServiceTypeSupport = _SetCameraInfoTypeSupport{}

// SetCameraInfoClient wraps rclgo.Client to provide type safe helper
// functions
type SetCameraInfoClient struct {
	*rclgo.Client
}

// NewSetCameraInfoClient creates and returns a new client for the
// SetCameraInfo
func NewSetCameraInfoClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*SetCameraInfoClient, error) {
	client, err := node.NewClient(serviceName, SetCameraInfoTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetCameraInfoClient{client}, nil
}

func (s *SetCameraInfoClient) Send(ctx context.Context, req *SetCameraInfo_Request) (*SetCameraInfo_Response, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*SetCameraInfo_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type SetCameraInfoServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s SetCameraInfoServiceResponseSender) SendResponse(resp *SetCameraInfo_Response) error {
	return s.sender.SendResponse(resp)
}

type SetCameraInfoServiceRequestHandler func(*rclgo.RmwServiceInfo, *SetCameraInfo_Request, SetCameraInfoServiceResponseSender)

// SetCameraInfoService wraps rclgo.Service to provide type safe helper
// functions
type SetCameraInfoService struct {
	*rclgo.Service
}

// NewSetCameraInfoService creates and returns a new service for the
// SetCameraInfo
func NewSetCameraInfoService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler SetCameraInfoServiceRequestHandler) (*SetCameraInfoService, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*SetCameraInfo_Request)
		responseSender := SetCameraInfoServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, SetCameraInfoTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &SetCameraInfoService{service}, nil
}