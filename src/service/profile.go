package service

import "github.com/dalikewara/ayapingping-go-crud/src/repository"

type profile struct {
	profileRepo repository.Profile
}

type NewProfileParam struct {
	ProfileRepo repository.Profile
}

// NewProfile generates new profile service that implements Profile.
func NewProfile(param NewProfileParam) Profile {
	return &profile{
		profileRepo: param.ProfileRepo,
	}
}

// UpdateImage updates profile image.
func (s *profile) UpdateImage(param ProfileUpdateImageParam) ProfileUpdateImageResult {
	result := ProfileUpdateImageResult{}

	if !param.UserID.IsValid() {
		result.Error = ErrParamUserID
		return result
	}

	return result
}
