package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/budiariyanto/cqrs-dummy/common"
	"github.com/budiariyanto/cqrs-dummy/model"
	"github.com/budiariyanto/cqrs-dummy/payloads"
	"github.com/budiariyanto/cqrs-dummy/repo"
	"github.com/budiariyanto/cqrs-dummy/service"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:32772"},
		GroupID:   "group1",
		Topic:     "donation-create",
		ReadBackoffMin: 50*time.Millisecond,
		ReadBackoffMax: 500*time.Millisecond,
		MinBytes:  10e1, // 1KB
		MaxBytes:  10e6, // 10MB
	})

	fmt.Println("Starting consumer donation...")

	dbDonation := common.NewDB("katresnan")
	donationRepo := repo.NewDonationRepo(dbDonation)
	donationService := service.NewDonationService(donationRepo)
	
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Read message from donation consumer:", string(m.Key), string(m.Value))
		var donation model.Donation
		err = json.Unmarshal(m.Value, &donation)
		if err != nil {
			fmt.Println(err)
			break
		}

		err = donationService.Create(donation)
		if err != nil {
			fmt.Println(err)
			break
		}

		sum, err := donationService.Sum(donation.CampaignID)
		if err != nil {
			fmt.Println(err)
			break
		}

		payload := payloads.CampaignDetailPayload {
			ID:            donation.CampaignID,
			TotalDonation: sum,
		}

		buf, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err)
			break
		}
		
		client := http.DefaultClient
		req, _ := http.NewRequest("PUT", "http://localhost:1323/campaign-detail/update-total-donation", bytes.NewBuffer(buf))
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			break
		}
		defer res.Body.Close()

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Response status: %d - %s\n", res.StatusCode, string(resBody))
		fmt.Println("Donation created successfully")

		// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	
	r.Close()
	fmt.Println("Stopping consumer donation...")
}