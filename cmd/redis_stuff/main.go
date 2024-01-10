package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type SearchResult struct {
	estimation_method              string
	origin_id                      string
	origin_type                    string `example:"kma"`
	origin_country                 string
	dest_id                        string
	dest_type                      string
	dest_country                   string
	load_source                    string
	equipment_type                 string
	search_lane_miles              int32
	search_lane_miles_low          int32
	search_lane_miles_high         int32
	search_min_partners            int8
	search_min_loads               int8
	search_max_mile_deviation      float32
	search_min_mile_deviation      int32
	search_daysback_steps          float32
	search_daysback_stop           int32
	search_origin_radius_steps     float32
	search_origin_radius_start     int32
	search_origin_radius_stop      int32
	search_dest_radius_steps       float32
	search_dest_radius_start       int32
	search_dest_radius_stop        int32
	agg_method                     string
	agg_confidence_level           float32
	agg_datewindow_decay           float32
	agg_dayback_decay              float32
	agg_origin_radius_decay        float32
	agg_dest_radius_decay          float32
	agg_mileage_decay              float32
	agg_datewindow_weight          int8
	agg_dayback_weight             int8
	agg_origin_radius_weight       int8
	agg_mileage_weight             int8
	agg_max_partner_weight         float32
	iterations_needed              int16
	time_needed                    int16
	daysback_needed                int16
	origin_radius_expansion        int32
	dest_radius_expansion          int32
	n_loads_found                  int16
	n_partners_found               int16
	miles_est_mean                 int32
	rpm_total_est_mean             float32
	rpm_total_est_sd               float32
	rpm_total_est_median           float32
	rpm_total_est_low              float32
	rpm_total_est_high             float32
	rate_total_est_mean            int32
	rate_total_est_sd              int32
	rate_total_est_median          int32
	rate_total_est_low             int32
	rate_total_est_high            int32
	origin_radius_confidence_score int8
	dest_radius_confidence_score   int8
	total_confidence_score         int8
	data_timestamp                 int64
	generated_at                   string
}

func ftSearch(rdb *redis.Client, indexName string, query string, resultsChannel chan<- SearchResult) {
	defer close(resultsChannel)

	results, err := rdb.Do(ctx, "FT.SEARCH", indexName, query).Slice()
	if err != nil {
		log.Panicln("Error: ", err)
		return
	}

	for _, result := range results {
		// Take the result and convert it to a SearchResult struct
		// Then send it to the resultsChannel
		fmt.Println(result)
		searchResult := SearchResult{}
		searchResult.estimation_method = result.([]interface{})[1].(string)
		searchResult.origin_id = result.([]interface{})[2].(string)
		searchResult.origin_type = result.([]interface{})[3].(string)
		searchResult.origin_country = result.([]interface{})[4].(string)
		searchResult.dest_id = result.([]interface{})[5].(string)
		searchResult.dest_type = result.([]interface{})[6].(string)
		searchResult.dest_country = result.([]interface{})[7].(string)
		searchResult.load_source = result.([]interface{})[8].(string)
		searchResult.equipment_type = result.([]interface{})[9].(string)
		searchResult.search_lane_miles = result.([]interface{})[10].(int32)
		searchResult.search_lane_miles_low = result.([]interface{})[11].(int32)
		searchResult.search_lane_miles_high = result.([]interface{})[12].(int32)
		searchResult.search_min_partners = result.([]interface{})[13].(int8)
		searchResult.search_min_loads = result.([]interface{})[14].(int8)
		searchResult.search_max_mile_deviation = result.([]interface{})[15].(float32)
		searchResult.search_min_mile_deviation = result.([]interface{})[16].(int32)
		searchResult.search_daysback_steps = result.([]interface{})[17].(float32)
		searchResult.search_daysback_stop = result.([]interface{})[18].(int32)
		searchResult.search_origin_radius_steps = result.([]interface{})[19].(float32)
		searchResult.search_origin_radius_start = result.([]interface{})[20].(int32)
		searchResult.search_origin_radius_stop = result.([]interface{})[21].(int32)
		searchResult.search_dest_radius_steps = result.([]interface{})[22].(float32)
		searchResult.search_dest_radius_start = result.([]interface{})[23].(int32)
		searchResult.search_dest_radius_stop = result.([]interface{})[24].(int32)
		searchResult.agg_method = result.([]interface{})[25].(string)
		searchResult.agg_confidence_level = result.([]interface{})[26].(float32)
		searchResult.agg_datewindow_decay = result.([]interface{})[27].(float32)
		searchResult.agg_dayback_decay = result.([]interface{})[28].(float32)
		searchResult.agg_origin_radius_decay = result.([]interface{})[29].(float32)
		searchResult.agg_dest_radius_decay = result.([]interface{})[30].(float32)
		searchResult.agg_mileage_decay = result.([]interface{})[31].(float32)
		searchResult.agg_datewindow_weight = result.([]interface{})[32].(int8)
		searchResult.agg_dayback_weight = result.([]interface{})[33].(int8)
		searchResult.agg_origin_radius_weight = result.([]interface{})[34].(int8)
		searchResult.agg_mileage_weight = result.([]interface{})[35].(int8)
		searchResult.agg_max_partner_weight = result.([]interface{})[36].(float32)
		searchResult.iterations_needed = result.([]interface{})[37].(int16)
		searchResult.time_needed = result.([]interface{})[38].(int16)
		searchResult.daysback_needed = result.([]interface{})[39].(int16)
		searchResult.origin_radius_expansion = result.([]interface{})[40].(int32)
		searchResult.dest_radius_expansion = result.([]interface{})[41].(int32)
		searchResult.n_loads_found = result.([]interface{})[42].(int16)
		searchResult.n_partners_found = result.([]interface{})[43].(int16)
		searchResult.miles_est_mean = result.([]interface{})[44].(int32)
		searchResult.rpm_total_est_mean = result.([]interface{})[45].(float32)
		searchResult.rpm_total_est_sd = result.([]interface{})[46].(float32)
		searchResult.rpm_total_est_median = result.([]interface{})[47].(float32)
		searchResult.rpm_total_est_low = result.([]interface{})[48].(float32)
		searchResult.rpm_total_est_high = result.([]interface{})[49].(float32)
		searchResult.rate_total_est_mean = result.([]interface{})[50].(int32)
		searchResult.rate_total_est_sd = result.([]interface{})[51].(int32)
		searchResult.rate_total_est_median = result.([]interface{})[52].(int32)
		searchResult.rate_total_est_low = result.([]interface{})[53].(int32)
		searchResult.rate_total_est_high = result.([]interface{})[54].(int32)
		searchResult.origin_radius_confidence_score = result.([]interface{})[55].(int8)
		searchResult.dest_radius_confidence_score = result.([]interface{})[56].(int8)
		searchResult.total_confidence_score = result.([]interface{})[57].(int8)
		searchResult.data_timestamp = result.([]interface{})[58].(int64)
		searchResult.generated_at = result.([]interface{})[59].(string)

		resultsChannel <- searchResult
	}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "34.173.154.159:6379", // Redis server address
		Password: "1234test",            // no password set
		DB:       0,                     // use default DB
	})

	fmt.Println(rdb.Ping(ctx))

	resultsChan := make(chan SearchResult)

	go ftSearch(rdb, "idx:kma_to_kma_spot_rates", "@origin_id:TN_CHA", resultsChan)

	for result := range resultsChan {
		fmt.Println("Recieved: ", result)
	}
}
