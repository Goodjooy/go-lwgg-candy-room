package index

import "go-lwgg-candy-room/src/manage"



func NewIndexApplication()manage.Application{
	app:=manage.NewApplication("/","D:\\goProject\\go-lwgg-candy-room\\src\\index\\templates\\**\\*","")

	app.AsignViewer(newMainPage())
	
	return app
}