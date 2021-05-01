package ci45

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	type testcase struct {
		name string
		args args
		want string
	}
	tests := []testcase{
		// TODO: Add test cases.
		{
			"test1",
			args{[]int{10, 2}},
			"102",
		},
		{
			"test2",
			args{[]int{3, 30, 34, 5, 9}},
			"3033459",
		},
		{
			"test4",
			args{[]int{1440, 7548, 4240, 6616, 733, 4712, 883, 8, 9576}},
			"1440424047126616733754888389576",
		},
	}
	for _, tt := range tests {
		t.Logf("\tChecking \"%#v\" for result: \"%s\"", tt.args, tt.want)
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			} else {
				t.Log("Get the result: ", got)
			}
		})
	}
}