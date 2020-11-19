package template

var (
	ServantImp = `package main

import (
	"context"
)

// {{.Servant}}Imp servant implementation
type {{.Servant}}Imp struct {
}

// Init servant init
func (imp *{{.Servant}}Imp) Init() (error) {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *{{.Servant}}Imp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *{{.Servant}}Imp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *{{.Servant}}Imp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
`
)
