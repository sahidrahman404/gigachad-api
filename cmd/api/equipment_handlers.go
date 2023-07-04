package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) createEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateEquipmentParams

	err := request.DecodeJSON(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if params.Validate(v); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	equipment := types.NewEquipmentFromParams(params)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Equipment.Insert(ctx, equipment)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		equipment,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	equipment, err := app.storage.Equipment.GetAll(ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, equipment)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := app.readParams("equipmentID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	equipmentParams := types.UpdateEquipmentParams{}
	err = request.DecodeJSON(w, r, &equipmentParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	equipment, err := app.storage.Equipment.Get(ctx, equipmentID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	types.UpdateEquipmentFromParams(&equipmentParams, equipment)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Equipment.Update(ctx, equipment)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, equipment)
	if err != nil {
		app.serverError(w, r, err)
	}
}
