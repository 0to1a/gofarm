package exampleModule

import "framework/framework"

type httpStructure struct{}
type ServiceStructure struct{}

var (
	http      httpStructure
	service   ServiceStructure
	utils     framework.Utils
	webserver framework.WebServer
)

func calculatePlus(n1 int64, n2 int64) int64 {
	return n1 + n2
}
