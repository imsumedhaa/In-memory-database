package database


type Database interface{
	Create()
	Update()
	Delete()
	Get()
	Show()
	Exit()
	Set()

}