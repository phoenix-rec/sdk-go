package sdk_go

import (
	"fmt"
	"github.com/phoenix-rec/sdk-go/phoenix"
	"testing"
)

func TestWrite(t *testing.T) {
	ProjectId := int64(10000043)
	CustomerId := int64(1)
	AccessId := "accessId"
	AccessKey := "secretId"
	WriteUrl := "http://phoenix-api.icocofun.com/data/openapi/write/"

	client, err := (&phoenix.ClientBuilder{}).
		ProjectId(ProjectId).
		CustomerId(CustomerId).
		AccessId(AccessId).
		AccessKey(AccessKey).
		WriteUrl(WriteUrl).Build()

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.WriteData("item", "test", []map[string]interface{}{{
		"x_item_id": "1",
		"x_status":  1,
	}})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func TestRec(t *testing.T) {
	ProjectId := int64(10000001)
	CustomerId := int64(1)
	AccessId := "accessId"
	AccessKey := "secretId"
	RecUrl := "http://phoenix-api.icocofun.com/recommend/openapi/rec/"

	client, err := (&phoenix.ClientBuilder{}).
		ProjectId(ProjectId).
		CustomerId(CustomerId).
		AccessId(AccessId).
		AccessKey(AccessKey).
		RecUrl(RecUrl).Build()

	if err != nil {
		fmt.Println(err)
	}
	userId := "10001"
	did := "asdfghjkl"
	ip := "127.0.0.1"
	channel := "HuaiWei"
	network := int32(1)
	os := "Android"
	model := "HUAWEI P40"

	requestId := "123456789"
	resp, err := client.Rec(10001, phoenix.UserInfo{
		UserID:  userId,
		Did:     did,
		Ip:      ip,
		Channel: channel,
		Network: network,
		Os:      os,
		Model:   model,
	}, []phoenix.Option{
		phoenix.WithRequestId(requestId),
	}...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
