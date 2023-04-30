package common

// import "x-ui/database/model"

// func GenVmessLink(inbound *model.Inbound) string {
// 	if inbound.Protocol.(*string) == "vmess" {
// 		return ""
// 	}
// 	return ""
// 	// network := inbound.StreamSettings
// 	// let network = this.stream.network;
// 	// let type = 'none';
// 	// let host = '';
// 	// let path = '';
// 	// if (network === 'tcp') {
// 	// 	let tcp = this.stream.tcp;
// 	// 	type = tcp.type;
// 	// 	if (type === 'http') {
// 	// 		let request = tcp.request;
// 	// 		path = request.path.join(',');
// 	// 		let index = request.headers.findIndex(header => header.name.toLowerCase() === 'host');
// 	// 		if (index >= 0) {
// 	// 			host = request.headers[index].value;
// 	// 		}
// 	// 	}
// 	// } else if (network === 'kcp') {
// 	// 	let kcp = this.stream.kcp;
// 	// 	type = kcp.type;
// 	// 	path = kcp.seed;
// 	// } else if (network === 'ws') {
// 	// 	let ws = this.stream.ws;
// 	// 	path = ws.path;
// 	// 	let index = ws.headers.findIndex(header => header.name.toLowerCase() === 'host');
// 	// 	if (index >= 0) {
// 	// 		host = ws.headers[index].value;
// 	// 	}
// 	// } else if (network === 'http') {
// 	// 	network = 'h2';
// 	// 	path = this.stream.http.path;
// 	// 	host = this.stream.http.host.join(',');
// 	// } else if (network === 'quic') {
// 	// 	type = this.stream.quic.type;
// 	// 	host = this.stream.quic.security;
// 	// 	path = this.stream.quic.key;
// 	// } else if (network === 'grpc') {
// 	// 	path = this.stream.grpc.serviceName;
// 	// }

// 	// if (this.stream.security === 'tls') {
// 	// 	if (!ObjectUtil.isEmpty(this.stream.tls.server)) {
// 	// 		address = this.stream.tls.server;
// 	// 	}
// 	// }

// 	// let obj = {
// 	// 	v: '2',
// 	// 	ps: remark,
// 	// 	add: address,
// 	// 	port: this.port,
// 	// 	id: this.settings.vmesses[clientIndex].id,
// 	// 	aid: this.settings.vmesses[clientIndex].alterId,
// 	// 	net: network,
// 	// 	type: type,
// 	// 	host: host,
// 	// 	path: path,
// 	// 	tls: this.stream.security,
// 	// 	sni: this.stream.tls.settings.serverName,
// 	// 	fp: this.stream.tls.settings.fingerprint,
// 	// 	alpn: this.stream.tls.alpn.join(','),
// 	// 	allowInsecure: this.stream.tls.settings.allowInsecure,
// 	// };
// 	// return 'vmess://' + base64(JSON.stringify(obj, null, 2));
// }
