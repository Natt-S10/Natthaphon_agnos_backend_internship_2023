package models

type PasswordStepsReqest struct {
	Password string `json:"init_password" binding:"required" validate:"regexp=^[a-zA-Z0-9.!]*$"`
}

type PasswordStepsResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}

//,regexp=^[a-zA-Z0-9.!]*$
