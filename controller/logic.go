package controller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/reechou/holmes"
	"github.com/reechou/wedding-chatroom/config"
	"github.com/reechou/wedding-chatroom/ext"
	"github.com/reechou/wedding-chatroom/models"
)

type Logic struct {
	sync.Mutex

	robotExt   *ext.RobotExt
	weddingExt *ext.WeddingServiceExt

	cfg *config.Config
}

func NewLogic(cfg *config.Config) *Logic {
	l := &Logic{
		cfg: cfg,
	}
	l.robotExt = ext.NewRobotExt(cfg)
	l.weddingExt = ext.NewWeddingServiceExt(cfg)

	models.InitDB(cfg)

	l.init()

	return l
}

func (self *Logic) init() {
	http.HandleFunc("/robot/receive_msg", self.RobotReceiveMsg)

	http.HandleFunc("/rpc", self.runRpc)
}

func (self *Logic) Run() {
	defer holmes.Start(holmes.LogFilePath("./log"),
		holmes.EveryDay,
		holmes.AlsoStdout,
		holmes.DebugLevel).Stop()

	if self.cfg.Debug {
		EnableDebug()
	}

	holmes.Info("server starting on[%s]..", self.cfg.Host)
	holmes.Infoln(http.ListenAndServe(self.cfg.Host, nil))
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func WriteBytes(w http.ResponseWriter, code int, v []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(v)
}

func EnableDebug() {

}
