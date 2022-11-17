package pkg

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Fermekoo/game-store/utils"
)

type VIPPayment struct {
	VIPPayment ApiGameInterface
	config     utils.Config
}

func NewVIPPayment(config utils.Config) *VIPPayment {
	return &VIPPayment{
		config: config,
	}
}

func (vip VIPPayment) callApi(path string, payload *bytes.Buffer) *http.Response {
	var client = &http.Client{}
	url := vip.config.VIPBaseURL + path

	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func (vip VIPPayment) generateSign() string {
	hash := md5.New()
	hash.Write([]byte(vip.config.VIPApiID + vip.config.VIPApiKey))
	sign := hex.EncodeToString(hash.Sum(nil))
	return sign
}

func (vip VIPPayment) Profile() (ProfileResponse, error) {

	var profileResponse ProfileResponse
	var param = url.Values{}
	sign := vip.generateSign()
	param.Set("key", vip.config.VIPApiKey)
	param.Set("sign", sign)

	payload := bytes.NewBufferString(param.Encode())

	response := vip.callApi("profile", payload)
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&profileResponse)

	return profileResponse, err
}

func (vip VIPPayment) Order(payload OrderCall) (OrderResponse, error) {
	var orderResponse OrderResponse
	var param = url.Values{}
	sign := vip.generateSign()

	param.Set("key", vip.config.VIPApiKey)
	param.Set("sign", sign)
	param.Set("type", "order")
	param.Set("service", payload.ServiceCode)
	param.Set("data_no", fmt.Sprint(payload.AccountID))
	param.Set("data_zone", payload.AccountZone)

	request_payload := bytes.NewBufferString(param.Encode())
	response := vip.callApi("game-feature", request_payload)
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&orderResponse)

	return orderResponse, err
}

func (vip VIPPayment) ListService(filter FilterListService) (ServiceResponse, error) {
	var listService ServiceResponse
	param := url.Values{}
	sign := vip.generateSign()

	param.Set("key", vip.config.VIPApiKey)
	param.Set("sign", sign)
	param.Set("type", "services")

	if filter.FilterType != "" && filter.FilterValue != "" {
		param.Set("filter_type", filter.FilterType)
		param.Set("filter_value", filter.FilterValue)
	}

	request_payload := bytes.NewBufferString(param.Encode())
	response := vip.callApi("game-feature", request_payload)
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&listService)

	return listService, err
}

func (vip VIPPayment) Game() ([]GameResponse, error) {
	var filter FilterListService
	listServices, _ := vip.ListService(filter)

	data := listServices.Data

	buckets := map[string]GameResponse{}
	outputs := []GameResponse{}
	for _, v := range data {
		out, exists := buckets[v.Game]
		if !exists {
			out = GameResponse{
				Name: v.Game,
			}
			buckets[v.Game] = out
			outputs = append(outputs, out)
		}
	}
	return outputs, nil
}
