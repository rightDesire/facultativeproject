// Package ouser provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package ouser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// AuthTokenDTO defines model for AuthTokenDTO.
type AuthTokenDTO struct {
	Token *string `json:"token,omitempty"`
}

// UserDTO defines model for UserDTO.
type UserDTO struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Email     *string    `json:"email,omitempty"`
	FullName  *string    `json:"fullName,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Uuid      *string    `json:"uuid,omitempty"`
}

// UserLoginRequest defines model for UserLoginRequest.
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegisterRequest defines model for UserRegisterRequest.
type UserRegisterRequest struct {
	Email    string  `json:"email"`
	FullName *string `json:"fullName,omitempty"`
	Password string  `json:"password"`
}

// UserUpdateRequest defines model for UserUpdateRequest.
type UserUpdateRequest struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UserId defines model for userId.
type UserId = string

// PostUsersLoginJSONRequestBody defines body for PostUsersLogin for application/json ContentType.
type PostUsersLoginJSONRequestBody = UserLoginRequest

// PostUsersRegisterJSONRequestBody defines body for PostUsersRegister for application/json ContentType.
type PostUsersRegisterJSONRequestBody = UserRegisterRequest

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody = UserUpdateRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Аутентификация пользователя
	// (POST /users/login)
	PostUsersLogin(ctx echo.Context) error
	// Регистрация нового пользователя
	// (POST /users/register)
	PostUsersRegister(ctx echo.Context) error
	// Получить профиль пользователя
	// (GET /users/{userId})
	GetUsersUserId(ctx echo.Context, userId UserId) error
	// Обновление профиля пользователя
	// (PUT /users/{userId})
	PutUsersUserId(ctx echo.Context, userId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostUsersLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersLogin(ctx)
	return err
}

// PostUsersRegister converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersRegister(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersRegister(ctx)
	return err
}

// GetUsersUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersUserId(ctx, userId)
	return err
}

// PutUsersUserId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutUsersUserId(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/users/login", wrapper.PostUsersLogin)
	router.POST(baseURL+"/users/register", wrapper.PostUsersRegister)
	router.GET(baseURL+"/users/:userId", wrapper.GetUsersUserId)
	router.PUT(baseURL+"/users/:userId", wrapper.PutUsersUserId)

}

type PostUsersLoginRequestObject struct {
	Body *PostUsersLoginJSONRequestBody
}

type PostUsersLoginResponseObject interface {
	VisitPostUsersLoginResponse(w http.ResponseWriter) error
}

type PostUsersLogin200JSONResponse AuthTokenDTO

func (response PostUsersLogin200JSONResponse) VisitPostUsersLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersLogin401Response struct {
}

func (response PostUsersLogin401Response) VisitPostUsersLoginResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PostUsersRegisterRequestObject struct {
	Body *PostUsersRegisterJSONRequestBody
}

type PostUsersRegisterResponseObject interface {
	VisitPostUsersRegisterResponse(w http.ResponseWriter) error
}

type PostUsersRegister201JSONResponse UserDTO

func (response PostUsersRegister201JSONResponse) VisitPostUsersRegisterResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersRegister400Response struct {
}

func (response PostUsersRegister400Response) VisitPostUsersRegisterResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetUsersUserIdRequestObject struct {
	UserId UserId `json:"userId"`
}

type GetUsersUserIdResponseObject interface {
	VisitGetUsersUserIdResponse(w http.ResponseWriter) error
}

type GetUsersUserId200JSONResponse UserDTO

func (response GetUsersUserId200JSONResponse) VisitGetUsersUserIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersUserId404Response struct {
}

func (response GetUsersUserId404Response) VisitGetUsersUserIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutUsersUserIdRequestObject struct {
	UserId UserId `json:"userId"`
	Body   *PutUsersUserIdJSONRequestBody
}

type PutUsersUserIdResponseObject interface {
	VisitPutUsersUserIdResponse(w http.ResponseWriter) error
}

type PutUsersUserId200Response struct {
}

func (response PutUsersUserId200Response) VisitPutUsersUserIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutUsersUserId400Response struct {
}

func (response PutUsersUserId400Response) VisitPutUsersUserIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Аутентификация пользователя
	// (POST /users/login)
	PostUsersLogin(ctx context.Context, request PostUsersLoginRequestObject) (PostUsersLoginResponseObject, error)
	// Регистрация нового пользователя
	// (POST /users/register)
	PostUsersRegister(ctx context.Context, request PostUsersRegisterRequestObject) (PostUsersRegisterResponseObject, error)
	// Получить профиль пользователя
	// (GET /users/{userId})
	GetUsersUserId(ctx context.Context, request GetUsersUserIdRequestObject) (GetUsersUserIdResponseObject, error)
	// Обновление профиля пользователя
	// (PUT /users/{userId})
	PutUsersUserId(ctx context.Context, request PutUsersUserIdRequestObject) (PutUsersUserIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostUsersLogin operation middleware
func (sh *strictHandler) PostUsersLogin(ctx echo.Context) error {
	var request PostUsersLoginRequestObject

	var body PostUsersLoginJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsersLogin(ctx.Request().Context(), request.(PostUsersLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsersLogin")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersLoginResponseObject); ok {
		return validResponse.VisitPostUsersLoginResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsersRegister operation middleware
func (sh *strictHandler) PostUsersRegister(ctx echo.Context) error {
	var request PostUsersRegisterRequestObject

	var body PostUsersRegisterJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsersRegister(ctx.Request().Context(), request.(PostUsersRegisterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsersRegister")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersRegisterResponseObject); ok {
		return validResponse.VisitPostUsersRegisterResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersUserId operation middleware
func (sh *strictHandler) GetUsersUserId(ctx echo.Context, userId UserId) error {
	var request GetUsersUserIdRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersUserId(ctx.Request().Context(), request.(GetUsersUserIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersUserId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersUserIdResponseObject); ok {
		return validResponse.VisitGetUsersUserIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutUsersUserId operation middleware
func (sh *strictHandler) PutUsersUserId(ctx echo.Context, userId UserId) error {
	var request PutUsersUserIdRequestObject

	request.UserId = userId

	var body PutUsersUserIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutUsersUserId(ctx.Request().Context(), request.(PutUsersUserIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutUsersUserId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutUsersUserIdResponseObject); ok {
		return validResponse.VisitPutUsersUserIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
