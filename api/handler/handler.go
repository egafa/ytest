package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/egafa/Spr1Inc1/api/model"
)

func MetricHandler(m model.Metric) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := string(r.URL.Path)

		//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
		words := strings.Split(ss, "/")
		if len(words) < 5 {
			w.Write([]byte("Ошибка запроса"))
			w.Write([]byte(http.StatusText(400)))
			return
		}

		nameMetric := words[3]
		strVal := words[4]
		strType := words[2]
		if strType == "gauge" {
			f, err := strconv.ParseFloat(strVal, 64)
			if err != nil {
				m.SaveGaugeVal(nameMetric, 0)
			}
			m.SaveGaugeVal(nameMetric, f)
		}

		if strType == "counter" {
			i, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				m.SaveCounterVal(nameMetric, 0)
			}
			m.SaveCounterVal(nameMetric, i)
		}

		for idx, word := range words {

			w.Write([]byte(fmt.Sprintf("Word %d is: %s\n", idx, word)))
		}

	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ss := string(r.URL.Path)

	//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	words := strings.Split(ss, "/")
	if len(words) < 3 {
		w.Write([]byte("Ошибка запроса"))
		w.Write([]byte(http.StatusText(400)))
		return
	}
	for idx, word := range words {

		w.Write([]byte(fmt.Sprintf("Word %d is: %s\n", idx, word)))
	}
}
