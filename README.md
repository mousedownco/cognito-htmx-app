Hypermedia Systems - Contact App
================================

## Cognito Integration

Fork of the `htmx-contact-app` repository to add Cognito integration.

## Building

GOOS=linux GOARCH=arm64 sam build

## Custom Domain

Why use a custom domain?  When attempting to use the default distribution domains with the API Gateway the SAM template would eventually fail with circular dependency errors.