package shortener

import (
	"encoding/json"
	"net/http"
	"sort"
	constant "url_shortener_server/constants"
)

func (s *Shortener) Metrics(w http.ResponseWriter, r *http.Request) {
	type domainStat struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
	}
	var stats []domainStat
	s.mu.RLock()
	for domain, count := range s.stats {
		stats = append(stats, domainStat{domain, count})
	}
	s.mu.RUnlock()

	if len(stats) == 0 {
		http.Error(w, constant.NoMetricCount, http.StatusOK)
		return
	}

	sort.Slice(stats, func(first, second int) bool {
		return stats[first].Count > stats[second].Count
	})

	if len(stats) > 3 {
		stats = stats[:3]
	}

	json.NewEncoder(w).Encode(stats)
}
