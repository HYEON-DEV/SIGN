package config

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		URL      string `yaml:"url"`      // 호스트와 포트 통합 URL
		Database string `yaml:"database"` // 데이터베이스 이름
		User     string `yaml:"user"`     // 사용자 이름
		Password string `yaml:"password"` // 비밀번호
	} `yaml:"mysql"`

	SpringBoot struct {
		Port      int    `yaml:"port"`       // springboot 포트
		Host      string `yaml:"host"`       // springboot 호스트
		SpringURL string `yaml:"spring_url"` // springboot URL
	} `yaml:"spring_boot"`

	Server struct {
		Port  int    `yaml:"port"`   // Go서버 포트
		Host  string `yaml:"host"`   // Go서버 호스트
		GoURL string `yaml:"go_url"` // Go서버 URL
	} `yaml:"server"`

	ECDSA struct {
		PrivateKeyPath string `yaml:"private_key_path"` // 개인키 파일경로
		PublicKeyPath  string `yaml:"public_key_path"`  // 공개키 파일경로
	} `yaml:"ecdsa"`
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

// config.yml 파일 업데이트 (주석과 형식 유지)
func SaveConfig(privateKeyPath, publicKeyPath string) error {
	// YAML 파일을 문자열로 읽기
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		return fmt.Errorf("config.yml 파일 읽기 실패: %v", err)
	}

	// 정규식을 사용해 private_key_path와 public_key_path 업데이트
	rePrivate := regexp.MustCompile(`private_key_path:\s*".*"`)
	rePublic := regexp.MustCompile(`public_key_path:\s*".*"`)

	updatedData := rePrivate.ReplaceAllString(string(data), fmt.Sprintf(`private_key_path: "%s"`, privateKeyPath))
	updatedData = rePublic.ReplaceAllString(updatedData, fmt.Sprintf(`public_key_path: "%s"`, publicKeyPath))

	// 업데이트된 내용을 파일에 저장
	err = os.WriteFile("config/config.yml", []byte(updatedData), 0644)
	if err != nil {
		return fmt.Errorf("config.yml 파일 저장 실패: %v", err)
	}

	return nil
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
