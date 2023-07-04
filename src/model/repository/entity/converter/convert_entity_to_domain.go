package converter

import (
	"github.com/Fccrelick/crudGO/src/model"
	"github.com/Fccrelick/crudGO/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())
	return domain
}
