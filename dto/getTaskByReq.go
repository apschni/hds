package dto

type GetTaskByReq struct {
	CategoryId string   `json:"category_id"`
	SubjectIds []string `json:"subject_ids"`
}
