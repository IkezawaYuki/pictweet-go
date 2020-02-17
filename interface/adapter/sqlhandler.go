package adapter

type SQLHandler interface {
	Find(out interface{}, where ...interface{}) error
	Exec(sql string, values ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Raw(sql string, values ...interface{}) error
	Create(value interface{}) error
	Save(value interface{}) error
	Delete(value interface{}) error
	Where(out interface{}, query interface{}, args ...interface{}) error
	Close() error
}
