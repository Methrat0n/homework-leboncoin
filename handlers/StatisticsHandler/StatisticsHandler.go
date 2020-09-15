package StatisticsHandler

import (
	"leboncoin/model"
	"strconv"

	routing "github.com/qiangxue/fasthttp-routing"
)

type StatisticsHandler interface {
	GetMostUsed(request *routing.Context) error
	SaveStats(request *routing.Context) error
}

type defaultStatisticsHandler struct {
	statistics model.Statistics
}

func New(statistics model.Statistics) StatisticsHandler {
	return defaultStatisticsHandler{statistics}
}

func (handler defaultStatisticsHandler) SaveStats(request *routing.Context) error {
	handler.statistics.Stats[request.Request.URI().String()]++
	request.SetStatusCode(200)
	return nil
}
func (handler defaultStatisticsHandler) GetMostUsed(request *routing.Context) error {
	var mostUsedRequest string
	maxUse := 0
	for key, element := range handler.statistics.Stats {
		if element > maxUse {
			maxUse = element
			mostUsedRequest = key
		}
	}

	if mostUsedRequest == "" {
		request.SetStatusCode(400)
		request.SetBody([]byte("No request made"))
		return nil
	}

	request.SetStatusCode(200)
	request.SetContentType("application/json")
	request.SetBody([]byte("{ mostUsed: \"" + mostUsedRequest + "\", used: \"" + strconv.Itoa(maxUse) + "\"}"))
	return nil
}
