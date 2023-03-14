//go:build unit

package setbk

import (
	"reflect"
	"testing"
)

func TestGetPrettySummary(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want PrettySummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrettySummary(tt.args.market); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrettySummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettyIndex(t *testing.T) {
	type args struct {
		i index
	}
	tests := []struct {
		name string
		args args
		want pIndex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyIndex(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prettyIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettyInvestorSummary(t *testing.T) {
	type args struct {
		is investorSummary
	}
	tests := []struct {
		name string
		args args
		want pInvestorSummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyInvestorSummary(tt.args.is); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prettyInvestorSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettyRankings(t *testing.T) {
	type args struct {
		rr []ranking
	}
	tests := []struct {
		name string
		args args
		want []pRanking
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyRankings(tt.args.rr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prettyRankings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettyStocks(t *testing.T) {
	type args struct {
		ss []stock
	}
	tests := []struct {
		name string
		args args
		want []pStock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyStocks(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prettyStocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSummary(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want Summary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSummary(tt.args.market); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIndex(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want index
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndex(tt.args.market); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInvestorSummary(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want investorSummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInvestorSummary(tt.args.market); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInvestorSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRanking(t *testing.T) {
	type args struct {
		market string
		rType  string
	}
	tests := []struct {
		name string
		args args
		want ranking
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRanking(tt.args.market, tt.args.rType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexURL(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexURL(tt.args.market); got != tt.want {
				t.Errorf("indexURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_investorSummaryURL(t *testing.T) {
	type args struct {
		market string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := investorSummaryURL(tt.args.market); got != tt.want {
				t.Errorf("investorSummaryURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rankingURL(t *testing.T) {
	type args struct {
		market string
		rType  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankingURL(tt.args.market, tt.args.rType); got != tt.want {
				t.Errorf("rankingURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
