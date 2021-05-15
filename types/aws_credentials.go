package types

const (
	AwsAuthTypeAccessKeys = "accessKeys"
	AwsAuthTypeAssumeRole = "assumeRole"
)

type AwsCredentials struct {
	AuthType             string `json:"authType"`
	AccessKeyId          string `json:"accessKeyId"`
	SecretAccessKey      string `json:"secretAccessKey"`
	AssumeRoleName       string `json:"assumeRoleName"`
	AssumeRoleExternalId string `json:"assumeRoleExternalId"`
}
