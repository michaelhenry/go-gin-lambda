# go-gin-lambda template

![Gopher-life](https://blog.iamkel.net/assets/blog/go-lang/go-life.png)

A very simple way to deploy your go-gin application to aws lambda.

## How to use

Plase generate or get your AWS credentials from [AWS IAM](https://us-east-1.console.aws.amazon.com/iamv2/home?region=ap-southeast-2#/home) and store it as part of Github's Action Secret. And that's it. 

```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
```

> Note: For security purposes, please only include the necessary permission needed to build your aws stacks.
