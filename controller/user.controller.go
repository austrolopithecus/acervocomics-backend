package controller

import (
	"acervocomics/models"
	"acervocomics/util"
	"fmt"
)

func Me(id uint) (*models.User, error) {
	db := util.DB
	// Verifica se o ID é válido
	if id == 0 {
		return nil, fmt.Errorf("invaid id")
	}
	// Busca o usuário pelo ID
	user := models.User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	// Retorna o usuário
	return &user, nil
}

func Register(body *models.RegisterUserRequest) (*models.RegisterResponse, error) {
	db := util.DB

	// Verifica se o email já existe
	existingUser := models.User{}
	if err := db.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		return nil, nil
	}
	// Cria o usuário
	user := models.User{}
	user.Email = body.Email
	user.Nome = body.Nome
	user.Admin = false
	user.Senha = body.Password
	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// Gera o token
	token, err := util.MintToken(&user)
	if err != nil {
		return nil, err
	}
	// Retorna a resposta
	return &models.RegisterResponse{
		Token: token,
		CommonResponse: models.CommonResponse{
			Message: "Usuario criado com sucesso!",
			Succes:  true,
		},
	}, nil
}

func Login(body *models.LoginUserRequest) (*models.LoginResponse, error) {
	db := util.DB
	// Verifica se o email e senha estão corretos
	user := models.User{}
	if err := db.Where("email = ? AND senha = ?", body.Email, body.Password).First(&user).Error; err != nil {
		return nil, err
	}
	// Gera o token
	token, err := util.MintToken(&user)
	if err != nil {
		return nil, err
	}
	// Retorna a resposta
	return &models.LoginResponse{
		Token: token,
		CommonResponse: models.CommonResponse{
			Message: "Login realizado com sucesso!",
			Succes:  true,
		},
	}, nil
}
