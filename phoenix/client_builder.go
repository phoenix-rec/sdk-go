package phoenix

import (
	"fmt"
	"net/http"
)

type ClientBuilder struct {
	PId  int64
	CId  int64
	AK   string
	Sk   string
	WUrl string
	RUrl string
}

type Client struct {
	client     *http.Client
	ProjectId  int64
	CustomerId int64
	AccessId   string
	AccessKey  string
	WriteUrl   string
	RecUrl     string
}

func (receiver *ClientBuilder) ProjectId(projectId int64) *ClientBuilder {
	receiver.PId = projectId
	return receiver
}

func (receiver *ClientBuilder) CustomerId(customerId int64) *ClientBuilder {
	receiver.CId = customerId
	return receiver
}

func (receiver *ClientBuilder) AccessId(accessId string) *ClientBuilder {
	receiver.AK = accessId
	return receiver
}

func (receiver *ClientBuilder) AccessKey(AccessKey string) *ClientBuilder {
	receiver.Sk = AccessKey
	return receiver
}

func (receiver *ClientBuilder) WriteUrl(writeUrl string) *ClientBuilder {
	if writeUrl == "" {
		receiver.WUrl = defaultWriteUrl
	} else {
		receiver.WUrl = writeUrl
	}
	return receiver
}

func (receiver *ClientBuilder) RecUrl(recUrl string) *ClientBuilder {
	if recUrl == "" {
		receiver.RUrl = defaultRecUrl
	} else {
		receiver.RUrl = recUrl
	}
	return receiver
}

func (receiver *ClientBuilder) Build() (Client, error) {
	if receiver.CId == 0 {
		return Client{}, fmt.Errorf("customerId is empty")
	}
	if receiver.AK == "" {
		return Client{}, fmt.Errorf("accessKeyId is empty")
	}
	if receiver.Sk == "" {
		return Client{}, fmt.Errorf("accessKey is empty")
	}

	client := Client{
		ProjectId:  receiver.PId,
		CustomerId: receiver.CId,
		AccessId:   receiver.AK,
		AccessKey:  receiver.Sk,
		RecUrl:     receiver.RUrl,
		WriteUrl:   receiver.WUrl,
		client:     &http.Client{},
	}

	return client, nil
}
