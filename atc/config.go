package atc

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"

	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/vars"
	"sigs.k8s.io/yaml"
)

type pendingVarSource struct {
	vs   types.VarSourceConfig
	deps []string
}

func OrderByDependency(c types.VarSourceConfigs) (types.VarSourceConfigs, error) {
	ordered := types.VarSourceConfigs{}
	pending := []pendingVarSource{}
	added := map[string]interface{}{}

	for _, vs := range c {
		b, err := yaml.Marshal(vs.Config)
		if err != nil {
			return nil, err
		}

		template := vars.NewTemplate(b)
		varNames := template.ExtraVarNames()

		dependencies := []string{}
		for _, varName := range varNames {
			parts := strings.Split(varName, ":")
			if len(parts) > 1 {
				dependencies = append(dependencies, parts[0])
			}
		}

		if len(dependencies) == 0 {
			// If no dependency, add the var source to ordered list.
			ordered = append(ordered, vs)
			added[vs.Name] = true
		} else {
			// If there are some dependencies, then check if dependencies have
			// already been added to ordered list, if yes, then add it; otherwise
			// add it to a pending list.
			miss := false
			for _, dep := range dependencies {
				if added[dep] == nil {
					miss = true
					break
				}
			}
			if !miss {
				ordered = append(ordered, vs)
				added[vs.Name] = true
			} else {
				pending = append(pending, pendingVarSource{vs, dependencies})
				continue
			}
		}

		// Once a var_source is added to ordered list, check if any pending
		// var_source can be added to ordered list.
		left := []pendingVarSource{}
		for _, pendingVs := range pending {
			miss := false
			for _, dep := range pendingVs.deps {
				if added[dep] == nil {
					miss = true
					break
				}
			}
			if !miss {
				ordered = append(ordered, pendingVs.vs)
				added[pendingVs.vs.Name] = true
			} else {
				left = append(left, pendingVs)
			}
		}
		pending = left
	}

	if len(pending) > 0 {
		names := []string{}
		for _, vs := range pending {
			names = append(names, vs.vs.Name)
		}
		return nil, fmt.Errorf("could not resolve inter-dependent var sources: %s", strings.Join(names, ", "))
	}

	return ordered, nil
}

func DefaultTLSConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,

		// https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.CurveP384,
			tls.CurveP521,
		},

		// Security team recommends a very restricted set of cipher suites
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},

		PreferServerCipherSuites: true,
		NextProtos:               []string{"h2"},
	}
}

func DefaultSSHConfig() ssh.Config {
	return ssh.Config{
		// use the defaults prefered by go, see https://github.com/golang/crypto/blob/master/ssh/common.go
		Ciphers: nil,

		// CIS recommends a certain set of MAC algorithms to be used in SSH connections. This restricts the set from a more permissive set used by default by Go.
		// See https://infosec.mozilla.org/guidelines/openssh.html and https://www.cisecurity.org/cis-benchmarks/
		MACs: []string{
			"hmac-sha2-256-etm@openssh.com",
			"hmac-sha2-256",
		},

		//[KEX Recommendations for SSH IETF](https://tools.ietf.org/html/draft-ietf-curdle-ssh-kex-sha2-10#section-4)
		//[Mozilla Openssh Reference](https://infosec.mozilla.org/guidelines/openssh.html)
		KeyExchanges: []string{
			"ecdh-sha2-nistp256",
			"ecdh-sha2-nistp384",
			"ecdh-sha2-nistp521",
			"curve25519-sha256@libssh.org",
		},
	}
}
