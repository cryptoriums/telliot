// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package private_file

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/cryptoriums/telliot/pkg/prompt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

func createHash(key string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(key)); err != nil {
		return "", errors.Wrap(err, "hasher.Write")
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	h, err := createHash(passphrase)
	if err != nil {
		return nil, errors.Wrap(err, "createHash")
	}
	block, _ := aes.NewCipher([]byte(h))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "NewGCM")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "read full")
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func Decrypt(data []byte, passphrase string) ([]byte, error) {
	h, err := createHash(passphrase)

	if err != nil {
		return nil, errors.Wrap(err, "createHash")
	}
	block, err := aes.NewCipher([]byte(h))
	if err != nil {
		return nil, errors.Wrap(err, "NewCipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "NewGCM")
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.Wrap(err, "gcm.Open")
	}
	return plaintext, nil
}

func EncryptFile(inFile string, outFile string, passphrase string) error {
	out, err := os.Create(outFile)
	if err != nil {
		return errors.Wrap(err, "creating output file")
	}
	defer out.Close()

	in, err := ioutil.ReadFile(inFile)
	if err != nil {
		return errors.Wrap(err, "reading input file")
	}

	bb, err := encrypt(in, passphrase)
	if err != nil {
		return errors.Wrap(err, "encrypt")
	}

	_, err = out.Write(bb)
	if err != nil {
		return errors.Wrap(err, "writing to output file")
	}
	return nil
}

func DecryptWithWebPassword(ctx context.Context, logger log.Logger, input []byte, host string, port uint) []byte {
	fileBytes := make(chan [][]byte)
	srv := &http.Server{Addr: host + ":" + strconv.Itoa(int(port))}
	defer func() {
		if err := srv.Shutdown(ctx); err != nil {
			level.Error(logger).Log("msg", "shutting down the password prompt web server", "err", err)
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			level.Error(logger).Log("msg", "starting the password web server", "err", err)
			fileBytes <- [][]byte{nil}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		postResult := ""
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("parsing form data:%v", err), http.StatusInternalServerError)
				return
			}
			pass := r.PostForm.Get("decryptPass")
			output, err := Decrypt(input, pass)
			if err != nil {
				postResult = "Decrypt error try again:" + err.Error()
			} else {
				postResult = `File unlocked!`
				fileBytes <- [][]byte{output}
			}
		}
		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang='en'>
		<head>
			<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1'>
			<title>Decrypt password</title>
			<style>
				body {
					font-family: arial;
				}
				label {
					min-width: 9em;
					float: left;
				}
			</style>
		</head>
		<body>
		`+postResult+`
		<form id="data" method="post">
			<label for="decryptPass">Decrypt pass:</label><input type="password" name="decryptPass" id="decryptPass"/>
			<input type="submit" value="GO">
		</form>
		</body>
		</html>
		`)

	})

	b := <-fileBytes
	return b[0]
}

func DecryptWithPasswordLoop(input []byte) []byte {
	for {
		pass, err := prompt.Prompt("Enter Password: ", true)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting password from terminal:", err)
			continue
		}
		output, err := Decrypt(input, pass)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("Decrypt error try again:", err)
			continue
		}
		return output
	}
}

func EncryptWithPasswordLoop(inFile string, outFile string) error {
	for {
		pass, err := prompt.Prompt("Enter Password: ", true)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting password from terminal:", err)
			continue
		}
		if pass == "" {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("password can't be empty")
			continue
		}

		err = EncryptFile(inFile, outFile, pass)
		if err != nil {
			return errors.Wrap(err, "encrypt file")
		}
		return nil
	}
}
