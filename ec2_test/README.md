1. install AWS-SDK for Go
   - $go get github.com/aws/aws-sdk-go
   
2. create the Access Key of AWS-EC2
   - ref) https://docs.aws.amazon.com/ko_kr/sdk-for-go/v1/developer-guide/setting-up.html 
   
3. create the Credentials File
   - ref) https://docs.aws.amazon.com/ko_kr/sdk-for-go/v1/developer-guide/configuring-sdk.html
   - $vi $HOME/.aws/credentials
      [default]
      aws_access_key_id = <YOUR_ACCESS_KEY_ID>
      aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>

4. download this test program
   - $ go get github.com/powerkimhub/go-test/ec_test

5. $ cd $GOPATH/src/github.com/powerkimhub/go-test/ec2_test

6. $ go build

7. $ ./ec2_test
