package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type PipelineFunc func(ctx *map[string]interface{}) error
type StepFunc func(handleFunc PipelineFunc) PipelineFunc

func new() func(ctx *map[string]interface{}) (error error) {
	return func(ctx *map[string]interface{}) (error error) {
		fmt.Println(0)
		return nil
	}
}
func sum1() StepFunc {
	return func(next PipelineFunc) PipelineFunc {
		return func(ctx *map[string]interface{}) error {
			fields := log.Fields{
				"k": "1",
			}
			log.WithFields(fields).Info("pushlog")
			return next(ctx)
		}
	}
}

func sum2() StepFunc {
	return func(next PipelineFunc) PipelineFunc {
		return func(ctx *map[string]interface{}) error {
			fields := log.Fields{
				"k": "2",
			}
			log.WithFields(fields).Info("pushlog")
			return nil
		}
	}
}

func sum3() StepFunc {
	return func(next PipelineFunc) PipelineFunc {
		return func(ctx *map[string]interface{}) error {
			err := next(ctx)
			fields := log.Fields{
				"k": "3",
			}
			log.WithFields(fields).Info("pushlog")
			return err
		}
	}
}

func Pipeline(h PipelineFunc, steps ...StepFunc) PipelineFunc {
	h = chainSteps(h, steps...)
	return h
}

func chainSteps(h PipelineFunc, middleware ...StepFunc) PipelineFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

func main() {
	steps := []StepFunc{
		sum3(),
		sum2(),
		sum1(),
	}
	h := Pipeline(new(), steps...)

	m := &map[string]interface{}{
		"k": "value",
	}

	err := h(m)
	if err != nil {
		fmt.Println(err)
		return
	}
}
