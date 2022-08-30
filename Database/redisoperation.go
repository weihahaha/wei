package Database

import (
	"log"
)

func (re *Redis)RedisLogin(user string, pwd string)  {
	conn := re.RedisDb.Get()
	defer conn.Close()

	_, err := conn.Do(user, pwd)
	if err != nil{
		log.Fatalln("Redis:", err)
		conn.Close()
	}

}
