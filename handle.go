// Copyright 2023 AWS Professional Services
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package remetrics

import (
	"context"
	"errors"
	"github.com/Hoverhuang-er/remetrics/pkg/conf"
	"sync"
)

type Remetrics struct {
	InputCfg     *conf.Config
	InputData    *conf.Input
	IsNewVersion bool
}

type RemetricFunc interface {
	InitLibs(ctx context.Context) error
	Relable(ctx context.Context) (bool, error)
	PusherV3(ctx context.Context, idx int, wg *sync.WaitGroup) (bool, error)
	CheckPromUp(ctx context.Context) (bool, error)
}

// InitLibs is a function that init libs
func (rt *Remetrics) InitLibs(ctx context.Context, r Remetrics) error {
	checkCfg := conf.LoadCfg()
	if checkCfg == nil {
		return errors.New("config is nil")
	}
	return nil
}

// Relable is a function that rebuild metrics labels from metrcis path to data_lable
func (rt *Remetrics) Relable(ctx context.Context, r Remetrics) (bool, error) {
	return false, nil
}

// PusherV3 is a function that push metrics to prometheus with coroutines pool
func (rt *Remetrics) PusherV3(ctx context.Context, r Remetrics, idx int, wg *sync.WaitGroup) (bool, error) {
	return false, nil
}

// CheckPromUp is a function that check prometheus is up or not
func (rt *Remetrics) CheckPromUp(ctx context.Context, r Remetrics) (bool, error) {
	return false, nil
}
