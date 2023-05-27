package table_operations

// Package to handle CRUD operations on DynamoDB table

import (
    "context"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

// Writes a new item to the DynamoDB table
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
    _, errr := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
        TableName: aws.String("Recipes"),
        Item: item,
    })

    if err != nil {
        panic(errr)
    }
}

// Scan/query Dyanmodb items
func ScanItems(r Recipe) []Recipe {
    cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
        o.Region = "us-east-1"
        return nil
    })
    if err != nil {
        panic(err)
    }

    svc := dynamodb.NewFromConfig(cfg)

    expr, err := expression.NewBuilder().WithFilter(
        expression.And(
            expression.Contains(expression.Name("RecipeName"), r.RecipeName),
            expression.Contains(expression.Name("Cuisine"), r.Cuisine),
            expression.Contains(expression.Name("Ingredients"), r.Ingredients),
        ),
    ).Build()
    if err != nil {
        panic(err)
    }

    out_raw, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
        TableName:                 aws.String("Recipes"),
        FilterExpression:          expr.Filter(),
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
    })
    if err != nil {
        panic(err)
    }

    // Unmarshals *dynamodb.ScanOutput to list of structs
    var out []Recipe
    err_unmarshal := attributevalue.UnmarshalListOfMaps(out_raw.Items, &out)
    if err_unmarshal != nil {
        fmt.Printf("Failed to unmarshal Items, %w", err_unmarshal)
    }

    return out
}
