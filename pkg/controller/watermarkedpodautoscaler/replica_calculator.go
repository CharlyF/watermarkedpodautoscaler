package watermarkedpodautoscaler

import (
	"fmt"
	"math"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	metricsclient "k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
)



const (
	// defaultTestingTolerance is default value for calculating when to
	// scale up/scale down.
	defaultTestingTolerance = 0.1
)

// ReplicaCalculator is responsible for calculation of the number of replicas
// It contains all the needed information
type ReplicaCalculator struct {
	metricsClient metricsclient.MetricsClient
	podsGetter    v1coreclient.PodsGetter
	tolerance     float64
}

// NewReplicaCalculator returns a ReplicaCalculator object reference
func NewReplicaCalculator(metricsClient metricsclient.MetricsClient, podsGetter v1coreclient.PodsGetter, tolerance float64) *ReplicaCalculator {
	return &ReplicaCalculator{
		metricsClient: metricsClient,
		podsGetter:    podsGetter,
		tolerance:     tolerance,
	}
}
// GetExternalMetricReplicas calculates the desired replica count based on a
// target metric value (as a milli-value) for the external metric in the given
// namespace, and the current replica count.
func (c *ReplicaCalculator) GetExternalMetricReplicas(currentReplicas int32, targetUtilization int64, metricName, namespace string, selector *metav1.LabelSelector) (replicaCount int32, utilization int64, timestamp time.Time, err error) {
	labelSelector, err := metav1.LabelSelectorAsSelector(selector)
	if err != nil {
		return 0, 0, time.Time{}, err
	}
	log.Info(fmt.Sprintf("Using %v", labelSelector))
	// should do c.metricsClient.GetExternalMetric(metricName, namespace, labelSelector)
	metrics, timestamp, err := []int64{14}, time.Now(), nil
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("unable to get external metric %s/%s/%+v: %s", namespace, metricName, selector, err)
	}
	utilization = 0
	for _, val := range metrics {
		utilization = utilization + val
	}

	usageRatio := float64(utilization) / float64(targetUtilization)
	if math.Abs(1.0-usageRatio) <= c.tolerance {
		// return the current replicas if the change would be too small
		return currentReplicas, utilization, timestamp, nil
	}

	return int32(math.Ceil(usageRatio * float64(currentReplicas))), utilization, timestamp, nil
}
