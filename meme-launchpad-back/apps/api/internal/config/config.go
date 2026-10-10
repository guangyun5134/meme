package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret  string `json:"accessSecret"`
		AccessExpire  int64  `json:"accessExpire"`
		RefreshExpire int64  `json:"refreshExpire"`
	} `json:"auth"`

	Database struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		User         string `json:"user"`
		Password     string `json:"password"`
		Name         string `json:"name"`
		SSLMode      string `json:"sslMode"`
		MaxOpenConns int    `json:"maxOpenConns"`
		MaxIdleConns int    `json:"maxIdleConns"`
	} `json:"database"`

	Redis struct {
		Host string `json:"host"`
		Pass string `json:"pass"`
		DB   int    `json:"db"`
	} `json:"redis"`

	Chain struct {
		Name            string `json:"name"`
		ChainID         int64  `json:"chainId"`
		RPC             string `json:"rpc"`
		CoreContract    string `json:"coreContract"`
		FactoryContract string `json:"factoryContract"`
		HelperContract  string `json:"helperContract"`
		VestingContract string `json:"vestingContract"`
		TokenBytecode   string `json:"tokenBytecode,omitempty"` // MEMEToken 合约字节码（用于 CREATE2 地址计算）
	} `json:"chain"`

	Storage struct {
		Type      string `json:"type"`
		Bucket    string `json:"bucket"`
		Region    string `json:"region"`
		AccessKey string `json:"accessKey,omitempty"`
		SecretKey string `json:"secretKey,omitempty"`
		Endpoint  string `json:"endpoint,omitempty"`
		CDNDomain string `json:"cdnDomain"`
	} `json:"storage"`

	// 腾讯云 COS 配置
	Cos struct {
		SecretID  string `json:"secretId,omitempty"`
		SecretKey string `json:"secretKey,omitempty"`
		Bucket    string `json:"bucket,omitempty"`
		Region    string `json:"region,omitempty"`
		AppID     string `json:"appId,omitempty"`
		Domain    string `json:"domain,omitempty"` // 自定义域名
	} `json:"cos"`

	SignerPrivateKey string `json:"signerPrivateKey,omitempty"`
}
