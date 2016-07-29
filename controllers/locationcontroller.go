//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/zjwdmlmx/odcsock/application"
	"github.com/zjwdmlmx/odcsock/application/proto"
	"github.com/zjwdmlmx/odcsock/global"
	"github.com/zjwdmlmx/odcsock/model"
)

const (
	loginPath  = "/user/devicelogin"
	oldSosPath = "/user/sos"
)

var (
	host string
)

type loginResponse struct {
	Req int    `json:"req"`
	Sid string `json:"sid"`
}

type loginRequest struct {
	Imei string `json:"imei"`
	Pass string `json:"pass"`
}

type oldSosRequest struct {
	Sid  string  `json:"sid"`
	Date int64   `json:"date"`
	La   float64 `json:"la"`
	Lo   float64 `json:"lo"`
	Acc  uint    `json:"acc"`
}

// LocationController controller of location logic
type LocationController struct {
	client  *http.Client
	_buffer *bytes.Buffer
}

// electricity report
func (ctrl *LocationController) reportElectoricity(cmd *proto.V1Command) {
	if cmd.Electricity < 0 || cmd.Electricity > 100 {
		return
	}

	R := global.RedisPool.Get()

	if _, err := R.Do("SETEX", cmd.Id+".ele", 3600, cmd.Electricity); err != nil {
		log.Println(err.Error())
	}
}

// login the odcser server
func (ctrl *LocationController) loginDevice(cmd *proto.V1Command) (sessionid string, err error) {
	var (
		device   model.WearableDevice
		jsonByte []byte
	)

	global.DB.Where("imei = ?", cmd.Id).First(&device)
	jsonByte, err = json.Marshal(loginRequest{Imei: cmd.Id, Pass: device.Password})

	if err != nil {
		return
	}
	log.Println(string(jsonByte))

	var res *http.Response
	res, err = ctrl.client.Post(host+loginPath, "application/json", bytes.NewBuffer(jsonByte))

	if err != nil {
		return
	}

	/**
	 * Reading login response for session id
	 */

	resData := make([]byte, 1024)
	var readCount int

	ctrl._buffer.Reset()
	for {
		readCount, err = res.Body.Read(resData)
		if readCount > 0 {
			ctrl._buffer.Write(resData[0:readCount])
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return
		}
	}

	resObj := loginResponse{}
	err = json.Unmarshal(ctrl._buffer.Bytes(), &resObj)

	if err != nil {
		return
	}

	log.Println(resObj)
	if resObj.Req == 0 {
		sessionid = resObj.Sid
	} else {
		err = errors.New("device login failed")
	}
	return
}

// send the sos message to odcser server for the device client
func (ctrl *LocationController) sendOldSos(cmd *proto.V1Command) (err error) {
	var (
		sessionid string
		jsonByte  []byte
	)

	sessionid, err = ctrl.loginDevice(cmd)

	if err != nil {
		return
	}

	jsonByte, err = json.Marshal(oldSosRequest{
		Sid:  sessionid,
		La:   cmd.Latitude,
		Lo:   cmd.Longitude,
		Date: cmd.Time.Unix() * 1000,
		Acc:  11,
	})

	_, err = ctrl.client.Post(host+oldSosPath, "application/json", bytes.NewBuffer(jsonByte))

	if err != nil {
		return
	}
	return
}

/**
 * get user' id with device IMEI
 *
 * @param imei the imei code for device
 * @return user's id and error
 */
func (ctrl *LocationController) getUID(imei string) (uid uint64, err error) {
	var (
		cacheErr error
		user     model.User
	)
	// take user id from cache. If user id is not exist take from database
	// and set to cache
	uid, cacheErr = global.Cached.SGetUint64(imei)

	if cacheErr != nil {
		global.DB.Select("id").Where("imei = ?", imei).First(&user)
		if user.Id == 0 {
			err = errors.New("unregiste device")
			return
		}

		uid = user.Id
		err = global.Cached.SSetUint64(imei, uid, 3600) // expire time is one hour

		if err != nil {
			return
		}
	}
	return
}

func (ctrl *LocationController) Init() (err error) {
	// initial the controller
	ctrl.client = global.Client
	ctrl._buffer = &bytes.Buffer{}

	var ok bool
	host, ok = global.Config.Get("odcserServerHost")
	if !ok {
		log.Println("odcser server 's host is not set in configure file.Using default https://localhost'")
		host = "https://localhost"
	}

	return
}

//Handle func to handle V1 command
func (ctrl *LocationController) Handle(incomingMsg *application.IncomingMessage, replay *application.Replay) (err error) {
	Command, ok := incomingMsg.Command.(*proto.V1Command)

	if !ok {
		err = errors.New("unknow error in controller LocationController, when parse Command")
		return
	}

	log.Println(*Command)
	log.Println(incomingMsg.Params)

	var uid uint64

	if uid, err = ctrl.getUID(Command.Id); err != nil {
		return
	}

	if uid == 0 {
		return
	}

	// is the incoming message a SOS
	if Command.State == proto.DEVICE_STATE_SOS {
		err = ctrl.sendOldSos(Command)

		if err != nil {
			return
		}
	}

	// electricity report
	ctrl.reportElectoricity(Command)

	// is the incoming message's data valid
	if !Command.Valid {
		return
	}

	currPosition := model.Position{
		Latitude:  Command.Latitude,
		Longitude: Command.Longitude,
		Time:      Command.Time,
		Uid:       uid,
	}

	// store the current position
	global.DB.Create(&currPosition)

	return
}
