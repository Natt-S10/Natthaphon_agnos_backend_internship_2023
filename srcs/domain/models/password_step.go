package models

type PasswordStepsReqest struct {
	Password string `json:"init_password"`
}

type PasswordStepsResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
