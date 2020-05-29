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
		return
	}
	
	campaignDetail, err := h.DenormalizeSvc.GetCampaignDetail(d.CampaignID)
	if err != nil {
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
		return
	}

	var kafkaWriter = common.NewKafkaWriter()
	
	kmsg := kafka.Message{
		Key: []byte(uuid.New().String()),
		Value: v,
	}

	err = kafkaWriter.WriteMessages(context.Background(), kmsg)
	if err != nil {
		return
	}

	fmt.Println(err)
	
	kafkaWriter.Close()

	return c.JSON(http.StatusOK, d)
}