package database


type Database interface{
	Create(key, value string) error
	Update(key,value string) error
	Delete(key string) error
	Get(key string) error
	Show() error
	Exit() error

}