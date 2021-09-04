package handler

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jeffprestes/test-bbolt/lib/contx"
	"github.com/jeffprestes/test-bbolt/repo"
)

func SaveUser(ctx *contx.Context) {
	newUser, err := repo.SaveUser("admin", "ze_"+strconv.Itoa(rand.Intn(99))+"@teste.com", "Ze da Silva", rand.Intn(99))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not save data")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func GetAllUsers(ctx *contx.Context) {
	users, err := repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not retrieve data")
		return
	}
	if len(users) < 1 {
		ctx.JSON(http.StatusNotFound, "no records in the database")
		return
	}
	ctx.JSON(http.StatusOK, users)
}
