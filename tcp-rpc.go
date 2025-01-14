package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"supervisord/message"
	"time"
)

const (
	MESSAGE_VERSION_SIZE     int = 4
	MESSAGE_ID_SIZE          int = 4
	MESSAGE_DATA_LENGTH_SIZE int = 4
	HEAD_SIZE                int = 10
)

type SupervisorTcp struct {
	supervisor  *Supervisor
	ipAddress   string
	port        string
	appID       string
	instanceID  string
	accessToken string
	groupID     string
	//	conn *net.TCPConn
}

func NewSupervisorTcp(ipAddress string, port string, supervisor *Supervisor) *SupervisorTcp {
	return &SupervisorTcp{supervisor: supervisor, port: port, ipAddress: ipAddress}
}

func connectServer(st *SupervisorTcp, done chan int)(*net.TCPConn,error)  {
	tcpAddress:=st.getAddress()
	tcpAdd, err := net.ResolveTCPAddr("tcp", tcpAddress)
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error:", err)
		return nil,err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdd) //raddr是指远程地址，laddr是指本地地址，连接服务端
	if err != nil {
		fmt.Println("net.DailTCP error:", err)
		return nil,err
	}
	fmt.Println("connected")
	readTimeout := 30 * time.Second
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(readTimeout)
	instanceLoginReq := &message.InstanceLoginReq{InstanceID: st.instanceID, AppId: st.appID, GroupID: st.groupID, AccessToken: st.accessToken}
	sendData, err := proto.Marshal(instanceLoginReq)
	var b bytes.Buffer
	x := int32(3)
	binary.Write(&b, binary.BigEndian, x)
	messageBuffer := packMessage(b.Bytes(), 3, sendData)
	_, err = conn.Write(messageBuffer)
	if err != nil {
		done <- 1
		<-done
		fmt.Println(err)
	}
	return conn,nil
}
func (st *SupervisorTcp) getAddress()string{
	if st.port == "" {
		st.port = "50001"
	}

	tcpAddress := st.ipAddress + ":" + st.port
	return tcpAddress
}
// CreateSocket create tcp socker
func (st *SupervisorTcp) CreateSocket() error {

	c := make(chan int)
	conn,err := connectServer(st,c)
	defer conn.Close()
	t1 := time.NewTicker(time.Second * 10)
	go onMessageReceived(conn, st, c, t1) //读取服务端广播的信息
	go sendHeartBeat(conn, c, t1,st)
	//st.conn = conn
	//for {
	//	// 自己发送的信息
	//	var data string
	//	fmt.Scan(&data)
	//	if data == "quit"{
	//		break
	//	}
	//	b := []byte(data + "\n")
	//	conn.Write(b)
	//}
	for {
		time.Sleep(10 * time.Second)
	}
	return err
}

func packMessage(messageVersionBuf []byte, messageID int32, sendData []byte) []byte {
	var b bytes.Buffer
	binary.Write(&b, binary.BigEndian, messageVersionBuf)
	binary.Write(&b, binary.BigEndian, messageID)
	dataLength := int32(len(sendData))
	binary.Write(&b, binary.BigEndian, &dataLength)
	binary.Write(&b, binary.BigEndian, sendData)
	return b.Bytes()
}

func sendHeartBeat(conn *net.TCPConn, done chan int, heartBeatTimer *time.Ticker,st *SupervisorTcp) {
	id := 0
	//t1 := time.NewTimer(time.Second * 5)

	if conn != nil {
		for {
			select {
			case <-done:
				fmt.Println("exiting...")
				done <- 1
				break
			case <-heartBeatTimer.C:

				value := int32(id + 1)
				if value > 10000000 {
					value = 0
				}

				heartbeatReq := &message.HeartbeatReq{Id: value}
				sendData, err := proto.Marshal(heartbeatReq)
				if err != nil {
					fmt.Println(err)
				}
				var b bytes.Buffer
				x := int32(1)
				binary.Write(&b, binary.BigEndian, x)
				messageBuffer := packMessage(b.Bytes(), 2, sendData)
				if messageBuffer != nil {
					_, err = conn.Write(messageBuffer)
					if err != nil {
						fmt.Println(err)
						tcpConn,err := connectServer(st,done)
						if err!=nil{
							fmt.Println(err)
						}
						conn = tcpConn
					}
				}
			default:
			}
		}
	}
}

//onMessageReceived 获取服务端发送来的信息
func onMessageReceived(conn *net.TCPConn, st *SupervisorTcp, done chan int, heartBeatTimer *time.Ticker) {
	var (
		buffer               = message.NewBuffer(conn, 10)
		messageDataLengthBuf []byte
		messageVersionBuf    []byte
		messageIdBuf         []byte
		contentSize          int
		contentBuf           []byte
	)
	for {
		select {
		case <-done:
			fmt.Println("exiting...")
			done <- 1
			break
		default:
		}

		_, err := buffer.ReadFromReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			messageVersionBuf, err = buffer.Seek(MESSAGE_VERSION_SIZE)
			if err != nil {
				break
			}

			messageIdBuf, err = buffer.Seek(MESSAGE_ID_SIZE)
			if err != nil {
				break
			}
			version := int(binary.BigEndian.Uint32(messageVersionBuf))
			fmt.Println(version)
			messageId := int(binary.BigEndian.Uint32(messageIdBuf))
			fmt.Println(messageId)

			messageDataLengthBuf, err = buffer.SeekOffset(MESSAGE_ID_SIZE+MESSAGE_VERSION_SIZE, MESSAGE_DATA_LENGTH_SIZE)
			if err != nil {
				break
			}
			contentSize = int(binary.BigEndian.Uint32(messageDataLengthBuf))
			if buffer.Len() >= contentSize+HEAD_SIZE {
				contentBuf = buffer.Read(HEAD_SIZE, contentSize)

				//fmt.Println(string(contentBuf))
				if messageId == 2 {
					heartbeat := &message.HeartbeatReq{}
					proto.Unmarshal(contentBuf, heartbeat)
					if heartbeat.Id > 0 {
						heartbeatAck := &message.HeartbeatAck{Id: heartbeat.Id + 1}
						sendData, err := proto.Marshal(heartbeatAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 1, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
						}

					}
				} else if messageId == 5 {
					startAppReq := &message.StartAppReq{}
					proto.Unmarshal(contentBuf, startAppReq)
					if startAppReq.ApplicationName != "" {
						startArgs := StartProcessArgs{Name: startAppReq.ApplicationName, Wait: true}
						result := struct{ Success bool }{false}
						err := st.supervisor.StartProcess(nil, &startArgs, &result)
						errMessage := ""
						if err != nil {
							errMessage = err.Error()
						}
						startAppAck := &message.StartAppAck{Result: result.Success, Message: errMessage}
						sendData, err := proto.Marshal(startAppAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 6, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
						//	heartBeatTimer.Reset(time.Millisecond)
						}
					}

				} else if messageId == 7 {
					stopAppReq := &message.StopAppReq{}
					proto.Unmarshal(contentBuf, stopAppReq)
					if stopAppReq.ApplicationName != "" {
						startArgs := StartProcessArgs{Name: stopAppReq.ApplicationName, Wait: true}
						result := struct{ Success bool }{false}
						err := st.supervisor.StopProcess(nil, &startArgs, &result)
						errMessage := ""
						if err != nil {
							errMessage = err.Error()
						}
						stopAppAck := &message.StopAppAck{Result: result.Success, Message: errMessage}
						sendData, err := proto.Marshal(stopAppAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 8, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
							//heartBeatTimer.(time.Millisecond)
						}
					}

				} else if messageId == 9 {
					getLogReq := &message.GetLogReq{}
					proto.Unmarshal(contentBuf, getLogReq)
					if getLogReq.ApplicationName != "" {
						processLogReadInfo := ProcessLogReadInfo{Name: getLogReq.ApplicationName, Offset: 5000, Length: 5000}

						logDataInfo := struct{ LogData string }{""}
						err := st.supervisor.ReadProcessStdoutLog(nil, &processLogReadInfo, &logDataInfo)
						errMessage := ""
						if err != nil {
							errMessage = err.Error()
						}
						getLogAck := &message.GetLogAck{ErrorMessage: errMessage, Message: logDataInfo.LogData}
						sendData, err := proto.Marshal(getLogAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 10, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
							//heartBeatTimer.Reset(time.Millisecond)
						}
					}

				} else if messageId == 11 {
					startProgramesReq := &message.StartProgramesReq{}
					proto.Unmarshal(contentBuf, startProgramesReq)
					errMessage := ""
					result := struct{ Success bool }{false}
					if len(startProgramesReq.ApplicationNames) > 0 {
						//processLogReadInfo := ProcessLogReadInfo{Name: getLogReq.ApplicationName,Offset:5000,Length:5000}

						for _, program := range startProgramesReq.ApplicationNames {
							//sr._startProgram(program)
							startArgs := StartProcessArgs{Name: program, Wait: true}
							oldResult := result.Success
							err := st.supervisor.StartProcess(nil, &startArgs, &result)
							if err != nil {
								errMessage = errMessage + err.Error()
							}
							if !result.Success {
								result.Success = result.Success && oldResult
							}
						}
						startProgramsAck := &message.StartProgramsAck{Status: result.Success, Message: errMessage}
						sendData, err := proto.Marshal(startProgramsAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 12, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
							//heartBeatTimer.Reset(time.Millisecond)
						}
					}

				} else if messageId == 13 {
					stopProgramesReq := &message.StopProgramesReq{}
					proto.Unmarshal(contentBuf, stopProgramesReq)
					errMessage := ""
					result := struct{ Success bool }{false}
					if len(stopProgramesReq.ApplicationNames) > 0 {
						//processLogReadInfo := ProcessLogReadInfo{Name: getLogReq.ApplicationName,Offset:5000,Length:5000}

						for _, program := range stopProgramesReq.ApplicationNames {
							//sr._startProgram(program)
							startArgs := StartProcessArgs{Name: program, Wait: true}
							oldResult := result.Success
							err := st.supervisor.StopProcess(nil, &startArgs, &result)
							if err != nil {
								errMessage = errMessage + err.Error()
							}
							if !result.Success {
								result.Success = result.Success && oldResult
							}
						}
						stopProgramsAck := &message.StopProgramsAck{Status: result.Success, Message: errMessage}
						sendData, err := proto.Marshal(stopProgramsAck)
						if err != nil {
							fmt.Println(err)
							break
						}
						messageBuffer := packMessage(messageVersionBuf, 14, sendData)
						if messageBuffer != nil {
							_, err = conn.Write(messageBuffer)
							if err != nil {
								fmt.Println(err)
							}
						}
					}

				}
				continue
			}
			break
		}
	}

	//reader := bufio.NewReader(conn)
	//for {
	//	// var data string
	//	msg,err := reader.ReadString('\n')  //读取直到输入中第一次发生 ‘\n’
	//	fmt.Println(msg)
	//	if err!=nil{
	//		fmt.Println("err:",err)
	//		os.Exit(1)    //服务端错误的时候，就将整个客户端关掉
	//	}
	//}
}
