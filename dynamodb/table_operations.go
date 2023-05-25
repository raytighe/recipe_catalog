package table_operations

// Package to handle CRUD operations on DynamoDB table

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type Recipe struct {
    RecipeId int
    RecipeName string
    Cuisine string
    Ingredients string
    Instructions string
    Source string
    CookTime int
}

// Function that returns a boolean if specified DynamoDB table exists
func CheckTable(TableName string) bool {
    cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
        opts.Region = "us-east-1"
        return nil
    })
    if err != nil {
        panic(err)
    }

    svc := dynamodb.NewFromConfig(cfg)
    p := dynamodb.NewListTablesPaginator(svc, nil, func(o *dynamodb.ListTablesPaginatorOptions) {
        o.StopOnDuplicateToken = true
    })

    for p.HasMorePages() {
        out, err := p.NextPage(context.TODO())
        if err != nil {
            panic(err)
        }

        for _, tn := range out.TableNames {
            if tn == TableName {
                return true
            }
        }
    }
    return false
}

func WriteItem(r Recipe) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
        o.Region = "us-east-1"
        return nil
    })
    if err != nil {
        panic(err)
    }

    item, err := attributevalue.MarshalMap(r)
    svc := dynamodb.NewFromConfig(cfg)
    out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
        TableName: aws.String("Recipes"),
        Item: item,
    })

    if err != nil {
        panic(err)
    }

    fmt.Println(out.Attributes)
}