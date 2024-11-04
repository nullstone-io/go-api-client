package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvEvent_IsEqual(t *testing.T) {
	tests := []struct {
		a    EnvEvent
		b    EnvEvent
		want bool
	}{
		{
			a: EnvEvent{
				Actions: []EventAction{EventActionAppDeployed, EventActionBlockNeedsApproval},
				Blocks:  []int64{31},
				Channels: map[IntegrationTool]ChannelData{
					IntegrationToolSlack: {
						"connections": []map[string]string{
							{
								"channel_id": "C01DBR86SRK",
							},
						},
					},
				},
				Statuses: []EventStatus{},
			},
			b: EnvEvent{
				Actions: []EventAction{EventActionBlockNeedsApproval, EventActionAppDeployed},
				Blocks:  []int64{31},
				Channels: map[IntegrationTool]ChannelData{
					IntegrationToolSlack: {
						"connections": []map[string]any{
							{
								"channel_id": "C01DBR86SRK",
							},
						},
					},
				},
				Statuses: nil,
			},
			want: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a, b := test.a, test.b
			got := a.IsEqual(b)
			assert.Equal(t, test.want, got)
		})
	}
}
