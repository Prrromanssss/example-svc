package http

type CreateSomeModel struct {
	Param1 string `json:"param1" validate:"required"`
	Param2 string `json:"param2" validate:"required"`
}

type GetExampleModel struct {
	Param1 string `query:"param1"`
	Param2 string `query:"param2"`
}
