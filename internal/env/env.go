package env
import (
	"os"
	"strconv"
	"time"
)
 
func GetString(key, fallback string) string{
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetDuration(duration string) time.Duration{

	val , err := time.ParseDuration(duration)
		if err != nil {
			return  0 
		}
		return val
}

func GetInt(key string, fallback int) int{
	val , ok := os.LookupEnv(key)
	if !ok{
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil{
		return fallback
	}
	return valAsInt
}