package services

type ProducerTokenService interface {
	ValidateToken() bool
	GetClaims() Claims
	ValidateClaims(claims_validate Claims) bool
}
type ConsumerTokenService interface {
	RequestToken() []map[string]interface{}
}
type Claims struct {
	KV []map[string]interface{}
}

type AzureAuthToken struct {
	Token            string
	ClaimsToValidate Claims
	Client_ID        string
}

func (a *AzureAuthToken) ValidateToken() bool {
	panic("Not Implemented")
}
func (a *AzureAuthToken) GetClaims() bool {
	panic("Not Implemented")
}

func (a *AzureAuthToken) ValidateClaims() bool {
	panic("Not Implemented")
}

type AzureConsumerClient struct {
	Client_ID     string
	Client_Secret string
	Scope         string
}

func (a *AzureConsumerClient) RequestToken() []map[string]interface{} {
	panic("Not Implemented")
}
