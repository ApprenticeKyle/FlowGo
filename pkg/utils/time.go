package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Time 自定义时间类型，用于JSON序列化和反序列化
// 零值会序列化为 null，非零值使用 RFC3339 格式
type Time struct {
	time.Time
}

// NewTime 创建新的 Time
func NewTime(t time.Time) Time {
	return Time{Time: t}
}

// MarshalJSON 实现 JSON 序列化
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format(time.RFC3339))
}

// UnmarshalJSON 实现 JSON 反序列化
// 支持多种时间格式：
// - RFC3339: "2006-01-02T15:04:05Z07:00"
// - RFC3339Nano: "2006-01-02T15:04:05.999999999Z07:00"
// - 日期时间: "2006-01-02T15:04:05", "2006-01-02 15:04:05"
// - 仅日期: "2006-01-02" (默认时间为 00:00:00，使用本地时区)
func (t *Time) UnmarshalJSON(data []byte) error {
	// 处理 null 或空字符串
	str := string(data)
	if str == "null" || str == `""` || str == `null` {
		t.Time = time.Time{}
		return nil
	}

	// 移除 JSON 字符串的引号
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// 如果为空字符串，返回零值
	if str == "" {
		t.Time = time.Time{}
		return nil
	}

	// 尝试多种时间格式（按常见程度排序）
	formats := []string{
		time.RFC3339,                    // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,                // "2006-01-02T15:04:05.999999999Z07:00"
		"2006-01-02T15:04:05Z07:00",     // 带时区
		"2006-01-02T15:04:05Z",          // UTC
		"2006-01-02T15:04:05",           // 无时区
		"2006-01-02 15:04:05",           // 空格分隔
		"2006-01-02T15:04",              // 无秒
		"2006-01-02 15:04",              // 无秒，空格分隔
		"2006-01-02",                    // 仅日期
	}

	for _, format := range formats {
		if parsed, err := time.Parse(format, str); err == nil {
			// 如果是仅日期格式（"2006-01-02"），使用本地时区的 00:00:00
			if format == "2006-01-02" {
				// 保持日期，但设置为本地时区的 00:00:00
				loc, _ := time.LoadLocation("Local")
				parsed = time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, loc)
			}
			t.Time = parsed
			return nil
		}
	}

	// 如果所有格式都失败，尝试标准 JSON 时间解析
	var stdTime time.Time
	if err := json.Unmarshal(data, &stdTime); err != nil {
		return err
	}
	t.Time = stdTime
	return nil
}

// Value 实现 driver.Valuer 接口，用于数据库存储
func (t Time) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现 sql.Scanner 接口，用于数据库读取
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		// 尝试直接解析为时间，如果失败则尝试 JSON 反序列化
		if parsed, err := time.Parse(time.RFC3339, string(v)); err == nil {
			t.Time = parsed
			return nil
		}
		return t.UnmarshalJSON(v)
	case string:
		// 尝试直接解析为时间，如果失败则尝试 JSON 反序列化
		if parsed, err := time.Parse(time.RFC3339, v); err == nil {
			t.Time = parsed
			return nil
		}
		return t.UnmarshalJSON([]byte(`"` + v + `"`))
	default:
		return fmt.Errorf("unsupported Scan type for Time: %T", value)
	}
}

// IsZero 检查时间是否为零值
func (t Time) IsZero() bool {
	return t.Time.IsZero()
}

// String 返回字符串表示
func (t Time) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(time.RFC3339)
}

// Format 格式化时间
func (t Time) Format(layout string) string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(layout)
}
