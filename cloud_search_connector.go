package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
)

func main() {
	//customers
	search := "(prefix field=name 'かす')"
	queryparser := "structured"
	start := int64(0)
	region := "ap-northeast-1"
	return_str := "id,userhash,name,ipcompanyid,company,ban,infoemail,tags"
	sort_str := "chatno desc"
	endpoint := "search-customers-vjydyrbgzy45wijs7r2z2s4bfq.ap-northeast-1.cloudsearch.amazonaws.com"
	svc := cloudsearchdomain.New(session.New(), aws.NewConfig().WithRegion(region).WithEndpoint(endpoint))
	params := &cloudsearchdomain.SearchInput{
		Query:       aws.String(search),
		QueryParser: &queryparser,
		Start:       &start,
		Size:        aws.Int64(100),
		Return:      aws.String(return_str),
		Sort:        aws.String(sort_str),
	}
	resp, err := svc.Search(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
