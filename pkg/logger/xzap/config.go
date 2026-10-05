/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 5:21 PM
 * Description:
 **/

package xzap

// Config 配置信息
type Config struct {
	ServiceName string
	Mode        string
	Path        string
	Level       string
	Compress    bool
	KeepDays    int
}
