package aiService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

type jsonResp struct {
	Status int    `json:"status" jsonschema:"title=status,description=博客友链申请审核的结果，内容正常返回1(int类型),内容有问题的话返回0(int类型),example=1,example=0"`
	Msg    string `json:"msg" jsonschema:"title=msg,description=博客友链申请审核的结果的原因，内容正常返回申请成功,内容有问题的话返回具体有问题的地方,example=申请成功,example=内容包含漏洞注入"`
}

func VerifyFriendLink(link, title, desc string) (int, string, error) {

	start := time.Now()
	defer func() {
		log.Printf("CreateNewFriendLinkReq aiService took %v", time.Since(start))
	}()
	type content struct {
		Link  string `json:"link"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	}

	// 将结构体序列化为 JSON 字符串
	contentJSON, err := json.Marshal(content{Link: link, Title: title, Desc: desc})
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return 0, "", fmt.Errorf("marshaling json: %w", err)
	}

	// 将 JSON 字符串转换为 string
	contentStr := string(contentJSON)
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemRole,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: systemRole,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: contentStr,
		},
	}

	//resp, err := schemaClient.CreateChatCompletion(
	//	ctx,
	//	openai.ChatCompletionRequest{
	//		Model:    "deepseek-chat",
	//		Messages: dialogue,
	//	},
	//	&jsonRespObj,
	//)

	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: dialogue,
		},
	)
	if err != nil {
		return 0, "", fmt.Errorf("CreateChatCompletion err: %w", err)
	}

	respStr := resp.Choices[0].Message.Content
	respObj, err := linkRespStrToModel(respStr)
	if err != nil {
		return 0, "", fmt.Errorf("linkRespStrToModel err: %w", err)
	}
	return respObj.Status, respObj.Msg, nil
}

func linkRespStrToModel(respStr string) (jsonResp, error) {
	var resp jsonResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		return jsonResp{}, err
	}
	return resp, nil
}
