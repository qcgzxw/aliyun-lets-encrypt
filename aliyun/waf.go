package aliyun

import (
	"errors"
	waf_openapi20190910 "github.com/alibabacloud-go/waf-openapi-20190910/client"
	"time"
)

type wafOpenapiClient struct {
	client     *waf_openapi20190910.Client
	instanceId *string
}

func (this *wafOpenapiClient) getClient() *waf_openapi20190910.Client {
	return this.client
}
func (this *wafOpenapiClient) getInstanceId() *string {
	if this.instanceId == nil {
		if resp, err := this.DescribeInstanceInfo(); err != nil {
			panic(err)
		} else {
			this.instanceId = resp.Body.InstanceInfo.InstanceId
		}
	}
	return this.instanceId
}

func (this *wafOpenapiClient) DescribeInstanceInfo() (resp *waf_openapi20190910.DescribeInstanceInfoResponse, err error) {
	r := &waf_openapi20190910.DescribeInstanceInfoRequest{}
	if resp, err = this.client.DescribeInstanceInfo(r); err != nil {
		return
	}
	invalidPayType := int32(0)
	if resp.Body.InstanceInfo == nil || resp.Body.InstanceInfo.PayType == nil || *resp.Body.InstanceInfo.PayType == invalidPayType || resp.Body.InstanceInfo.InstanceId == nil {
		panic("无法获取实例ID")
	}
	return
}

func (this *wafOpenapiClient) DescribeDomainNames() (resp *waf_openapi20190910.DescribeDomainNamesResponse, err error) {
	r := &waf_openapi20190910.DescribeDomainNamesRequest{
		InstanceId: this.getInstanceId(),
	}
	return this.client.DescribeDomainNames(r)
}

func (this *wafOpenapiClient) CreateCertificate(domain, certificate, privateKey, certificateName *string) (resp *waf_openapi20190910.CreateCertificateResponse, err error) {
	if domain == nil || certificate == nil || privateKey == nil ||
		*domain == "" || *certificate == "" || *privateKey == "" {
		err = errors.New("非法参数")
		return
	}
	if certificateName == nil || *certificateName == "" {
		now := time.Now().Format("20060102150403")
		certificateName = &now
	}
	r := &waf_openapi20190910.CreateCertificateRequest{
		Certificate:     certificate,
		CertificateName: certificateName,
		Domain:          domain,
		InstanceId:      this.getInstanceId(),
		PrivateKey:      privateKey,
	}
	if resp, err = this.client.CreateCertificate(r); err != nil {
		return
	}
	if resp.Body.CertificateId == nil {
		return nil, errors.New("上传失败")
	}
	return
}

func (this *wafOpenapiClient) DescribeCertificates(domain *string) (resp *waf_openapi20190910.DescribeCertificatesResponse, err error) {
	r := &waf_openapi20190910.DescribeCertificatesRequest{
		Domain:     domain,
		InstanceId: this.getInstanceId(),
	}
	return this.client.DescribeCertificates(r)
}
func (this *wafOpenapiClient) CreateCertificateByCertificateId(domain *string, certificateId *int64) (resp *waf_openapi20190910.CreateCertificateByCertificateIdResponse, err error) {
	r := &waf_openapi20190910.CreateCertificateByCertificateIdRequest{
		Domain:        domain,
		CertificateId: certificateId,
		InstanceId:    this.getInstanceId(),
	}
	return this.client.CreateCertificateByCertificateId(r)
}
