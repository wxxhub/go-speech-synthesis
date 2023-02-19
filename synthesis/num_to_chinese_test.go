package synthesis

import (
	"fmt"
	"github.com/go-ego/gpy"
	"github.com/mozillazg/go-pinyin"
	"testing"
)

func TestGoPinYin(t *testing.T) {
	a := pinyin.NewArgs()
	a.Heteronym = true
	a.Style = pinyin.Tone3

	t.Log(pinyin.Pinyin("音频重置重要璦", a))
	t.Log(pinyin.Convert("音频重置重要璦", &a))
	t.Log(pinyin.LazyPinyin("音频重置重要璦", a))
	t.Log(pinyin.LazyConvert("音频重置重要璦龔鷦", &a))
	//t.Log(pinyin.SinglePinyin("音频重置重要", &a))
}
func TestGoEgo(t *testing.T) {
	hans := "音频重置重要"
	// 默认
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong] [guo] [hua]]

	// 包含声调
	a.Style = gpy.Tone
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhōng] [guó] [huà]]

	// 声调用数字表示
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng] [guo2] [hua4]]

	// 开启多音字模式
	a = gpy.NewArgs()
	a.Heteronym = true
	fmt.Println(gpy.Pinyin(hans, a))
	//gpy.AddDict()
	// [[zhong zhong] [guo] [hua]]
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [hua4]]

	fmt.Println(gpy.LazyPinyin(hans, gpy.NewArgs()))
	// [zhong guo hua]

	fmt.Println(gpy.Convert(hans, nil))
	// [[zhong] [guo] [hua]]

	fmt.Println(gpy.LazyConvert(hans, nil))
}
