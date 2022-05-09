package main

import (
	"AliyunLetsEncrypt/aliyun"
	"io/ioutil"
	"os"
)

func main() {
	// args1: accessKeyId
	// args2: accessKeySecret
	// args3: domain
	// args4: certificate path
	// args5: privateKey path
	const DEFAULT_REGION = "cn-hangzhou"

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 5 {
		panic("参数缺失")
	}
	accessKeyId := argsWithoutProg[0]
	accessKeySecret := argsWithoutProg[1]
	domain := argsWithoutProg[2]
	certificatePath := argsWithoutProg[3]
	privateKeyPath := argsWithoutProg[4]

	certificate, err := ioutil.ReadFile(certificatePath)
	if err != nil {
		panic(err)
	}
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	wafClient, err := aliyun.CreateWafOpenapiClient(DEFAULT_REGION, &accessKeyId, &accessKeySecret)
	if err != nil {
		panic(err)
	}
	_, err = wafClient.CreateCertificate(domain, string(certificate), string(privateKey), "")
	if err != nil {
		panic(err)
	}
	println("上传成功！")
}
