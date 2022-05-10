package aliyun

import (
	"AliyunLetsEncrypt/aliyun/models"
	"errors"
	cas20180713 "github.com/alibabacloud-go/cas-20180713/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	waf_openapi20190910 "github.com/alibabacloud-go/waf-openapi-20190910/client"
)

const DEFAULT_REGION_ID = "cn-hangzhou"

type WafOpenapiClient interface {
	// 实例
	DescribeInstanceInfo() (*waf_openapi20190910.DescribeInstanceInfoResponse, error) // 查询当前阿里云账号下WAF实例的详情

	// 域名
	DescribeDomainNames() (*waf_openapi20190910.DescribeDomainNamesResponse, error)                                                               // 获取指定WAF实例中已添加的域名名称列表
	CreateCertificate(domain, certificate, privateKey, certificateName *string) (*waf_openapi20190910.CreateCertificateResponse, error)           // 为已添加的域名配置记录上传证书及私钥信息
	DescribeCertificates(domain *string) (*waf_openapi20190910.DescribeCertificatesResponse, error)                                               // 为已添加的域名配置记录上传证书及私钥信息
	CreateCertificateByCertificateId(domain *string, certificateId *int64) (*waf_openapi20190910.CreateCertificateByCertificateIdResponse, error) // 根据证书ID为指定域名创建证书
}

type CasOpenapiClient interface {
	// 证书
	CreateUserCertificate(name, certificate, privateKey *string) (*cas20180713.CreateUserCertificateResponse, error) // 添加证书
}

type Aliyun interface {
	CreateWafOpenapiClient() (client WafOpenapiClient, _err error)
	CreateCasOpenapiClient() (client CasOpenapiClient, _err error)
}
type aliyunClient struct {
	accessKeyId     string
	accessKeySecret string
	regionId        string
}

func CreateAliyun(accessKeyId, accessKeySecret, regionId string) (aliyun Aliyun, err error) {
	if accessKeySecret == "" || accessKeyId == "" {
		return nil, errors.New("非法参数")
	}
	if regionId == "" {
		regionId = DEFAULT_REGION_ID
	}
	aliyun = &aliyunClient{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		regionId:        regionId,
	}
	return
}

func (aliyun *aliyunClient) CreateWafOpenapiClient() (client WafOpenapiClient, _err error) {
	var (
		ok      bool
		regions string
	)
	if regions, ok = models.WAF_REGIONS[aliyun.regionId]; !ok {
		_err = errors.New("非法参数")
		return
	}
	_result := &waf_openapi20190910.Client{}
	if _result, _err = waf_openapi20190910.NewClient(&openapi.Config{
		AccessKeyId:     &aliyun.accessKeyId,
		AccessKeySecret: &aliyun.accessKeySecret,
		Endpoint:        &regions,
	}); _err != nil {
		return
	} else {
		client = &wafOpenapiClient{
			client:     _result,
			instanceId: nil,
		}
	}
	return
}
func (aliyun *aliyunClient) CreateCasOpenapiClient() (client CasOpenapiClient, _err error) {
	var (
		ok      bool
		regions string
	)
	if regions, ok = models.CAS_REGIONS[aliyun.regionId]; !ok {
		_err = errors.New("非法参数")
		return
	}
	_result := &cas20180713.Client{}
	if _result, _err = cas20180713.NewClient(&openapi.Config{
		AccessKeyId:     &aliyun.accessKeyId,
		AccessKeySecret: &aliyun.accessKeySecret,
		Endpoint:        &regions,
	}); _err != nil {
		return
	} else {
		client = &casOpenapiClient{
			client: _result,
		}
	}
	return
}
