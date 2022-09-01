package otten

import (
	"net/http"
	"otten/pkg"
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetTransaction(c echo.Context) error {

	var result Result
	var received string
	result.Status.Code = viper.GetString("message.errorCode")
	result.Status.Message = viper.GetString("message.errorText")
	result.Data.ReceivedBy = ""

	client := resty.New()

	url := viper.GetString("data.url")

	resp, err := client.R().
		EnableTrace().
		SetDoNotParseResponse(true).
		Get(url)

	if err != nil {
		return c.String(http.StatusBadRequest, string("error read URL"))
	}

	lookUp, err := pkg.LookUpHTML(resp.RawBody())
	if err != nil {
		return c.String(http.StatusInternalServerError, string("error lookup HTML"))
	}

	history := []History{}
	for i, val := range lookUp {
		tmp := History{}
		val = strings.Trim(val, "\t")
		if i%3 != 0 {
			if val != "" {
				dateStr := strings.Trim(string(lookUp[i-1]), "\t")
				dateValue, _ := time.Parse("02-01-2006 15:04", dateStr)
				createdAt := dateValue.Format("2006-01-02T15:04:05+07:00")
				formattedAt := dateValue.Format("02 January 2006, 15:04 WIB")
				tmp.Description = val
				tmp.CreatedAt = createdAt
				tmp.Formatted.CreatedAt = formattedAt
				history = append(history, tmp)
				found, _ := regexp.MatchString("DELIVERED TO", val)
				if found {
					trimStr := strings.TrimLeft(strings.TrimRight(val, "]"), "DELIVERED TO [")
					s := strings.Split(trimStr, "  |")
					received = s[0]
				}
			}
		}
	}

	if len(lookUp) > 0 {
		result.Status.Code = viper.GetString("message.successCode")
		result.Status.Message = viper.GetString("message.successText")
		result.Data.ReceivedBy = received
	}
	result.Data.History = history

	return c.JSON(http.StatusOK, result)
}
