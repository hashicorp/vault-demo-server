// Copyright IBM Corp. 2015, 2026

package api

// SegmentList returns all the available LAN segments.
func (op *Operator) SegmentList(q *QueryOptions) ([]string, *QueryMeta, error) {
	var out []string
	qm, err := op.c.query("/v1/operator/segment", &out, q)
	if err != nil {
		return nil, nil, err
	}
	return out, qm, nil
}
