package exception

import (
	"net/http"
	"sso-jwt/helper"
	"sso-jwt/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if authenticationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func authenticationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AuthenticationError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusUnauthorized,
			"error message": exception.Error,
		}).Warn("UNAUTHORIZED")

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusBadRequest,
			"error message": exception.Error,
		}).Warn("BAD REQUEST")

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusNotFound,
			"error message": exception.Error,
		}).Warn("NOT FOUND")

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusBadRequest,
			"error message": exception.Error,
		}).Warn("BAD REQUEST")

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	logger := helper.NewLogger()

	logger.WithFields(logrus.Fields{
		"status":        http.StatusInternalServerError,
		"error message": err,
	}).Error("INTERNAL SERVER ERROR")

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
