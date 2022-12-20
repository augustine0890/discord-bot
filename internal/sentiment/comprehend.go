package sentiment

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type AwsClient struct {
	comprehendClient *comprehend.Comprehend
}

func NewAwsClient() *AwsClient {
	// Create a Session with a custom region
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			os.Getenv("AWS_SESSION_TOKEN"),
		),
	}))

	// Create a Comprehend client from Session
	client := &AwsClient{comprehend.New(sess)}
	return client
}

func (c *AwsClient) DetectSentiment(text string) (result *comprehend.DetectSentimentOutput, err error) {
	params := &comprehend.DetectSentimentInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	}

	req, result := c.comprehendClient.DetectSentimentRequest(params)
	err = req.Send()
	if err != nil {
		return nil, err
	}

	return result, nil
}
