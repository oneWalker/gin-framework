package utils_test

import (
	"gin-practice/pkg/utils"
	"testing"
)

//unit testing
func TestEncodeMD5(t *testing.T) {
	cryptoStr := utils.EncodeMD5("testing")
	if cryptoStr == "" {
		t.Error("GenShortID failed")
	}
}

//performance testing
func BenchmarkEncodeMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.EncodeMD5("string")
	}
}

//performace analysis
//#使用命令，生成CPU的profile
//go test -v -bench="BenchmarkGenShortID$" --run=none -cpuprofile cpu.out ./utils/

//观察耗费时间，Graphviz可视化分析图依赖包
//go tool 主要是对生成的文件进行分析
//go tool pprof cpu.out
//top
//# or
//go tool pprof -http=":" cpu.out

//测试覆盖率
//go test -v -coverprofile=cover.out ./util/
//-coverprofile=cover.out 选项可以统计测试覆盖率
//cover -func=cover.out 可以查看更加详细的测试覆盖率的结果
//go test -v -coverprofile=cover.out ./utils/
