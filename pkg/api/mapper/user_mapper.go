package mapper

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
)

// UserFromModel is method for mapping db model to dto
func UserFromModel(user *model.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:       user.ID.Hex(),
		Account:  userAccountFromModel(user.Account),
		Personal: userPersonalFromModel(user.Personal),
	}
}

func userPersonalFromModel(personal model.PersonalInformation) dto.PersonalInformationDTO {
	return dto.PersonalInformationDTO{
		Email:     personal.Email,
		BirthDate: personal.BirthDate,
		FirstName: personal.FirstName,
		LastName:  personal.LastName,
		Address:   userAddressFromModel(personal.Address),
	}
}

func userAddressFromModel(address model.UserAddress) dto.UserAddressDTO {
	return dto.UserAddressDTO{
		City:       address.City,
		Country:    address.Country,
		Line1:      address.Line1,
		Line2:      address.Line2,
		PostalCode: address.PostalCode,
		Region:     address.Region,
	}
}

func userAccountFromModel(account model.UserAccount) dto.UserAccountDTO {
	return dto.UserAccountDTO{
		Nickname:            account.Nickname,
		PrimaryLanguage:     account.PrimaryLanguage,
		AdditionalLanguages: getArrayOrEmpty(account.AdditionalLanguages),
		Socials:             userSocialsFromModel(account.Socials),
	}
}

func getArrayOrEmpty(arr []string) []string {
	if arr == nil {
		return []string{}
	}
	return arr
}

func userSocialsFromModel(accounts []model.UserSocialAccount) []dto.UserSocialAccountDTO {
	socials := []dto.UserSocialAccountDTO{}
	for _, account := range accounts {
		socials = append(socials, socialFromModel(account))
	}
	return socials
}

func socialFromModel(account model.UserSocialAccount) dto.UserSocialAccountDTO {
	return dto.UserSocialAccountDTO{
		ID:       account.ID,
		Provider: account.Provider,
	}
}
