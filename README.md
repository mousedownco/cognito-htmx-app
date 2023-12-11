> [!IMPORTANT]
> This was a fork of the `htmx-contact-app` repository that was meant to show how to integrate AWS Cognito with HTMX. I've move that demo to a the simpler [`htmx-cognito-demo`](https://github.com/mousedownco/htmx-cognito-demo) project.



## Cognito Integration

Fork of the `htmx-contact-app` repository to add Cognito integration.

## Building

GOOS=linux GOARCH=arm64 sam build

## Custom Domain

Why use a custom domain?  When attempting to use the default distribution domains with the API Gateway the SAM template would eventually fail with circular dependency errors.