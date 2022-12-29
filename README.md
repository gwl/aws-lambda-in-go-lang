# AWS Lambda in GoLang — The Ultimate Guide
*Revised: Monday, December 26, 2022*

A tutorial project for AWS Lambda in Golang 
**This tutorial project** builds with Go 1.19 on macOS for AWS Lambda/Linux

- [AWS Lambda in GoLang — The Ultimate Guide](https://www.softkraft.co/aws-lambda-in-golang/)


### What is this repository for? ###

This is a tutorial project which demonstrates the use of the DynamoDB persistent store with AWS Lambda in Go.

* [AWS Lambda](https://aws.amazon.com/lambda/?nc2=h_ql_prod_cp_lbd)
* [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)
* Go
* JSON

### How do I get set up? (initial build instructions) ###
The instructions in the tutorial are pretty good. These two additional bits might help...

1. Create the folder structure for this project in your ~/go project folder…
*  `% cd ~/go`
*  `% mkdir aws-lambda-in-go-lang`
*  `% cd aws-lambda-in-go-lang`
*  `% go mod init aws-lambda-in-go-lang`
*  `% mkdir build cmd pkg pkg/user pkg/handlers pkg/validators`
*  `% git init`

2. IMPORTANT BUILD NOTE: The tutorial has the remaining instructions, with one minor error. The author happens to be building on the same architecture as the AWS Lambda host (Linux/amd64). If your build system is somewhere else (say, macOS) you'll need to modify the build instruction, thusly:

* % `GOARCH=amd64 GOOS=linux go build -o build/main cmd/main.go`   [or, to build and zip it in one step…]
* % `GOARCH=amd64 GOOS=linux go build -o build/main cmd/main.go && zip -jrm build/main.zip build/main`


### Testing it ###
You can poke it with a web browser and it will provide the list of entires in the database as JSON.

You'll need to replace the URL with your own.

1. add a new user…

* % `curl --header "Content-Type: application/json" --request POST --data '{"email": "jonnyquest@example.com", "firstName": "Jonny", "lastName": "Quest"}' https://YOUR-Lambda-URL.execute-api.us-east-1.amazonaws.com/staging`

{"email":"jonnyquest@example.com","firstName":"Jonny","lastName":"Quest"}

2. list all users…

* % `curl -X GET https://YOUR-Lambda-URL.execute-api.us-east-1.amazonaws.com/staging`

[{"email":"s.karasiewicz@softkraft.co","firstName":"Sebastian","lastName":"Karasiewicz"},{"email":"jonnyquest@example.com","firstName":"Jonny","lastName":"Quest"},{"email":"mork@example.com","firstName":"Mork","lastName":"FromOrk"}]

3. fetch the record of a user given email address…

* % `curl -X GET https://YOUR-Lambda-URL.execute-api.us-east-1.amazonaws.com/staging\?email\=jonnyquest@example.com`

{"email":"jonnyquest@example.com","firstName":"Jonny","lastName":"Quest"}

4. update a user record

* `% curl --header "Content-Type: application/json" --request PUT --data '{"email": "s.karasiewicz@softkraft.co", "firstName": "Sebas", "lastName": "Karasiewicz"}' https://YOUR-Lambda-URL.execute-api.us-east-1.amazonaws.com/staging`

5. delete a user…

* % `curl -X DELETE https://YOUR-Lambda-URL.execute-api.us-east-1.amazonaws.com/staging\?email\=s.karasiewicz@softkraft.co`


### Resources ### 
[AWS Toolkit for JetBrains](https://docs.aws.amazon.com/toolkit-for-jetbrains/latest/userguide/welcome.html)


### Contribution guidelines ###

* This isn't a living project so there are no contribution guidelines.
* [Markdown Cheatsheet](https://www.markdownguide.org/cheat-sheet/)

