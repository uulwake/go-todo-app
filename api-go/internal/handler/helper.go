package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/internal/constant"
	"todo-app/internal/libs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func PreprocessedRequest(ctx echo.Context, validate *validator.Validate, body any) error {
	
	if body != nil {
		err := json.NewDecoder(ctx.Request().Body).Decode(&body)
		if err != nil {
			return libs.DefaultInternalServerError(ctx, err)
		}
		
		err = validate.Struct(body)
		if err != nil {
			return libs.DefaultInternalServerError(ctx, err)
		}
	}

	return nil
}

func GetUserIdFromContext(c echo.Context) (int64, error) {
	userId := c.Get("user_id")

	switch v := userId.(type) {
	case int64:
		return v, nil
	default:
		return 0, libs.DefaultInternalServerError(c, fmt.Errorf("invalid userID %v", userId))
	}
}

func GetIdFromPathParam(c echo.Context, key string) (int64, error) {
	idParam := c.Param(key)
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, libs.CustomError{
			HTTPCode: http.StatusBadRequest,
			Message: fmt.Sprintf("invalid %s ID",key),
		}
	}

	return id, nil
}

func ConvertCommonQueryParam(c echo.Context) (libs.QueryParam, error) {
	pageSizeStr := c.QueryParam("page_size")
	pageNumberStr := c.QueryParam("page_number")

	var queryParam libs.QueryParam
	var pageSize, pageNumber int
	var err error

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return queryParam, libs.CustomError{
				HTTPCode: http.StatusBadRequest,
				Message: "invalid page size query",
			}
		}

		if pageSize < constant.PAGE_SIZE_MIN || pageSize > constant.PAGE_SIZE_MAX {
			pageSize = constant.PAGE_SIZE_DEFAULT
		}
	} else {
		pageSize = constant.PAGE_SIZE_DEFAULT
	} 

	if pageNumberStr != "" {
		pageNumber, err = strconv.Atoi(pageNumberStr)
		if err != nil {
			return queryParam, libs.CustomError{
				HTTPCode: http.StatusBadRequest,
				Message: "invalid page number query",
			}
		}

		if pageNumber < constant.PAGE_NUMBER_MIN || pageNumber > constant.PAGE_NUMBER_MAX {
			pageNumber = constant.PAGE_NUMBER_DEFAULT
		}
	} else {
		pageNumber = constant.PAGE_NUMBER_DEFAULT
	}

	queryParam.PageNumber = pageNumber
	queryParam.PageSize = pageSize

	return queryParam, nil
}