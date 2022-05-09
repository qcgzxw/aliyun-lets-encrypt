package aliyun

import (
	"AliyunLetsEncrypt/aliyun/models"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	waf_openapi20190910 "github.com/alibabacloud-go/waf-openapi-20190910/client"
	"time"
)

const DEFAULT_REGION_ID = "cn-hangzhou"

type WafOpenapiClient interface {
	// 实例
	DescribeInstanceInfo() (*waf_openapi20190910.DescribeInstanceInfoResponse, error) // 查询当前阿里云账号下WAF实例的详情

	// 域名
	DescribeDomainNames() (*waf_openapi20190910.DescribeDomainNamesResponse, error)                                                    // 获取指定WAF实例中已添加的域名名称列表
	CreateCertificate(domain, certificate, privateKey, certificateName string) (*waf_openapi20190910.CreateCertificateResponse, error) // 为已添加的域名配置记录上传证书及私钥信息
}

type wafOpenapiClient struct {
	client     *waf_openapi20190910.Client
	instanceId *string
}

func CreateWafOpenapiClient(regionId string, accessKeyId *string, accessKeySecret *string) (client WafOpenapiClient, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	if regionId == "" {
		config.Endpoint = tea.String(DEFAULT_REGION_ID)
	} else {
		if region, ok := models.REGIONS[regionId]; !ok {
			return nil, errors.New("非法region id")
		} else {
			config.Endpoint = tea.String(region)
		}
	}
	_result := &waf_openapi20190910.Client{}
	if _result, _err = waf_openapi20190910.NewClient(config); _err == nil {
		client = &wafOpenapiClient{
			client:     _result,
			instanceId: nil,
		}
	}
	return
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
	if resp.Body.InstanceInfo == nil || resp.Body.InstanceInfo.PayType == tea.Int32(0) || resp.Body.InstanceInfo.InstanceId == nil {
		panic("无法获取实例ID")
	}
	return
}

func (this *wafOpenapiClient) DescribeDomainNames() (resp *waf_openapi20190910.DescribeDomainNamesResponse, err error) {
	r := &waf_openapi20190910.DescribeDomainNamesRequest{}
	return this.client.DescribeDomainNames(r)
}

func (this *wafOpenapiClient) CreateCertificate(domain, certificate, privateKey, certificateName string) (resp *waf_openapi20190910.CreateCertificateResponse, err error) {
	if certificateName == "" {
		certificateName = time.Now().Format("20060102150403")
	}
	r := &waf_openapi20190910.CreateCertificateRequest{
		Certificate:     tea.String(certificate),
		CertificateName: tea.String(certificateName),
		Domain:          tea.String(domain),
		InstanceId:      this.getInstanceId(),
		PrivateKey:      tea.String(privateKey),
	}
	if resp, err = this.client.CreateCertificate(r); err != nil {
		return
	}
	if resp.Body.CertificateId == nil {
		return nil, errors.New("上传失败")
	}
	return
}
