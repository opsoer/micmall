package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/util"
	"go.uber.org/zap"
)


func InitSentinel(){
	initFlow()
	initBreaker()
}
func initFlow() {
	err := sentinel.InitDefault()
	if err != nil {
		zap.S().Fatalf("初始化sentinel 异常: %v", err)
	}

	//配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "banner_list",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //匀速通过
			Threshold:              10000,
			StatIntervalInMs:       100,
		},

		{
			Resource:               "goods_list",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //匀速通过
			Threshold:              10000,
			StatIntervalInMs:       100,
		},
		{
			Resource:               "goods_detail",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //匀速通过
			Threshold:              10000,
			StatIntervalInMs:       100,
		},

	})

	if err != nil {
		zap.S().Fatalf("加载规则失败: %v", err)
	}
}
type StateChangeTestListener struct {}
func (s *StateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	zap.S().Infof("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *StateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	zap.S().Infof("rule.steategy: %+v, From %s to Open, snapshot: %.2f, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *StateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	zap.S().Infof("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}
func initBreaker() {
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		zap.S().Error(err.Error())
	}
	circuitbreaker.RegisterStateChangeListeners(&StateChangeTestListener{})
	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:         "good_list_breaker",
			Strategy:         circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:   3000,
			MinRequestAmount: 10,
			StatIntervalMs:   500,
			MaxAllowedRtMs:   500,
			Threshold:        0.1,
		},
	})
	if err != nil {
		zap.S().Error(err.Error())
	}

}