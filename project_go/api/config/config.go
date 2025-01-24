package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		URL      string `yaml:"url"`      // 호스트와 포트 통합 URL
		Database string `yaml:"database"` // 데이터베이스 이름
		User     string `yaml:"user"`     // 사용자 이름
		Password string `yaml:"password"` // 비밀번호
	} `yaml:"mysql"`
}

// 전역 변수 설정 - 사용 시, main에서 설정 로드만 해두면 GConfig로 언제든지 불러올 수 있음
var GConfig *Config

// yml 파일을 읽어 구조체로 변환
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config/config.yml") // config.yml 파일 경로 입력
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// 전역변수 설정
	GConfig = &config

	return &config, nil
}

// DB 연동을 위한 접속 명령어 생성 - yml에 입력한 DB 접속 정보를 조합하여 명령어 생성(문자열)
func (cfg *Config) MySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=5s",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.URL,
		cfg.MySQL.Database,
	)
}
