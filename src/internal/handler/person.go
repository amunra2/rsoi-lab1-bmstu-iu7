package handler

import (
	"errors"
	"fmt"
	"net/http"
	"persserv/internal/dto"
	myerror "persserv/internal/error-my"
	"persserv/internal/handler/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	getAllPersonFuncName  = "handler - persons - get all"
	getByIdPersonFuncName = "handler - persons - get by id"
	createPersonFuncName  = "handler - persons - create"
	updatePersonFuncName  = "handler - persons - update"
	deletePersonFuncName  = "handler - persons - delete"
)

// @Summary Get all persons
// @Tags persons
// @Description get all persons
// @ID getall-person
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Person
// @Router /api/v1/persons [get]
func (h *Handler) getAllPersons(ctx *gin.Context) {
	persons, err := h.useCases.Person.GetAll()

	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.GetMessage())
		return
	}

	ctx.JSON(http.StatusOK, persons)
}

// @Summary Get person by id
// @Tags persons
// @Description get person by id
// @ID getbyid-person
// @Param id path int true "Person Id"
// @Accept json
// @Produce json
// @Success 200 {object} dto.Person
// @Failure 404 {object} responses.ErrorResponse
// @Router /api/v1/persons/{id} [get]
func (h *Handler) getByIdPerson(ctx *gin.Context) {
	personId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.NewErrorResponse(ctx,
			http.StatusBadRequest,
			myerror.NewError(getByIdPersonFuncName, err).GetMessage())
		return
	}

	person, errGet := h.useCases.Person.GetById(personId)
	if errGet != nil {
		if errors.Is(errGet.Err, myerror.NotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, errGet.GetMessage())
		} else {
			responses.NewErrorResponse(ctx, http.StatusInternalServerError, errGet.GetMessage())
		}

		return
	}

	ctx.JSON(http.StatusOK, person)
}

// @Summary Create person
// @Tags persons
// @Description create person
// @ID create-person
// @Accept  json
// @Produce  json
// @Param input body dto.PersonCreate true "person info"
// @Success 201
// @Header 201 {string} Location "Path to new person"
// @Failure 400 {object} responses.ValidationErrorResponse
// @Router /api/v1/persons [post]
func (h *Handler) createPerson(ctx *gin.Context) {
	var personInput dto.PersonCreate
	if err := ctx.BindJSON(&personInput); err != nil {
		responses.NewValidationErrorResponse(ctx,
			myerror.NewError(createPersonFuncName, myerror.ValidationError).GetMessage(),
			err,
		)
		return
	}

	id, err := h.useCases.Person.Create(personInput)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.GetMessage())
		return
	}

	ctx.Header("Location", fmt.Sprintf(`/api/v1/persons/%d`, id))
	ctx.Status(http.StatusCreated)
}

// @Summary Update person
// @Tags persons
// @Description update person
// @ID update-person
// @Accept  json
// @Produce  json
// @Param id path int true "Person Id"
// @Param input body dto.PersonUpdate true "person update info"
// @Success 200 {object} dto.Person
// @Failure 400 {object} responses.ValidationErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /api/v1/persons/{id} [patch]
func (h *Handler) updatePerson(ctx *gin.Context) {
	personId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.NewErrorResponse(ctx,
			http.StatusBadRequest,
			myerror.NewError(updatePersonFuncName, err).GetMessage())
		return
	}

	var personInput dto.PersonUpdate
	if err := ctx.BindJSON(&personInput); err != nil {
		responses.NewValidationErrorResponse(ctx,
			myerror.NewError(createPersonFuncName, myerror.ValidationError).GetMessage(),
			err,
		)
		return
	}

	person, errUpd := h.useCases.Person.Update(personId, personInput)
	if errUpd != nil {
		if errors.Is(errUpd.Err, myerror.NotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, errUpd.GetMessage())
		} else if errors.Is(errUpd.Err, myerror.UpdateStructureIsEmpty) {
			responses.NewErrorResponse(ctx, http.StatusBadRequest, errUpd.GetMessage())
		} else {
			responses.NewErrorResponse(ctx, http.StatusInternalServerError, errUpd.GetMessage())
		}

		return
	}

	ctx.JSON(http.StatusOK, person)
}

// @Summary Delete person by id
// @Tags persons
// @Description delete person by id
// @ID delete-person
// @Param id path int true "Person Id"
// @Accept json
// @Produce json
// @Success 204
// @Failure 404 {object} responses.ErrorResponse
// @Router /api/v1/persons/{id} [delete]
func (h *Handler) deletePerson(ctx *gin.Context) {
	personId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.NewErrorResponse(ctx,
			http.StatusBadRequest,
			myerror.NewError(updatePersonFuncName, err).GetMessage())
		return
	}

	errDel := h.useCases.Person.Delete(personId)
	if errDel != nil {
		if errors.Is(errDel.Err, myerror.NotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, errDel.GetMessage())
		} else {
			responses.NewErrorResponse(ctx, http.StatusInternalServerError, errDel.GetMessage())
		}

		return
	}

	ctx.Status(http.StatusNoContent)
}
