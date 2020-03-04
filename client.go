package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"
    "github.com/aws/aws-sdk-go/aws/endpoints"
    "github.com/aws/aws-sdk-go/aws"

    "fmt"
    "os"
)

import "C"

type SnsSession struct {
    Sess  *sns.SNS
}


//export SayHello
func SayHello(a int) {
    fmt.Printf("This is the hello world from golang %d\n", a)
    return
}

//export ListTopic
func ListTopic() {
    myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
        if service == endpoints.SnsServiceID {
            return endpoints.ResolvedEndpoint{
                URL:           "http://localhost:4575",
                SigningRegion: "us-east-1",
            }, nil
        }

        return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
    }

    sess := session.Must(session.NewSessionWithOptions(session.Options{
        Config: aws.Config{
            Region: aws.String("us-east-1"),
            EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
        },
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := sns.New(sess)

    result, err := svc.ListTopics(nil)
    if err != nil {
        fmt.Printf("error\n")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    for _, t := range result.Topics {
        fmt.Println(*t.TopicArn)
    }

}

func main() {}
