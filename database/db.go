package database


type Database interface{
	Create() error
	Update() error
	Delete() error
	Get() error
	Show() error
	Exit() error

}