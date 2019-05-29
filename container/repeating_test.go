package cal

import "testing"

func TestSubString(t *testing.T){
	//1 第一步写测试数据 和期望数据
	tests:= []struct{a string;b int}{
		 {"abcabcbb",3},
		 {"bbbbb",1},
		 {"pwwkew",3},
		 {"这里是慕课网",6},
		 {"一二三二一",3},
	}

	//2.第二步循环判断
	for _,value:= range tests{
		result:=lengthOfNonRepeatingSubStr(value.a)
		if result!=value.b {
			t.Errorf("lengthOfNonRepeatingSubStr(%s);got %d,expected %d",value.a,result,value.b)
		}
	}
}


//BenchmarkSubString-8   	10000000	       219 ns/op   进行了1000W次   每个for循环里面执行1次耗时 219NS
func BenchmarkSubString(b *testing.B){
	test:=struct{a string;b int}{"abcabcbb",3}

	for i:=0;i<b.N;i++{
		result:=lengthOfNonRepeatingSubStr(test.a)
		if result!=test.b {
			b.Errorf("lengthOfNonRepeatingSubStr(%s);got %d,expected %d",test.a,result,test.b)
		}
	}
}
