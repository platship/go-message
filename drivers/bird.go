package drivers

import (
	"encoding/json"
	"errors"

	"github.com/rehok/go-message/dto"
	"github.com/rehok/go-utils/curlx"
)

const (
	BaseUrl         = "https://api.bird.com"
	BirdCallbackUrl = ""
)

type Bird struct {
	AccessKey   string `json:"accessKey"`
	WorkspaceId string `json:"workspaceId"`
	SigningKey  string `json:"signingKey"`
}

func (Bird) Set(accessKey, workspaceId string) *Bird {
	return &Bird{AccessKey: accessKey, WorkspaceId: workspaceId, SigningKey: "bird"}
}

func (c *Bird) headers() []*curlx.Headers {
	var headers []*curlx.Headers
	headers = append(headers, &curlx.Headers{
		Name:  "Authorization",
		Value: "AccessKey " + c.AccessKey,
	})
	headers = append(headers, &curlx.Headers{
		Name:  "Content-Type",
		Value: "application/json",
	})
	return headers
}

// 获取钩子列表
func (c *Bird) WebhookSubscriptionList(organizationId string) (res []*dto.BirdResultBase, err error) {
	path := BaseUrl + "/organizations/" + organizationId + "/workspaces/" + c.WorkspaceId + "/webhook-subscriptions"
	data, err := curlx.Get(path, c.headers()...)
	if err != nil {
		return nil, err
	}
	var birdCallback dto.BirdCallback
	json.Unmarshal(data, &birdCallback)
	return birdCallback.Results, nil
}

// 更新钩子
func (c *Bird) WebhookSubscriptionPacth(organizationId, id string) (err error) {
	path := BaseUrl + "/organizations/" + organizationId + "/workspaces/" + c.WorkspaceId + "/webhook-subscriptions"
	newData := make(map[string]interface{})
	newData["url"] = BirdCallbackUrl
	newData["signingKey"] = c.SigningKey
	newData["status"] = "active"
	_, err = curlx.Request("PATCH", path+"/"+id, newData, c.headers()...)
	if err != nil {
		return errors.New("更新钩子失败")
	}
	return
}

func (c *Bird) WebhookSubscriptionPost(organizationId, event string) (err error) {
	path := BaseUrl + "/organizations/" + organizationId + "/workspaces/" + c.WorkspaceId + "/webhook-subscriptions"
	var datas = make(map[string]interface{})
	datas["service"] = "channels"
	datas["url"] = BirdCallbackUrl
	datas["event"] = event
	// datas["eventFilters"] = [] // 事件过滤
	_, err = curlx.Post(path, datas, c.headers()...)
	if err != nil {
		return errors.New("添加钩子失败")
	}
	return
}

type TextStr struct {
	Text string `json:"text"`
}

type MessageBody struct {
	Type string  `json:"type"`
	Text TextStr `json:"text"`
}

type MessageReceiver struct {
	Contacts []Contacts `json:"contacts"`
}

type Message struct {
	Body     MessageBody     `json:"body"`
	Receiver MessageReceiver `json:"receiver"`
}

type Contacts struct {
	IdentifierValue string `json:"identifierValue"`
	IdentifierKey   string `json:"identifierKey"`
}

func (c *Bird) ChannelMessagePost(channelId, phone, text string) (id string, err error) {
	path := BaseUrl + "/workspaces/" + c.WorkspaceId + "/channels/{channelId}/messages"
	var contacts []Contacts
	contacts = append(contacts, Contacts{IdentifierKey: "phonenumber", IdentifierValue: phone})
	datas := &Message{
		MessageBody{Type: "text", Text: TextStr{Text: text}},
		MessageReceiver{Contacts: contacts},
	}
	result, err := curlx.Post(path, datas, c.headers()...)
	if err != nil {
		return "", errors.New("添加钩子失败")
	}
	var res *dto.ResChannelMessage
	json.Unmarshal(result, &res)
	return res.ID, nil
}

func (s *Bird) Init() (err error) {
	if s.AccessKey == "" {
		return errors.New("API Key is undefined")
	}
	if s.AccessKey == "" {
		return errors.New("API Url is undefined")
	}
	return
}

func (s *Bird) Send() error {
	// 发送信息
	return nil
}
