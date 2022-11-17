package service

import (
	"Gin_todo/model"
	"Gin_todo/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0 未做 1 已做
}

// 展示任务详情的服务
type ShowTaskService struct {
}

// 展示任务列表的服务
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.Db.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
	}
	err := model.Db.Create(&task).Error
	if err != nil {
		code := 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.Db.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败！",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

// 列表返回所有用户备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 5
	}
	model.Db.Model(&model.Task{}).Preload("User").Where("uid = ?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTasks(tasks),
	}
}
