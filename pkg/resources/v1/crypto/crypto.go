package crypto

import "crypto"

// not a real salt, just testing
const salt = "AAAAB3NzaC1yc2EAAAADAQABAAABgQCxtDNo1HK28mqt22522Zm+5oaZfkLzTmoi81ExWbYWStRs/sSZwIlC8VlgaZQrymmJ78J0GoFoHogNbr+zLlp2G+mDo/kHFMJ8s+Uxy3zh9sl1jRMG2qkT6yjLJCNVWXvIMnywqItA0p3ZXtoAs+IW7aseMRK/IBrb+9cs/78VyMYR9zmCaPSOQIm6MsfihM2LxKERXkyphla5pZZQN4gp+Tp8v/sIZU3XAFyejhzEESH8Nb5joLVvamOx5bnG+O2kL/y/qEUIib8rBv0ZWS0qLQKYSBHQi6bZ7WP/RitWN5JIOMrChdlFTiOrQiK/kzWvBAVNhvJGJ9U/snCFtGOthdktNMir6f2XcaBrkpc6vhoXZdNaWh4qzCxpQM+YAKeJviLli0Ske/O/7lcVi9ItsQznJMpsUkJfdt8V/McyS8cvjn0R0S0sKTvV+ckEmMPYvpWyTJ90L5wmNegRzeJ1Od7XeDjwGw4xbUiRVwxhgsxlgA5rwyf5XOqPcHNmnTU="

func HashText(text string) []byte {
	hash := crypto.SHA512_256.New()
	hash.Write([]byte(text))
	return hash.Sum([]byte(salt))
}
