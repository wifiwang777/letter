package service

import "github.com/aisuosuo/letter/api/models"

type OutMessages struct {
	*models.Messages
	CreateAt string `json:"createAt"`
}

func ConvertMessageTime(from []*models.Messages) (out []*OutMessages) {
	for _, message := range from {
		out = append(out, &OutMessages{
			Messages: message,
			CreateAt: message.CreateAt.Format("2006/01/02 15:04:05"),
		})
	}
	return
}
