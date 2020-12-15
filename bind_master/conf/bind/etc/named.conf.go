//include "/etc/bind/zones/go/internal/go.com.br.conf";
view "internal-view" {
	match-clients { net_85; };
	include "/etc/bind/zones/go/internal/go.com.br.conf";
	include "/etc/bind/zones.rfc1918";
	include "/etc/bind/named.conf.default-zones";
};

view "external-view" {
//	include "/etc/bind/zones/go/external/go.com.br.conf";
	match-clients { any; };
	recursion no;
	include "/etc/bind/zones/go/external/go.com.br.conf";
};

acl net_85 {
	10.85.0.20/32;
};
