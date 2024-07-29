package schemas

type CreateSchema struct {
	Prop1 string `json:"prop1" xml:"prop1"`
	Prop2 int    `json:"prop2" xml:"prop2"`
}

type UpdateSchema struct {
	Prop1 string `json:"prop1" xml:"prop1"`
	Prop2 string `json:"prop2" xml:"prop2"`
}

type EntitySchema struct {
	ID    interface{} `json:"id" xml:"id"`
	Prop1 string      `json:"prop1" xml:"prop1"`
	Prop2 string      `json:"prop2" xml:"prop2"`
}
