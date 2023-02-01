// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/hashicorp/vault/physical"
)

// Physical is a vault physical.Backend implementation that limits the
// maximum size of data that can be written to Vault in order to prevent
// abuse.
type Physical struct {
	physical.Backend

	Limit uint64

	current map[string]uint64
	total   uint64
	lock    sync.Mutex
	once    sync.Once
}

func (p *Physical) Put(ctx context.Context, e *physical.Entry) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.once.Do(p.init)

	// Determine if we'll pass the threshold
	written := p.total
	written -= p.current[e.Key]
	written += uint64(len(e.Value))
	if written > p.Limit {
		return fmt.Errorf("physical storage exceeded for client")
	}

	// Nope, write it
	if err := p.Backend.Put(ctx, e); err != nil {
		return err
	}

	// Update accounting
	p.current[e.Key] = uint64(len(e.Value))
	p.total = written
	return nil
}

func (p *Physical) Delete(ctx context.Context, key string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.once.Do(p.init)

	// Nope, write it
	if err := p.Backend.Delete(ctx, key); err != nil {
		return err
	}

	// Update accounting
	p.total -= p.current[key]
	delete(p.current, key)
	return nil
}

func (p *Physical) init() {
	p.current = make(map[string]uint64)
}
