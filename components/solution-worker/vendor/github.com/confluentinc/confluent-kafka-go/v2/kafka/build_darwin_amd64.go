// +build !dynamic


// This file was auto-generated by librdkafka_vendor/bundle-import.sh, DO NOT EDIT.

package kafka

// #cgo CFLAGS: -DUSE_VENDORED_LIBRDKAFKA -DLIBRDKAFKA_STATICLIB
// #cgo LDFLAGS: ${SRCDIR}/librdkafka_vendor/librdkafka_darwin_amd64.a  -lm -lsasl2 -ldl -lpthread -framework CoreFoundation -framework SystemConfiguration
import "C"

// LibrdkafkaLinkInfo explains how librdkafka was linked to the Go client
const LibrdkafkaLinkInfo = "static darwin_amd64 from librdkafka-static-bundle-v2.1.1.tgz"