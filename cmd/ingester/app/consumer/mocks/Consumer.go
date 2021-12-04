package mocks

import cluster "github.com/bsm/sarama-cluster"

import mock "github.com/stretchr/testify/mock"

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Consumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkPartitionOffset provides a mock function with given fields: topic, partition, offset, metadata
func (_m *Consumer) MarkPartitionOffset(topic string, partition int32, offset int64, metadata string) {
	_m.Called(topic, partition, offset, metadata)
}

// Partitions provides a mock function with given fields:
func (_m *Consumer) Partitions() <-chan cluster.PartitionConsumer {
	ret := _m.Called()

	var r0 <-chan cluster.PartitionConsumer
	if rf, ok := ret.Get(0).(func() <-chan cluster.PartitionConsumer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan cluster.PartitionConsumer)
		}
	}

	return r0
}
