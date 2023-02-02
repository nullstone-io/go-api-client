package types

type EnvVariables map[string]EnvVariable

type EnvVariable struct {
	Value     string `json:"value"`
	Sensitive bool   `json:"sensitive"`
	Redacted  bool   `json:"redacted"`
}

type EnvVariableInput struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Sensitive bool   `json:"sensitive"`
}
