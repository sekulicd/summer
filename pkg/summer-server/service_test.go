package summerservice

import "testing"

func TestSummerService_addTuples(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given 2&3, addTuples should return sum 5",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "given 4&3, addTuples should return sum 5",
			args: args{
				a: 4,
				b: 3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SummerService{}
			if got := s.AddTuple(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addTuples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummerService_addTriple(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given 4&3&7, addTriples should return sum 14",
			args: args{
				a: 4,
				b: 3,
				c: 7,
			},
			want: 14,
		},
		{
			name: "given 5&3&7, addTriples should return sum 15",
			args: args{
				a: 5,
				b: 3,
				c: 7,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SummerService{}
			if got := s.AddTriple(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("addTriple() = %v, want %v", got, tt.want)
			}
		})
	}
}
