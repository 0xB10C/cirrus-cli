/*
 * Provides a container compatible interface.
 *
 * This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /var/run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v1.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/libpod/images/json' | jq
 *
 * API version: 0.0.1
 * Contact: podman@lists.podman.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// VolumeInfo Volume list response
type VolumeInfo struct {
	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`
	// Name of the volume driver used by the volume. Only supports local driver
	Driver string `json:"Driver"`
	// User-defined key/value metadata. Always included
	Labels map[string]string `json:"Labels,omitempty"`
	// Mount path of the volume on the host.
	Mountpoint string `json:"Mountpoint"`
	// Name of the volume.
	Name string `json:"Name"`
	// The driver specific options used when creating the volume.
	Options map[string]string `json:"Options"`
	// The level at which the volume exists. Libpod does not implement volume scoping, and this is provided solely for Docker compatibility. The value is only \"local\".
	Scope string `json:"Scope"`
}
