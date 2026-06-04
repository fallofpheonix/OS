// STATUS: STUB
// RUNTIME: NO
// PRODUCTION_READY: NO
package live

import "errors"

var ErrNotImplemented = errors.New("live: not implemented")

type ProbeLoader struct{}

func (p *ProbeLoader) Load() error { return ErrNotImplemented }
