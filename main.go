package main

import (
	"AliyunLetsEncrypt/aliyun"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
	"os"
)

func main() {
	// args1: accessKeyId
	// args2: accessKeySecret
	// args3: domain
	// args4: certificate path
	// args5: privateKey path

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
	aliyunClient, err := aliyun.CreateAliyun(accessKeyId, accessKeySecret, "")
	if err != nil {
		panic(err)
	}
	waf, err := aliyunClient.CreateWafOpenapiClient()
	if err != nil {
		panic(err)
	}
	if resp, err := waf.DescribeDomainNames(); err != nil {
		panic(err)
	} else {
		if len(resp.Body.DomainNames) == 0 {
			panic("无法获取WAF域名列表")
		}
		for key, item := range resp.Body.DomainNames {
			if domain == *item {
				break
			}
			if key == len(resp.Body.DomainNames)-1 {
				panic("域名未添加到WAF")
			}
		}
	}
	cas, err := aliyunClient.CreateCasOpenapiClient()
	if err != nil {
		panic(err)
	}
	var certId *int64
	if resp, err := cas.CreateUserCertificate(nil, tea.String(string(certificate)), tea.String(string(privateKey))); err != nil {
		panic(err)
	} else {
		if resp.Body == nil || resp.Body.CertId == nil {
			panic("上传失败")
		}
		certId = resp.Body.CertId
	}
	_, err = waf.CreateCertificateByCertificateId(&domain, certId)
	if err != nil {
		panic(err)
	}
	println("上传成功！")
}
