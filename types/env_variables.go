package types

type EnvVariables map[string]EnvVariable

type EnvVariable struct {
	Value     string `json:"value"`
	Sensitive bool   `json:"sensitive"`
	Redacted  bool   `json:"redacted"`

	// Secret is populated during redaction when Value was a `{{ secret(<id>) }}`
	// reference to a cloud secrets manager. It lets the dashboard surface the
	// managed secret as a first-class object instead of an opaque mask.
	Secret *SecretIdentity `json:"secret,omitempty"`
}

type EnvVariableInput struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Sensitive bool   `json:"sensitive"`
}
