package requests

type ClassRoomRequest struct {
	ClassName string `json:"class_name" validate:"required"`
	Major     string `json:"major" validate:"required"`
	ClassYear string `json:"class_year" validate:"required"`
}

type CreateClassRoomRequest struct {
	Major     string `json:"major" validate:"required"`
	ClassName int    `json:"class_name" validate:"required"`
	ClassYear int    `json:"class_year" validate:"required"`
}

type UpdateClassRoomRequest struct {
	Id            uint   `json:"id" validate:"required"`
	ClassName     string `json:"class_name" validate:"required,min=4,max=5"`
	ClassRoomCode string `json:"class_room_code" validate:"required,min=4,max=5"`
}
type DeleteClassRoomRequest struct {
	Id            uint   `json:"id" validate:"required"`
	ClassRoomCode string `json:"class_room_code" validate:"required,min=4,max=5"`
}

type ClassRoomIDRequest struct {
	Id uint `json:"id" validate:"required"`
}
