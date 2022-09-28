package phoenix

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) apiRequest(url, requestId string, httpBody []byte) (resp []byte, err error) {
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := RandAllString(8)
	if requestId == "" {
		requestId = RandAllString(16)
	}
	signature := CalSignature(strconv.FormatInt(c.CustomerId, 10), timeNow, c.AccessId, c.AccessKey, nonce, httpBody)
	request, err := http.NewRequest("POST", url, bytes.NewReader(httpBody))
	if err != nil {
		return
	}
	// 设置请求头
	request.Header.Add("Customer-Id", strconv.FormatInt(c.CustomerId, 10))
	request.Header.Add("Access-Id", c.AccessId)
	request.Header.Add("Time", timeNow)
	request.Header.Add("Nonce", nonce)
	request.Header.Add("Request-Id", requestId)
	request.Header.Add("Signature", signature)
	// 发起请求
	httpResp, err := c.client.Do(request)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	resp, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	return
}

func (c *Client) WriteData(topic, stage string, param []map[string]interface{}, opts ...Option) (resp *DataResp, err error) {
	url := c.WriteUrl + strconv.FormatInt(c.ProjectId, 10) + "/" + topic + "/" + stage
	data, err := json.Marshal(param)
	if err != nil {
		return
	}
	respBody, err := c.apiRequest(url, Conv2Options(opts...).RequestId, data)
	json.Unmarshal(respBody, &resp)
	return
}

func (c *Client) Rec(scene int64, user UserInfo, opts ...Option) (resp *RecResp, err error) {
	url := c.RecUrl + strconv.FormatInt(c.ProjectId, 10)

	req := RecommendReq{
		CustomID:  c.CustomerId,
		ProjectID: c.ProjectId,
		SceneID:   scene,
		User:      user,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	respBody, err := c.apiRequest(url, Conv2Options(opts...).RequestId, data)
	json.Unmarshal(respBody, &resp)
	return
}
