package Server

import (
	"ReAce/Logger"
	"ReAce/Package"
	"ReAce/Sessions"
	"ReAce/Task"
	"errors"
	"fmt"
	"net"
	"sync"
)

type AceServer struct {
	listeners  *Listeners
	sessions   sync.Map
	C          chan Package.Package
	Downloader Task.DownLoader
}

func NewAceServer(listeners *Listeners) *AceServer {

	temp := &AceServer{listeners: listeners,
		C: make(chan Package.Package),
	}
	go temp.LoopChanMessage()
	return temp
}
func (s *AceServer) CheckUUID(uuid string) bool {
	_, ok := s.sessions.Load(uuid)
	if !ok {
		return false
	} else {
		return true
	}
}
func (s *AceServer) GetListeners() *Listeners {
	return s.listeners
}

func (s *AceServer) AddSession(uuid string, conn *Sessions.Session) {
	_, ok := s.sessions.Load(uuid)
	if !ok {

		s.sessions.Store(uuid, conn)
		Logger.ALogger.Debug("有主机上线:", conn.GetIp())
		return
	} else {
		Logger.ALogger.Debug("主机已存在:", conn.GetIp())
	}

}
func (s *AceServer) GetSession(uuid string) (error, *Sessions.Session) {
	val, ok := s.sessions.Load(uuid)
	if ok {
		return nil, val.(*Sessions.Session)
	} else {
		return errors.New("未找到该uuid"), nil
	}

}
func (s *AceServer) DelSession(uuid string) {
	val, ok := s.sessions.Load(uuid)
	if ok {
		session := val.(*Sessions.Session)
		session.GetConn().(net.Conn).Close()
		Logger.ALogger.Debug("有主机下线:", session.GetIp())
		s.sessions.Delete(uuid)
		return
	} else {

		Logger.ALogger.Debug(uuid, "的主机不存在:")
	}

}
func (s *AceServer) ShowSessions() {
	s.sessions.Range(func(id, val any) bool {
		session := val.(*Sessions.Session)
		fmt.Println("uuid:", id, "-ip:", session.GetIp())
		return true
	})
}
func (s *AceServer) GetSessionsUUIDS() []string {
	result := make([]string, 0)
	result = append(result, "")
	s.sessions.Range(func(id, val any) bool {
		//session := val.(*Sessions.Session)
		result = append(result, id.(string))
		return true
	})

	return result
}

func (s *AceServer) LoopChanMessage() {
	for {
		select {
		case pk := <-s.C:
			{
				err, session := s.GetSession(pk.GetUUID())

				if err != nil {
					fmt.Println("无法获取")
					break
				}
				//fmt.Println(session.GetIp())
				session.SendPkg(pk)
			}

		}
	}

}
func (s *AceServer) AddSendTask(pack Package.Package) {
	s.C <- pack
}

var Server *AceServer
