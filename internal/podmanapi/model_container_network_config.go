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

// ContainerNetworkConfig contains information on a container's network configuration.
type ContainerNetworkConfig struct {
	// CNINetworks is a list of CNI networks to join the container to. If this list is empty, the default CNI network will be joined instead. If at least one entry is present, we will not join the default network (unless it is part of this list). Only available if NetNS is set to bridge. Optional.
	CniNetworks []string `json:"cni_networks,omitempty"`
	// DNSOptions is a set of DNS options that will be used in the container's resolv.conf, replacing the host's DNS options which are used by default. Conflicts with UseImageResolvConf. Optional.
	DnsOption []string `json:"dns_option,omitempty"`
	// DNSSearch is a set of DNS search domains that will be used in the container's resolv.conf, replacing the host's DNS search domains which are used by default. Conflicts with UseImageResolvConf. Optional.
	DnsSearch []string `json:"dns_search,omitempty"`
	// DNSServers is a set of DNS servers that will be used in the container's resolv.conf, replacing the host's DNS Servers which are used by default. Conflicts with UseImageResolvConf. Optional.
	DnsServer [][]int32 `json:"dns_server,omitempty"`
	// Expose is a number of ports that will be forwarded to the container if PublishExposedPorts is set. Expose is a map of uint16 (port number) to a string representing protocol. Allowed protocols are \"tcp\", \"udp\", and \"sctp\", or some combination of the three separated by commas. If protocol is set to \"\" we will assume TCP. Only available if NetNS is set to Bridge or Slirp, and PublishExposedPorts is set. Optional.
	Expose *interface{} `json:"expose,omitempty"`
	// HostAdd is a set of hosts which will be added to the container's etc/hosts file. Conflicts with UseImageHosts. Optional.
	Hostadd []string `json:"hostadd,omitempty"`
	Netns *Namespace `json:"netns,omitempty"`
	// NetworkOptions are additional options for each network Optional.
	NetworkOptions map[string][]string `json:"network_options,omitempty"`
	// PortBindings is a set of ports to map into the container. Only available if NetNS is set to bridge or slirp. Optional.
	Portmappings []PortMapping `json:"portmappings,omitempty"`
	// PublishExposedPorts will publish ports specified in the image to random unused ports (guaranteed to be above 1024) on the host. This is based on ports set in Expose below, and any ports specified by the Image (if one is given). Only available if NetNS is set to Bridge or Slirp.
	PublishImagePorts bool `json:"publish_image_ports,omitempty"`
	StaticIp *[]int32 `json:"static_ip,omitempty"`
	StaticIpv6 *[]int32 `json:"static_ipv6,omitempty"`
	StaticMac *[]int32 `json:"static_mac,omitempty"`
	// UseImageHosts indicates that /etc/hosts should not be managed by Podman, and instead sourced from the image. Conflicts with HostAdd.
	UseImageHosts bool `json:"use_image_hosts,omitempty"`
	// UseImageResolvConf indicates that resolv.conf should not be managed by Podman, but instead sourced from the image. Conflicts with DNSServer, DNSSearch, DNSOption.
	UseImageResolveConf bool `json:"use_image_resolve_conf,omitempty"`
}
