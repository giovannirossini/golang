package controllers

import (
	"net/http"
	"strconv"

	"github.com/giovannirossini/curso/models"
	"github.com/labstack/echo"
)

// Home is the homepage
func Home(c echo.Context) error {
	var users []models.Users

	if err := models.UsersModel.Find().All(&users); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"mensagem": "Erro ao tentar recuperar os dados!",
		})
	}
	date := map[string]interface{}{
		"titulo": "Listagem De Usuários",
		"users":  users,
	}
	return c.Render(http.StatusOK, "index.html", date)
}

// Add users function redirect
func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

// Edit function for user
func Edit(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	var user models.Users

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não foi encontrado!",
		})
	}
	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Não foi possível encontrar o usuario.",
		})
	}

	var date = map[string]interface{}{
		"user": user,
	}

	return c.Render(http.StatusOK, "edit.html", date)
}

// Post is the POST function
func Post(c echo.Context) error {
	nome := c.FormValue("name")
	email := c.FormValue("email")

	var user models.Users

	user.Nome = nome
	user.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsersModel.Insert(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Não foi possível adicionar o registro no banco! Tente novamente.",
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Os campos precisam ser preenchidos!",
	})
}

// Delete users function
func Delete(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensage": "Usuário não encontrado!",
		})
	}
	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensage": "Usuário não pôde ser deletado!",
		})
	}
	return c.JSON(http.StatusAccepted, map[string]string{
		"mensage": "Usuário deletado com sucesso!",
	})

}

// Put to edit users
func Put(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("name")
	email := c.FormValue("email")

	var user = models.Users{
		ID:    userID,
		Nome:  nome,
		Email: email,
	}

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não existe!",
		})
	}

	if err := result.Update(user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar atualizar o registro!",
		})
	}

	return c.JSON(http.StatusAccepted, user)
}
