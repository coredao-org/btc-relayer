package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	CrossChainConfig CrossChainConfig `json:"cross_chain_config"`
	BTCConfig        BTCConfig        `json:"btc_config"`
	COREConfig       COREConfig       `json:"core_config"`
	LogConfig        LogConfig        `json:"log_config"`
	AlertConfig      AlertConfig      `json:"alert_config"`
}

type CrossChainConfig struct {
	RecursionHeight int64 `json:"recursion_height"`
}

func (cfg *CrossChainConfig) Validate() {
}

type BTCRpcAddrs struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type BTCConfig struct {
	RpcAddrs                     []BTCRpcAddrs `json:"rpc_addrs"`
	SleepSecond                  uint64        `json:"sleep_second"`
	DataSeedDenyServiceThreshold float64       `json:"data_seed_deny_service_threshold"`
}

func (cfg *BTCConfig) Validate() {
	if len(cfg.RpcAddrs) == 0 {
		panic("rpc endpoint of BTC chain should not be empty")
	}
}

type COREConfig struct {
	PrivateKey                   string   `json:"private_key"`
	Providers                    []string `json:"providers"`
	GasLimit                     uint64   `json:"gas_limit"`
	GasPrice                     uint64   `json:"gas_price"`
	GasIncrease                  uint64   `json:"gas_increase"`
	SleepSecond                  uint64   `json:"sleep_second"`
	DataSeedDenyServiceThreshold float64  `json:"data_seed_deny_service_threshold"`
}

func (cfg *COREConfig) Validate() {
	if len(cfg.Providers) == 0 {
		panic(fmt.Sprintf("provider address of Core Chain should not be empty"))
	}

	if cfg.GasLimit == 0 {
		panic(fmt.Sprintf("gas_limit of Core Chain should be larger than 0"))
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if using file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if using file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if using file logger")
		}
	}
}

type AlertConfig struct {
	EnableAlert     bool  `json:"enable_alert"`
	EnableHeartBeat bool  `json:"enable_heart_beat"`
	Interval        int64 `json:"interval"`

	Identity       string `json:"identity"`
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	BalanceThreshold     string `json:"balance_threshold"`
	SequenceGapThreshold uint64 `json:"sequence_gap_threshold"`
}

func (cfg *AlertConfig) Validate() {
	if !cfg.EnableAlert {
		return
	}
	if cfg.Interval <= 0 {
		panic("alert interval should be positive")
	}
	balanceThreshold, ok := big.NewInt(1).SetString(cfg.BalanceThreshold, 10)
	if !ok {
		panic(fmt.Sprintf("unrecognized balance_threshold"))
	}

	if balanceThreshold.Cmp(big.NewInt(0)) <= 0 {
		panic(fmt.Sprintf("balance_threshold should be positive"))
	}

	if cfg.SequenceGapThreshold <= 0 {
		panic(fmt.Sprintf("sequence_gap_threshold should be positive"))
	}
}

func (cfg *Config) Validate() {
	cfg.CrossChainConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.BTCConfig.Validate()
	cfg.COREConfig.Validate()
	cfg.AlertConfig.Validate()
	//cfg.DBConfig.Validate()
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}

func ParseConfigFromFile(filePath string) *Config {

	fullFilePath := path.Join(GetCurrentAbPath(), string(os.PathSeparator)+filePath)

	fmt.Println("config path:" + fullFilePath)
	//fullFilePath := dir + string(os.PathSeparator) + filePath
	bz, err := ioutil.ReadFile(fullFilePath)
	if err != nil {
		panic(err)
	}

	var config Config

	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}

	config.Validate()

	return &config
}

//getCurrentAbPath
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// go build path
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	res = res + string(os.PathSeparator) + "config"
	return res
}

// go run path
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
