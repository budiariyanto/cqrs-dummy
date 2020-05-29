package main

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"

	"github.com/budiariyanto/cqrs-dummy/service"
	"github.com/budiariyanto/cqrs-dummy/repo"
	"github.com/budiariyanto/cqrs-dummy/payloads"
	"github.com/budiariyanto/cqrs-dummy/common"
)

func main() {
	fmt.Println("Starting consumer campaignership...")
	dbCampaignership := common.NewDB("campaignership")
	campaignershipRepo := repo.NewCampaignershipRepo(dbCampaignership)
	campaignershipService := service.NewCampaignershipService(campaignershipRepo)
	
	UpdateTotalDonationConsumer(campaignershipService)
}

func UpdateTotalDonationConsumer(s *service.CampaignershipService) {
	r := common.NewKafkaReader("group2", "donation-create")
	
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Println("Read message from campaignership consumer:", string(m.Key), string(m.Value))
		var msg payloads.CreateDonationMessage
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}

		totalDonation := int64(msg.Amount) + msg.TotalDonation
		if totalDonation > msg.TargetDonation {
			reason := "Donation target reached"
			err = s.Close(msg.CampaignID, reason)
			if err != nil {
				fmt.Println(err)
				break
			}

			payload := payloads.CampaignDetailPayload {
				ID:             msg.CampaignID,
				IsClosed:       true,
				Reason:         reason,
			}
			
			buf, err := json.Marshal(payload)
			if err != nil {
				fmt.Println(err)
				break
			}

			client := http.DefaultClient
			req, _ := http.NewRequest("PUT", "http://localhost:1323/campaign-detail/update-campaign-status", bytes.NewBuffer(buf))
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
			fmt.Println("Campaign updated successfully")
		}
		// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	
	r.Close()
	fmt.Println("Stopping consumer campaignership...")
}