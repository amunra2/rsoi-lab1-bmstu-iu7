package handler

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"persserv/internal/dto"
	myerror "persserv/internal/error-my"
	"persserv/internal/usecase"
	mock_usecase "persserv/internal/usecase/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
	"go.uber.org/mock/gomock"
)

func TestHandler_createPerson(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockPerson, person dto.PersonCreate)

	tests := []struct {
		name                 string
		inputBody            string
		inputPerson          dto.PersonCreate
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"testName", "age":23, "address":"testAddress", "work":"testWork"}`,
			inputPerson: dto.PersonCreate{
				Name:    "testName",
				Age:     23,
				Address: "testAddress",
				Work:    "testWork",
			},
			mockBehavior: func(s *mock_usecase.MockPerson, person dto.PersonCreate) {
				s.EXPECT().Create(person).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: "",
		},
		{
			name:                 "BadRequest:no-name",
			inputBody:            `{"age":23, "address":"testAddress", "work":"testWork"}`,
			inputPerson:          dto.PersonCreate{},
			mockBehavior:         func(s *mock_usecase.MockPerson, person dto.PersonCreate) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"handler - persons - create: validation error","errors":{"Name":"required"}}`,
		},
		{
			name:                 "BadRequest:age-less-than-0",
			inputBody:            `{"name":"testName", "age":-1, "address":"testAddress", "work":"testWork"}`,
			inputPerson:          dto.PersonCreate{},
			mockBehavior:         func(s *mock_usecase.MockPerson, person dto.PersonCreate) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"handler - persons - create: validation error","errors":{"Age":"gte"}}`,
		},
	}

	logrus.SetOutput(ioutil.Discard)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			persMockUsecase := mock_usecase.NewMockPerson(c)
			test.mockBehavior(persMockUsecase, test.inputPerson)

			usecases := &usecase.UseCase{Person: persMockUsecase}
			handler := Handler{useCases: usecases}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/api/v1/persons", handler.createPerson)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/persons", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updatePerson(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockPerson, person dto.PersonUpdate)

	tests := []struct {
		name                 string
		inputBody            string
		inputPerson          dto.PersonUpdate
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"updatedName"}`,
			inputPerson: dto.PersonUpdate{
				Name: "updatedName",
			},
			mockBehavior: func(s *mock_usecase.MockPerson, person dto.PersonUpdate) {
				s.EXPECT().Update(1, person).Return(dto.Person{
					Id:      1,
					Name:    "updatedName",
					Age:     23,
					Address: "testAddress",
					Work:    "testWork",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"updatedName","age":23,"address":"testAddress","work":"testWork"}`,
		},
		{
			name:                 "BadRequest:age-less-than-0",
			inputBody:            `{"age":-5}`,
			inputPerson:          dto.PersonUpdate{},
			mockBehavior:         func(s *mock_usecase.MockPerson, person dto.PersonUpdate) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"handler - persons - create: validation error","errors":{"Age":"gte"}}`,
		},
		{
			name:        "BadRequest:empty-update-struct",
			inputBody:   `{}`,
			inputPerson: dto.PersonUpdate{},
			mockBehavior: func(s *mock_usecase.MockPerson, person dto.PersonUpdate) {
				s.EXPECT().Update(1, person).Return(dto.Person{}, myerror.NewError(updatePersonFuncName, myerror.UpdateStructureIsEmpty))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"handler - persons - update: update structure has no values"}`,
		},
	}

	logrus.SetOutput(ioutil.Discard)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			persMockUsecase := mock_usecase.NewMockPerson(c)
			test.mockBehavior(persMockUsecase, test.inputPerson)

			usecases := &usecase.UseCase{Person: persMockUsecase}
			handler := Handler{useCases: usecases}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.PATCH("/api/v1/persons/:id", handler.updatePerson)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/api/v1/persons/1", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getByIdPerson(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockPerson)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_usecase.MockPerson) {
				s.EXPECT().GetById(1).Return(dto.Person{
					Id:      1,
					Name:    "testName",
					Age:     23,
					Address: "testAddress",
					Work:    "testWork",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"testName","age":23,"address":"testAddress","work":"testWork"}`,
		},
		{
			name: "NotFound",
			mockBehavior: func(s *mock_usecase.MockPerson) {
				s.EXPECT().GetById(1).Return(dto.Person{}, myerror.NewError("test - repo", myerror.NotFound))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"test - repo: content not found"}`,
		},
	}

	logrus.SetOutput(ioutil.Discard)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			persMockUsecase := mock_usecase.NewMockPerson(c)
			test.mockBehavior(persMockUsecase)

			usecases := &usecase.UseCase{Person: persMockUsecase}
			handler := Handler{useCases: usecases}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/api/v1/persons/:id", handler.getByIdPerson)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/persons/1", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getAllPerson(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockPerson)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_usecase.MockPerson) {
				s.EXPECT().GetAll().Return([]dto.Person{{
					Id:      1,
					Name:    "testName",
					Age:     23,
					Address: "testAddress",
					Work:    "testWork",
				}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"name":"testName","age":23,"address":"testAddress","work":"testWork"}]`,
		},
	}

	logrus.SetOutput(ioutil.Discard)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			persMockUsecase := mock_usecase.NewMockPerson(c)
			test.mockBehavior(persMockUsecase)

			usecases := &usecase.UseCase{Person: persMockUsecase}
			handler := Handler{useCases: usecases}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/api/v1/persons/", handler.getAllPersons)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/persons/", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deletePerson(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockPerson)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_usecase.MockPerson) {
				s.EXPECT().Delete(1).Return(nil)
			},
			expectedStatusCode:   204,
			expectedResponseBody: ``,
		},
		{
			name: "NotFound",
			mockBehavior: func(s *mock_usecase.MockPerson) {
				s.EXPECT().Delete(1).Return(myerror.NewError("test - repo", myerror.NotFound))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"test - repo: content not found"}`,
		},
	}

	logrus.SetOutput(ioutil.Discard)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			persMockUsecase := mock_usecase.NewMockPerson(c)
			test.mockBehavior(persMockUsecase)

			usecases := &usecase.UseCase{Person: persMockUsecase}
			handler := Handler{useCases: usecases}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.DELETE("/api/v1/persons/:id", handler.deletePerson)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/persons/1", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
