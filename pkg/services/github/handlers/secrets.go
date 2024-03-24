package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/alpha-omega-corp/services/types"
	"github.com/google/go-github/v56/github"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/nacl/box"
)

// Encrypt
const (
	keySize   = 32
	nonceSize = 24
)

var generateKey = box.GenerateKey

type SecretsHandler interface {
	GetAll(ctx context.Context) ([]*github.Secret, error)
	Create(ctx context.Context, name string, content []byte) error
	Delete(ctx context.Context, name string) error
	GetKey(ctx context.Context) (*github.PublicKey, error)
	Encrypt(recipientPublicKey string, content string) (string, error)
}

type secretsHandler struct {
	SecretsHandler

	client *github.Client
	org    string
}

func NewSecretsHandler(cli *github.Client, c types.Config) SecretsHandler {
	return &secretsHandler{
		client: cli,
		org:    c.Viper.GetString("name"),
	}
}

func (h *secretsHandler) GetAll(ctx context.Context) ([]*github.Secret, error) {
	data, _, err := h.client.Actions.ListOrgSecrets(ctx, h.org, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	return data.Secrets, nil
}

func (h *secretsHandler) Create(ctx context.Context, name string, content []byte) error {
	key, err := h.GetKey(ctx)
	if err != nil {
		return err
	}

	encodedString, err := h.Encrypt(key.GetKey(), string(content))
	if err != nil {
		return err
	}

	_, err = h.client.Actions.CreateOrUpdateOrgSecret(ctx, h.org, &github.EncryptedSecret{
		Name:           name,
		KeyID:          key.GetKeyID(),
		EncryptedValue: encodedString,
		Visibility:     "all",
	})
	if err != nil {
		return err
	}

	return nil
}

func (h *secretsHandler) Delete(ctx context.Context, name string) error {
	_, err := h.client.Actions.DeleteOrgSecret(ctx, h.org, name)
	if err != nil {
		return err
	}

	return nil
}

func (h *secretsHandler) GetKey(ctx context.Context) (*github.PublicKey, error) {
	data, _, err := h.client.Actions.GetOrgPublicKey(ctx, h.org)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *secretsHandler) Encrypt(recipientPublicKey string, content string) (string, error) {
	// decode the provided public key from base64
	recipientKey := new([keySize]byte)
	b, err := base64.StdEncoding.DecodeString(recipientPublicKey)
	if err != nil {
		return "", err
	} else if size := len(b); size != keySize {
		return "", fmt.Errorf("recipient public key has invalid length (%d bytes)", size)
	}

	copy(recipientKey[:], b)

	// create an ephemeral key pair
	pubKey, pKey, err := generateKey(rand.Reader)
	if err != nil {
		return "", err
	}

	// create the nonce by hashing together the two public keys
	nonce := new([nonceSize]byte)
	nonceHash, err := blake2b.New(nonceSize, nil)
	if err != nil {
		return "", err
	}

	if _, err := nonceHash.Write(pubKey[:]); err != nil {
		return "", err
	}

	if _, err := nonceHash.Write(recipientKey[:]); err != nil {
		return "", err
	}

	copy(nonce[:], nonceHash.Sum(nil))

	// begin the output with the ephemeral public key and append the encrypted content
	out := box.Seal(pubKey[:], []byte(content), nonce, recipientKey, pKey)

	// base64-encode the final output
	return base64.StdEncoding.EncodeToString(out), nil
}
