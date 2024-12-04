package api

import "go.thebigfile.com/renterd/webhooks"

type WebhookResponse struct {
	Webhooks []webhooks.Webhook          `json:"webhooks"`
	Queues   []webhooks.WebhookQueueInfo `json:"queues"`
}
