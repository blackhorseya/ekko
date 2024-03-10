//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

// IWorkflowBiz is the interface for workflow business logic.
type IWorkflowBiz interface {
}
