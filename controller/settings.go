package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LoadSettings(c *gin.Context) {
	// 从settings.json文件中读取字典列表
	settings, err := readSettingsFromFile("settings.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "无法加载设置！",
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": settings,
	})
}

func SaveSettings(c *gin.Context) {
	// 从请求中解析字典列表
	var settings []map[string]interface{}
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "无效的请求数据！",
			"data": err,
		})
		return
	}
	// 将字典列表保存到settings.json文件
	err := writeSettingsToFile(settings, "settings.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "设置保存失败！",
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": settings,
	})
}

func readSettingsFromFile(filename string) ([]map[string]interface{}, error) {
	// 检查settings.json文件是否存在
	if !fileExists(filename) {
		return []map[string]interface{}{}, nil
	}
	// 从文件中读取JSON数据
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 检查文件是否为空
	if len(data) == 0 {
		return []map[string]interface{}{}, nil
	}
	// 解析JSON数据为字典列表
	var settings []map[string]interface{}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func writeSettingsToFile(settings []map[string]interface{}, filename string) error {
	// 检查settings.json文件是否存在，如果不存在则创建
	if !fileExists(filename) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	// 将字典列表转换为JSON数据
	data, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	// 将JSON数据写入文件
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}