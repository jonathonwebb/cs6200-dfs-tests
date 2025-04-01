package cs6200_dfs_tests

import (
	"flag"
	"testing"
)

var clientPath string

func TestPart1ClientStat(t *testing.T) {

}

func TestMain(m *testing.M) {
	flag.StringVar(&clientPath, "client", "bin/dfs-client-p1", "client binary path")
	flag.Parse()
	m.Run()
}
