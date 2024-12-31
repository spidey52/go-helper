package helper

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	Page      int64     `json:"page" form:"page"`
	Limit     int64     `json:"limit" form:"limit"`
	Skip      int64     `json:"skip" form:"skip"`
	StartDate time.Time `json:"start_date" form:"start_date"`
	EndDate   time.Time `json:"end_date" form:"end_date"`
}

func ParseDate(str string) time.Time {
	formats := []string{
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
	}

	for _, format := range formats {
		t, err := time.Parse(format, str)

		if err == nil {
			return t
		}
	}

	return time.Time{}
}

func GetPaginationData(c *gin.Context) PaginationData {
	page, _ := c.GetQuery("page")
	limit, _ := c.GetQuery("limit")
	start_date, _ := c.GetQuery("start_date")
	end_date, _ := c.GetQuery("end_date")

	paginationData := PaginationData{}

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)

	if err != nil {
		limitInt = 20
	}

	skip := pageInt * limitInt

	if start_date != "" {
		paginationData.StartDate = ParseDate(start_date)
	}

	if end_date != "" {
		paginationData.EndDate = ParseDate(end_date)
	}

	paginationData.Page = pageInt
	paginationData.Limit = limitInt
	paginationData.Skip = skip

	return paginationData
}
