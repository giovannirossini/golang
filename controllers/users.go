package controllers

import (
	"net/http"
	"strconv"

	"github.com/giovannirossini/golang/models"
	"github.com/labstack/echo"
)

// Home function to list and render
//  Render homepage and listing all users
func Home(c echo.Context) error {
	var users []models.Users

	if err := models.UsersModel.Find().All(&users); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"mensagem": "Error when trying to get data.",
		})
	}
	date := map[string]interface{}{
		"titulo": "Users list",
		"users":  users,
	}
	return c.Render(http.StatusOK, "index.html", date)
}

// Add users function redirect
// Render new user page
func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

// Edit function for user
// Homepage button that return the data information
func Edit(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	var user models.Users

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found.",
		})
	}
	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "User not found.",
		})
	}

	var date = map[string]interface{}{
		"user": user,
	}

	return c.Render(http.StatusOK, "edit.html", date)
}

// Post is the POST function
// Receives via ajax function the data to insert on database
func Post(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user models.Users

	user.Name = name
	user.Email = email

	if name != "" && email != "" {
		if _, err := models.UsersModel.Insert(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Could not add user. Try again.",
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Fields needs to be filled.",
	})
}

// Delete users function
// Receives via ajax function the data to delete in homepage
func Delete(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensage": "User not found.",
		})
	}
	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensage": "User cannot be deleted.",
		})
	}
	return c.JSON(http.StatusAccepted, map[string]string{
		"mensage": "User successfully deleted.",
	})

}

// Put to edit users
// Receives via ajax function the data to change
func Put(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user = models.Users{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	result := models.UsersModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "User does not exist.",
		})
	}

	if err := result.Update(user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error trying to update registry.",
		})
	}

	return c.JSON(http.StatusAccepted, user)
}
