package repositories

import (
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/models"
)

func GetTeams(offset int, recordPerPage int) ([]dto.TeamDTO, error) {
	//recebo aqui em baixo as equipas completas
	var teams []models.Team
	if err := database.DB.
		Offset(offset).
		Limit(recordPerPage).
		Order("id ASC").
		Preload("Coach").
		Find(&teams).Error; err != nil {
		return nil, &helper.CustomError{
			Type:    helper.ErrNotFound,
			Message: helper.ErrorResponse{
        Message: "Teams not found for the requested parameters.",
    },
		}
	}

	// Converte os dados das equipas para DTO, removendo informações desnecessárias
	var teamsDTO []dto.TeamDTO
	for _, team := range teams {
		teamsDTO = append(teamsDTO, dto.TeamDTO{
			ID:        team.ID,
			Name:      team.Name,
			Country:   team.Country,
			CoachID:   team.CoachID,
			Pts:       team.Pts,
			CreatedAt: team.CreatedAt.String(),
			Coach: dto.CoachDTO{ // Apenas os campos necessários do Coach
				Name:  team.Coach.Name,
				Email: team.Coach.Email,
			},
		})
	}

	return teamsDTO, nil
}

/* func GetUserByID(userID int) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return user, errors.New("notfound")
	}

	return user, nil
}

func UpdateUser(user *models.User) (*models.User, error) {

	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(user *models.User) error {
log.Println("user repositorie:", user)
	if err := database.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
*/
