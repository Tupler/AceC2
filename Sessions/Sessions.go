package Sessions

import (
	"ReAce/Logger"
	"ReAce/Package"
	"net"
	"time"
)

type Session struct {
	uuid        string
	ip          string
	installTime string
	conn        any
}

func (s *Session) GetIp() string {
	return s.ip
}
func (s *Session) GetInstallTime() string {
	return s.installTime
}
func (s *Session) GetConn() any {
	return s.conn
}

func NewSession(uuid string, conn any, ip string) *Session {

	return &Session{
		uuid:        uuid,
		conn:        conn,
		ip:          ip,
		installTime: time.Now().String(),
	}
}

func (s *Session) SendPkg(pkg Package.Package) {

	if tcp, ok := s.conn.(net.Conn); ok {

		_, err := tcp.Write(pkg.MakePackage())
		if err != nil {
			Logger.ALogger.Error("错误传输！")
			return
		}
		Logger.ALogger.Success("发送成功")
	}
	//} else if ws, ok := s.conn.(websocket.Conn); ok {
	//
	//}
}
