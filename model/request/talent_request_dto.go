package request

import "github.com/google/uuid"

type TalentLoginRequestDTO struct {
	Email    string
	Password string
}

type TalentRegisterRequestDTO struct {
	Email       string
	Password    string
	FullName    string
	PhoneNumber string
}

type GetAllTalentByAppIDRequestDTO struct {
	AppID uuid.UUID `json:"app_id"`
}

type EditTalentRequestDTO struct {
	Email       string
	FullName    string
	PhoneNumber string
	CV          string
	Bio         string
	Location    string
	Skills      []string
}

type GetAllTalentByAppAndJobIDRequestDTO struct {
	AppID uuid.UUID `json:"app_id"`
	JobID uuid.UUID `json:"job_id"`
}
