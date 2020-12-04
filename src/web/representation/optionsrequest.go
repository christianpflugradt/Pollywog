package representation

type OptionsRequest struct {
	KeepOptions []int `json:"keep_options"`
	CreateOptions []string `json:"create_options"`
}
