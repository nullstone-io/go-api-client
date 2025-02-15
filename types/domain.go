package types

type Domain struct {
	Block       `json:",inline"`
	Registrar   string `json:"registrar"`
	Certificate string `json:"certificate"`
}
