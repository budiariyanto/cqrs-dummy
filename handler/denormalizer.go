package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/budiariyanto/cqrs-dummy/payloads"
)

func (h *Handler) GetCampaignDetail(c echo.Context) (err error) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}

	campaignDetail, err := h.DenormalizeSvc.GetCampaignDetail(idInt)
	if err != nil {
		return
	}

	resp := payloads.CampaignDetailPayload(campaignDetail)
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateCampaignStatus(c echo.Context) (err error) {
	cd := new(payloads.CampaignDetailPayload)
	if err = c.Bind(cd); err != nil {
		return
	}

	oldCampaignDetail, err := h.DenormalizeSvc.GetCampaignDetail(cd.ID)
	if err != nil {
		return
	}

	err = h.DenormalizeSvc.UpdateCampaignStatus(oldCampaignDetail.ID, cd.IsClosed, cd.Reason)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, cd)
}

func (h *Handler) UpdateTotalDonation(c echo.Context) (err error) {
	cd := new(payloads.CampaignDetailPayload)
	if err = c.Bind(cd); err != nil {
		return
	}

	oldCampaignDetail, err := h.DenormalizeSvc.GetCampaignDetail(cd.ID)
	if err != nil {
		return
	}

	err = h.DenormalizeSvc.UpdateTotalDonation(oldCampaignDetail.ID, cd.TotalDonation)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, cd)
}