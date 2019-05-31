package cal

import "testing"

func TestTriangle(t *testing.T){ //右键 代码覆盖率

	//分别有3个元素结构的数组
	tests:=[]struct{a,b,c float64}{
		{3,4,5},
		{5,12,13},
		{0,1,1},
		{8,15,17},
	}


	for _,tt :=range tests{
		actual:=CalTriangle(tt.a,tt.b)
		if actual!=tt.c {
			t.Errorf("calTriange(%f,%f);got %f,expected %f",tt.a,tt.b,actual,tt.c)
		}
	}



}