package openai

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func ImgFromText(text string) (img string, err error) {
	c := openai.NewClient(os.Getenv("OpenAISecret"))
	ctx := context.Background()

	reqUrl := openai.ImageRequest{
		Prompt:         text,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		return
	}
	img = respUrl.Data[0].URL
	return
}