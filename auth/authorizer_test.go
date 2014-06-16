package auth

import (
	"code.google.com/p/go.crypto/ssh"
	"testing"
)

var TestAuthorizerKeys = [][]byte{
	[]byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC/E9vvgNYuyqJtEA5KECqdVg80ncVY7jBWn00UzB9r+NmPfgNV4+IgfYLSyIEdIvIia+dHxijCLPE1rth8BAYeShS5LKHo0V304VuoyHUJ4F1rdYi6u44xxGIYj3poP9E35moyJbGyT7efs8D7s2t4SeL/xdtBBsbZQB7vVpfR+InmzyOwrlurepueaaJGnPRpIC3sSF+6Wp3XQ1wNfjQkx8B37diAWzJG4x/09HV82A/6RlW11GWC9ueeSK1MgOgzovE+ApURj8jEmWZwxQ3fFezTGxQonwTSfcwSdro/dUaxBAQbt4vylXxDzp/smP8UwU8cWTDHuOce+FlrIKYZ key1"),
	[]byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDEqAvKUyefV/WTSCLfX3V/HhyWHsX3ergMLDjOe6MzLH1thkUQ068O2DOnV03zL6ItwmWPNMIo9IdLMcYR5nTLqm2ALD8HPde3uzaOvHU64XoAbM39sTjNPbxMbjGpVRZVdAFYUiisBssQxpDb1ayOmGacVeaDXcFds97wrrxl9AQ6N4wBvsgXrTKbu0HfRVn5/fEg2hyoZQYtxRoDMyduHo94s9fb7STI+/nWLnH/6c8c0sTCI8rVrNujvqsZygj2P+6deA9bdGgIGoC6tTqJhgdZvkz3D2XGfndjPCaJGVtyqPeZX9oAtlQJXnn1Iz8vEwhQPiWYKLvuS//mK2Zf key2"),
	[]byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDjEDzUrQo+eiAqgxRBJzrY9EPpTq4CyVDm9NMs6MinZ06PQLbxshZxruxrsrZtfnps1cz0HYqFCkSLTk/2KYdpstYxb0af2k6xxmGtTNdCDQTD3MThW+T3Qv58wXo75EgFI3ZEgGgo/bcALpaMYN0skil4cSv/Vd1B36tKgF6QKf3vFSsiiVJ8pPrM1v32DJ/DvLvQb5VsPrDfTcMqsfNXyRjXNMZqGmK2uAlZBcbO6o4M+yrRtt9wRnTJvAdLIip+ZF2O1q28atlluUWRrneW39oy9p/M5w/Gw/BuCRoshZadBovziXdHnlCh00JbM/xS+m1GsDXhfWIYuSXqTSCd key3"),
}

func TestAuthorizer(t *testing.T) {
	authorizer := NewAuthorizer()

	publicKeys := make([]ssh.PublicKey, len(TestAuthorizerKeys))

	for i, rawKey := range TestAuthorizerKeys {
		publicKey, _, _, _, err := ssh.ParseAuthorizedKey(rawKey)
		if err != nil {
			t.Fatal(err)
		}

		// Save the key.
		publicKeys[i] = publicKey
	}

	for _, publicKey := range publicKeys {
		if authorizer.Authorize(publicKey) {
			t.Error("authorized a key that hasn't been added.")
		}

		// Add the key to the authorizer.
		authorizer.AddKey(publicKey)
	}

	for _, publicKey := range publicKeys {
		if !authorizer.Authorize(publicKey) {
			t.Error("failed to authorized a key that has been added.")
		}
	}
}