package algorithms_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*************************
 *   通用断言工具类
 *************************/

// AssertEqual 通用相等断言（支持基本类型和error）
func AssertEqual(t *testing.T, got, want interface{}, msg ...string) {
	t.Helper()
	if !isEqual(got, want) {
		errorMsg := fmt.Sprintf("got %v (%T), want %v (%T)", got, got, want, want)
		if len(msg) > 0 {
			errorMsg += " | " + msg[0]
		}
		t.Error(errorMsg)
	}
}

// isEqual 深层比较实现
func isEqual(a, b interface{}) bool {
	switch aVal := a.(type) {
	case []int:
		bVal, ok := b.([]int)
		if !ok || len(aVal) != len(bVal) {
			return false
		}
		for i := range aVal {
			if aVal[i] != bVal[i] {
				return false
			}
		}
		return true
	case error:
		if b == nil {
			return aVal == nil
		}
		return aVal.Error() == b.(error).Error()
	default:
		return a == b
	}
}

/*************************
 *   测试数据生成器
 *************************/

// GenerateRandomIntSlice 生成随机整数切片
// 参数：length-长度, min/max-取值范围
// algorithms/testutils_test.go
func GenerateRandomIntSlice(t *testing.T, length, min, max int) []int {
	t.Helper()
	if length < 0 {
		t.Fatal("长度不能为负数")
	}

	// 创建独立随机源（替代弃用的 rand.Seed）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, length)
	for i := range result {
		result[i] = r.Intn(max-min+1) + min
	}
	return result
}

// GenerateSortedSlice 生成有序切片（可指定升序/降序）
func GenerateSortedSlice(t *testing.T, length int, ascending bool) []int {
	t.Helper()
	result := make([]int, length)
	for i := range result {
		if ascending {
			result[i] = i
		} else {
			result[i] = length - i - 1
		}
	}
	return result
}

/*************************
 *   特殊测试场景工具
 *************************/

// CreatePanicFunc 创建预期panic的测试函数包装器
func CreatePanicFunc(fn func()) (didPanic bool, msg any) {
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
			msg = r
		}
	}()
	fn()
	return
}
