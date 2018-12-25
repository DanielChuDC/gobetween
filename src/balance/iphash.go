package balance

/**
 * iphash.go - iphash balance impl
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

import (
	"errors"
	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/core"
	"hash/fnv"
)

/**
 * Iphash balancer
 */
type IphashBalancer struct{}

/**
 * Constructor
 */
func NewIphashBalancer(cfg config.BalanceConfig) interface{} {
	return &IphashBalancer{}
}

/**
 * Elect backend using iphash strategy
 * Using fnv1a for speed
 */
func (b *IphashBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {

	if len(backends) == 0 {
		return nil, errors.New("Can't elect backend, Backends empty")
	}

	hash := fnv.New32a()
	hash.Write(context.Ip())
	backend := backends[hash.Sum32()%uint32(len(backends))]

	return backend, nil
}