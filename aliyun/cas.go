package aliyun

import (
	"errors"
	cas20180713 "github.com/alibabacloud-go/cas-20180713/client"
	"time"
)

type casOpenapiClient struct {
	client *cas20180713.Client
}

func (this *casOpenapiClient) getClient() *cas20180713.Client {
	return this.client
}

func (this *casOpenapiClient) CreateUserCertificate(name, certificate, privateKey *string) (resp *cas20180713.CreateUserCertificateResponse, err error) {
	if certificate == nil || privateKey == nil || *certificate == "" || *privateKey == "" {
		err = errors.New("非法参数")
		return
	}
	if name == nil || *name == "" {
		now := time.Now().Format("20060102150403")
		name = &now
	}
	createUserCertificateRequest := &cas20180713.CreateUserCertificateRequest{
		Name: name,
		Cert: certificate,
		Key:  privateKey,
	}
	return this.getClient().CreateUserCertificate(createUserCertificateRequest)
}
