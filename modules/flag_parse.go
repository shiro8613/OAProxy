package modules

import "flag"

type path struct {
	Config_path string
	Logsdir_path string
	UserList_path string
}

var (
	config_path = flag.String("config", "./config.yml", "config.yml path")
	log_dir_path = flag.String("logDir", "./Logs", "Log Directory path")
	userlist_path = flag.String("userlist", "./Userlog.json", "LoginUserList path")
)

var path_config []path


func ParseFlag() {
	flag.Parse()
	path_config = append(path_config, path{
		Config_path: *config_path,
		Logsdir_path: *log_dir_path,
		UserList_path: *userlist_path,
	})
}

func GetFlag() path {
	for _, v := range path_config {
		return v
	}
	return path_config[0]
}