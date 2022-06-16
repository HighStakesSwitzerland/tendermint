// Code generated by metricsgen. DO NOT EDIT.

package p2p

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		PeersConnected: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_connected",
			Help:      "Number of peers connected.",
		}, labels).With(labelsAndValues...),
		PeersStored: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_stored",
			Help:      "Nomber of peers in the peer store database.",
		}, labels).With(labelsAndValues...),
		PeersInactivated: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_inactivated",
			Help:      "Number of inactive peers stored.",
		}, labels).With(labelsAndValues...),
		PeerReceiveBytesTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_receive_bytes_total",
			Help:      "Number of bytes per channel received from a given peer.",
		}, append(labels, "peer_id", "chID", "message_type")).With(labelsAndValues...),
		PeerSendBytesTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_send_bytes_total",
			Help:      "Number of bytes per channel sent to a given peer.",
		}, append(labels, "peer_id", "chID", "message_type")).With(labelsAndValues...),
		PeerPendingSendBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_pending_send_bytes",
			Help:      "Number of bytes pending being sent to a given peer.",
		}, append(labels, "peer_id")).With(labelsAndValues...),
		PeersConnectedSuccess: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_connected_success",
			Help:      "Number of successful connection attempts",
		}, labels).With(labelsAndValues...),
		PeersConnectedFailure: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_connected_failure",
			Help:      "Number failed connection attempts",
		}, labels).With(labelsAndValues...),
		PeersConnectedIncoming: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_connected_incoming",
			Help:      "Number of peers connected as a result of dialing the peer.",
		}, labels).With(labelsAndValues...),
		PeersConnectedOutgoing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers_connected_outgoing",
			Help:      "Number of peers connected as a result of the peer dialing this node.",
		}, labels).With(labelsAndValues...),
		RouterPeerQueueRecv: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_peer_queue_recv",
			Help:      "The time taken to read off of a peer's queue before sending on the connection.",
		}, labels).With(labelsAndValues...),
		RouterPeerQueueSend: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_peer_queue_send",
			Help:      "The time taken to send on a peer's queue which will later be read and sent on the connection.",
		}, labels).With(labelsAndValues...),
		RouterChannelQueueSend: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_channel_queue_send",
			Help:      "The time taken to send on a p2p channel's queue which will later be consued by the corresponding reactor/service.",
		}, labels).With(labelsAndValues...),
		PeerQueueDroppedMsgs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_channel_queue_dropped_msgs",
			Help:      "The number of messages dropped from a peer's queue for a specific p2p Channel.",
		}, append(labels, "ch_id")).With(labelsAndValues...),
		PeerQueueMsgSize: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_queue_msg_size",
			Help:      "The size of messages sent over a peer's queue for a specific p2p Channel.",
		}, append(labels, "ch_id")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		PeersConnected:         discard.NewGauge(),
		PeersStored:            discard.NewGauge(),
		PeersInactivated:       discard.NewGauge(),
		PeerReceiveBytesTotal:  discard.NewCounter(),
		PeerSendBytesTotal:     discard.NewCounter(),
		PeerPendingSendBytes:   discard.NewGauge(),
		PeersConnectedSuccess:  discard.NewCounter(),
		PeersConnectedFailure:  discard.NewCounter(),
		PeersConnectedIncoming: discard.NewGauge(),
		PeersConnectedOutgoing: discard.NewGauge(),
		RouterPeerQueueRecv:    discard.NewHistogram(),
		RouterPeerQueueSend:    discard.NewHistogram(),
		RouterChannelQueueSend: discard.NewHistogram(),
		PeerQueueDroppedMsgs:   discard.NewCounter(),
		PeerQueueMsgSize:       discard.NewGauge(),
	}
}
