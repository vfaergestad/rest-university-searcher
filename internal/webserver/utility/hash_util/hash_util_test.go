package hash_util

import "testing"

func TestHashPolicy(t *testing.T) {
	type args struct {
		country string
		scope   string
	}
	tests := []struct {
		name string
		args []args
		want bool
	}{
		{
			name: "EqualPolicies",
			args: []args{
				{
					country: "Norway",
					scope:   "2020-01-01",
				},
				{
					country: "Norway",
					scope:   "2020-01-01",
				},
			},
			want: true,
		},
		{
			name: "DifferentPolicies",
			args: []args{
				{
					country: "Norway",
					scope:   "2020-01-01",
				},
				{
					country: "Sweden",
					scope:   "2020-01-01",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashPolicy(tt.args[0].country, tt.args[0].scope) == HashPolicy(tt.args[1].country, tt.args[1].scope)
			if got != tt.want {
				t.Errorf("HashPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashWebhook(t *testing.T) {
	type args struct {
		url     string
		country string
		calls   int
	}
	tests := []struct {
		name string
		args []args
		want bool
	}{
		{
			name: "EqualWebhooks",
			args: []args{
				{
					url:     "https://example.com/webhook",
					country: "Norway",
					calls:   1,
				},
				{
					url:     "https://example.com/webhook",
					country: "Norway",
					calls:   1,
				},
			},
			want: true,
		},
		{
			name: "DifferentWebhooks",
			args: []args{
				{
					url:     "https://example.com/webhook",
					country: "Norway",
					calls:   1,
				},
				{
					url:     "https://example.com/webhook2",
					country: "Sweden",
					calls:   2,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashWebhook(tt.args[0].url, tt.args[0].country, tt.args[0].calls) == HashWebhook(tt.args[1].url, tt.args[1].country, tt.args[1].calls)
			if got != tt.want {
				t.Errorf("HashWebhook() = %v, want %v", got, tt.want)
			}
		})
	}
}
