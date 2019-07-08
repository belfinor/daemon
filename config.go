package daemon

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-07-08

type Config struct {
	PidFile string `json:"pid"`
	LogFile string `json:"log"`
	WordDir string `json:"dir"`
}
