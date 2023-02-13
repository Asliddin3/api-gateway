package repo

type RedisRepo interface{
	Set(key,value string) error
	Get(key string) (interface{},error)
}
