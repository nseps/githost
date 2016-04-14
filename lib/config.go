package githost

import (
	"os/exec"
	"strconv"
)

type Config string

func get(conf string, key string) (value string, err error) {
	cmd := exec.Command("git", "config", "-f", conf, key)
	out, err := cmd.Output()
	out = out[:len(out)-1]
	return string(out), err
}

func put(conf string, key string, value string) error {
	cmd := exec.Command("git", "config", "-f", conf, key, value)
	_, err := cmd.Output()
	return err
}

//String functions

func (c Config) Get(section string, key string) (value string, err error)  {
	return get(string(c), section + "." + key)
}

func (c Config) GetWithURL(section string, url string, key string) (value string, err error)  {
	return get(string(c), section + "." + url + "." + key)
}

func (c Config) Put(section string, key string, value string) error  {
	return put(string(c), section + "." + key, value)
}

func (c Config) PutWithURL(section string, url string, key string, value string) error  {
	return put(string(c), section + "." + url + "." + key, value)
}

//Int functions

func (c Config) GetInt(section string, key string) (value int, err error)  {
	val, err := get(string(c), section + "." + key)
	if err != nil {
		return
	}
	v, err := strconv.Atoi(val)
	return v, err
}

func (c Config) GetIntWithURL(section string, url string, key string) (value int, err error)  {
	val, err := get(string(c), section + "." + url + "." + key)
	if err != nil {
		return
	}
	v, err := strconv.Atoi(val)
	return v, err
}

func (c Config) PutInt(section string, key string, value int) error  {
	return put(string(c), section + "." + key, strconv.Itoa(value))
}

func (c Config) PutIntWithURL(section string, url string, key string, value int) error  {
	return put(string(c), section + "." + url + "." + key, strconv.Itoa(value))
}