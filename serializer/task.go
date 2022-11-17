package serializer

import "Gin_todo/model"

type Task struct {
	ID        uint   `json:"ID" example:"1"`       // 任务ID
	Title     string `json:"Title" example:"吃饭"`   // 题目
	Content   string `json:"Content" example:"睡觉"` // 内容
	View      uint64 `json:"View" example:"32"`    // 浏览量
	Status    int    `json:"Status" example:"0"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreateAt:  item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
