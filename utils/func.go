package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 生成MD5字符
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成密码
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 获取当前执行文件父级目录路径
func GetCurrentAbsPath() string {
	dir := getCurrentAbsPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		dir = getCurrentAbsPathByCaller()
	}
	dir = filepath.Dir(dir)
	return dir
}

// 获取当前执行文件绝对路径(go build)
func getCurrentAbsPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		return "."
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径(go run)
func getCurrentAbsPathByCaller() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		absPath = path.Dir(filename)
	}
	return absPath
}

// interface转字符串
func ItoS(value interface{}) string {
	var result string
	switch value.(type) {
	case string:
		// 将interface转为string字符串类型
		result, _ = value.(string)

	case int:
		// 将interface转为int类型
		rs, _ := value.(int)
		result = strconv.Itoa(rs)

	case int64:
		// 将interface转为int64类型
		rs, _ := value.(int64)
		result = strconv.FormatInt(rs, 10)

	case float32:
		// 将interface转为float32类型
		rs, _ := value.(float32)
		result = strconv.FormatFloat(float64(rs), 'f', 2, 64)

	case float64:
		// 将interface转为float64类型
		rs, _ := value.(float64)
		result = strconv.FormatFloat(rs, 'f', 2, 64)

	default:
		result = ""
	}

	return result
}
