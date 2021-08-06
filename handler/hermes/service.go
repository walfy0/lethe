package hermes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
	"github.com/sirupsen/logrus"
)

type Config struct {
	User     		string `json:"user"`
	Password 		string `json:"password"`
	Host     		string `json:"host"`
	Port     		int    `json:"port"`
	EncrypMethod 	string `json:"encryp_method"`
}

func GetConfig()(Config, error){
	ConfigJson := Config{}
	data, err := ioutil.ReadFile("../hermes/config.json")
	if err != nil {
		logrus.Warnf("[GetConfig]err: %v", err)
		return ConfigJson, err
	}
	err = json.Unmarshal(data, &ConfigJson)
	if err != nil {
		logrus.Warnf("[GetConfig]err: %v", err)
		return ConfigJson, err
	}
	return ConfigJson, nil
}

func GetStatus(c *gin.Context) {
	configJson, err := GetConfig()
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, common.SuccessResp(map[string]interface{}{
			"method": configJson.Host,
			"status": false,
		}))
		return
	}
	cmd := exec.Command("/bin/bash", "-c", "lsof -i:1081")
	lsof, err := cmd.Output()
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, common.SuccessResp(map[string]interface{}{
			"method": configJson.EncrypMethod,
			"status": false,
		}))
		return
	}
	logrus.Info(strings.Trim(string(lsof), "\n"))
	c.JSON(http.StatusOK, common.SuccessResp(map[string]interface{}{
		"method": configJson.EncrypMethod,
		"status": true,
	}))
}

type HermesReq struct{
	Hermes			string `json:"hermes"`
	Command 		string `json:"command"`
	EncrypMethod 	string `json:"encryp_method"`
}

func ChangeHermes(c *gin.Context){
	configJson, err := GetConfig()
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusOK, common.SuccessResp(map[string]interface{}{
			"method": configJson.Host,
			"status": false,
		}))
		return
	}
	req := HermesReq{}
	if err := c.BindJSON(&req); err != nil {
		logrus.Info("err: ",err)
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
		return
	}
	configJson.EncrypMethod = req.EncrypMethod
	result, _ := json.MarshalIndent(configJson, "", "    ")
	ioutil.WriteFile("../hermes/config.json",result, os.FileMode(os.O_TRUNC))
	if req.Command == "start" {
		shell := "cd ../hermes/ && nohup sh build.sh &"
		if req.Hermes == "service" {
			shell = "cd ../hermes/proxy && nohup sh build.sh &"
		}
		cmd := exec.Command("/bin/bash", "-c", shell)
		err := cmd.Start()
		if err != nil {
			logrus.Info(err)
			c.JSON(200, common.ErrorResp(common.ServiceError, nil))
			return 
		}
	} else {
		cmd := exec.Command("/bin/bash", "-c", "lsof -i:1081")
		lsof, err := cmd.Output()
		if err != nil {
			logrus.Info(err)
			c.JSON(200, common.ErrorResp(common.ServiceError, nil))
			return 
		}
		result := strings.Split(string(lsof), "\n")[1]
		tmp := strings.Split(result, " ")
		for i, t := range tmp {
			if i>0&&t!=""{
				result = t
				break
			}
		}
		shell := ""
		if req.Command == "close" {
			shell = "kill -9 " + result
		}else if req.Command == "restart"{
			shell = "kill " + result
		}
		logrus.Info("start " + req.Command)
		cmd = exec.Command("/bin/bash", "-c", shell)
		_, err = cmd.Output()
		if err != nil {
			logrus.Info(err)
			c.JSON(200, common.ErrorResp(common.ServiceError, nil))
			return 
		}
	}
	c.JSON(http.StatusOK, common.SuccessResp(nil))
}
