package main

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/koesie10/webauthn/protocol"
)

func main() {

	rpId := "https://example.org"
	rpOrigin := "https://example.org"

	assertionChallenge, err := base64.RawURLEncoding.DecodeString("rtnHiVQ7")
	if err != nil {
		fmt.Printf("Challenge Format Error: %v", err)
		return
	}

	b64Id := "1iBgFW1tSMSO_4mre7JJeQ"
	rawId, err := base64.RawURLEncoding.DecodeString(b64Id)
	if err != nil {
		fmt.Printf("ID Format Error: %v", err)
		return
	}

	attsClientData, err := base64.RawURLEncoding.DecodeString("eyJjaGFsbGVuZ2UiOiJydG5IaVZRNyIsIm9yaWdpbiI6Imh0dHBzOi8vZXhhbXBsZS5vcmciLCJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIn0")
	if err != nil {
		fmt.Printf("ClientData Format Error: %v", err)
		return
	}

	attsObj, err := base64.RawURLEncoding.DecodeString("o2hhdXRoRGF0YVkAlFDXqQXjBGuIY4NizDSjGhrlNHZspV46o5eVHv5lOwYrRAAAAAAAAAAAAAAAAAAAAAAAAAAAABDWIGAVbW1IxI7_iat7skl5pQECAyYgASFYIIWPShqhQ_g1b_ZhzAWWN5UWh8gHOwbNWtdui9M2xf-jIlggLzvDxIPB69pTsHL0u7YNApNdcMaQd3uDiz6Y-0ENYLtjZm10ZnBhY2tlZGdhdHRTdG10omNzaWdYRzBFAiEA0zMrn6VENt3UNZ7TbRR06SgJmMIaVm3EOYq6SbTcWRYCID6zT1AsTTKzYhfYSiqq-a3guNvRiaitDthkxRNAthAXY2FsZyY")
	if err != nil {
		fmt.Printf("Attestation Format Error: %v", err)
		return
	}

	assertionClientData, err := base64.RawURLEncoding.DecodeString("")
	if err != nil {
		fmt.Printf("ClientData Format Error: %v", err)
		return
	}

	authData, err := base64.RawURLEncoding.DecodeString("")
	if err != nil {
		fmt.Printf("AuthData Format Error: %v", err)
		return
	}
	signature, err := base64.RawURLEncoding.DecodeString("")
	if err != nil {
		fmt.Printf("Signature Format Error: %v", err)
		return
	}
	userHandle, err := base64.RawURLEncoding.DecodeString("bHlva2F0bw")
	if err != nil {
		fmt.Printf("UserHandle Format Error: %v", err)
		return
	}

	attsRes := protocol.AttestationResponse{
		PublicKeyCredential: protocol.PublicKeyCredential{
			ID:    b64Id,
			RawID: rawId,
			Type:  "public-key",
		},
		Response: protocol.AuthenticatorAttestationResponse{
			AuthenticatorResponse: protocol.AuthenticatorResponse{
				ClientDataJSON: attsClientData,
			},
			AttestationObject: attsObj,
		},
	}

	fmt.Println("Parse Attestation Response")
	atts, err := protocol.ParseAttestationResponse(attsRes)
	if err != nil {
		e := protocol.ToWebAuthnError(err)
		fmt.Printf("Error: %s, %s, %s", e.Name, e.Debug, e.Hint)
		return
	}

//	 This returns err, because this webauthn library doesn't support self-attestation
	validAtts, err := protocol.IsValidAttestation(atts, assertionChallenge, rpId, rpOrigin)
	if err != nil {
		e := protocol.ToWebAuthnError(err)
		fmt.Printf("Error: %s, %s, %s", e.Name, e.Debug, e.Hint)
		return
	}
	if !validAtts {
		fmt.Println("Invalid Attestation!")
		return
	}

	return

	pubKey := atts.Response.Attestation.AuthData.AttestedCredentialData.COSEKey

	cert := &x509.Certificate{
		PublicKey: pubKey,
	}

	assertionRes := protocol.AssertionResponse{
		PublicKeyCredential: protocol.PublicKeyCredential{
			ID:    b64Id,
			RawID: rawId,
			Type:  "public-key",
		},
		Response: protocol.AuthenticatorAssertionResponse{
			AuthenticatorResponse: protocol.AuthenticatorResponse{
				ClientDataJSON: assertionClientData,
			},
			AuthenticatorData: authData,
			Signature:         signature,
			UserHandle:        userHandle,
		},
	}
	assertion, err := protocol.ParseAssertionResponse(assertionRes)
	if err != nil {
		e := protocol.ToWebAuthnError(err)
		fmt.Printf("Error: %s, %s, %s", e.Name, e.Debug, e.Hint)
		return
	}

	valid, err := protocol.IsValidAssertion(assertion, assertionChallenge, rpId, rpOrigin, cert)
	if err != nil {
		e := protocol.ToWebAuthnError(err)
		fmt.Printf("Error: %s, %s, %s", e.Name, e.Debug, e.Hint)
		return
	}

	if !valid {
		fmt.Println("Invalid Assertion!")
		return
	}

	fmt.Println("Valid Assertion!!!")

}
