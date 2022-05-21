package models

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/models"
)

const CALLBACK_EVENT_CHANGE_TYPE_UPDATE_TAG   = "update_tag"

type EventTagUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	TagID         string   `xml:"TagId"`
	AddUserItems  string   `xml:"AddUserItems"`
	DelUserItems  string   `xml:"DelUserItems"`
	AddPartyItems string   `xml:"AddPartyItems"`
	DelPartyItems string   `xml:"DelPartyItems"`
}

