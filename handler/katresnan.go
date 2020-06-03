package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/budiariyanto/cqrs-dummy/common"
	"github.com/budiariyanto/cqrs-dummy/payloads"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func (h *Handler) CreateDonation(c echo.Context) (err error) {
	d := new(payloads.DonationReq)
	if err = c.Bind(d); err != nil {
		fmt.Println("Bind error:", err)
		return
	}
	
	campaignDetail, err := h.DenormalizeSvc.GetCampaignDetail(d.CampaignID)
	if err != nil {
		fmt.Println("Campaign detail error:", err)
		return
	}

	msg := payloads.CreateDonationMessage{
		CampaignID:     d.CampaignID,
		DonorName:      d.DonorName,
		Amount:         d.Amount,
		TargetDonation: campaignDetail.TargetDonation,
		TotalDonation: campaignDetail.TotalDonation,
	}

	v, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	var kafkaWriter = common.NewKafkaWriter()
	
	kmsg := kafka.Message{
		Key: []byte(uuid.New().String()),
		Value: v,
	}

	err = kafkaWriter.WriteMessages(context.Background(), kmsg)
	if err != nil {
		fmt.Println("Write kafka mesage error:", err)
		return
	}

	kafkaWriter.Close()

	return c.JSON(http.StatusOK, d)
}