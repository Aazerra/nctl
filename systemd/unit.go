package systemd

import (
	"context"
	"fmt"
	"log"
	"nctl/utils"

	"github.com/coreos/go-systemd/v22/dbus"
)

type SystemD struct {
	con *dbus.Conn
}

func New(ctx context.Context) SystemD {
	con, err := dbus.NewWithContext(ctx)
	utils.FailOnErr(err)

	return SystemD{
		con: con,
	}
}

func (s *SystemD) RestartUnit(ctx context.Context, name string) {
	var result chan (string)

	_, err := s.con.TryRestartUnitContext(ctx, name, "fail", result)
	utils.FailOnErr(err)
	log.Printf("%s has been restarted.", name)
}
func (s *SystemD) StartUnit(ctx context.Context, name string) {
	var result chan (string)

	_, err := s.con.StartUnitContext(ctx, name, "fail", result)
	utils.FailOnErr(err)
	log.Printf("%s has been started.", name)

}
func (s *SystemD) StopUnit(ctx context.Context, name string) {
	var result chan (string)

	_, err := s.con.StopUnitContext(ctx, name, "fail", result)
	utils.FailOnErr(err)
	log.Printf("%s has been stopped.", name)

}
func (s *SystemD) StatusUnit(ctx context.Context, name string) {
	res, err := s.con.ListUnitsByNamesContext(ctx, []string{name})
	utils.FailOnErr(err)
	fmt.Println(res[0].ActiveState)
}

func (s *SystemD) Close() {
	s.con.Close()
}
