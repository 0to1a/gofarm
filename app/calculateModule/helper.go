package calculateModule

import "framework/framework"

type httpStructure struct{}
type serviceStructure struct{}

var (
	http      httpStructure
	service   serviceStructure
	cron      framework.CronUtils
	utils     framework.Utils
	webserver framework.WebServer
)

func calculatePlus(n1 int64, n2 int64) int64 {
	return n1 + n2
}
