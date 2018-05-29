/*
   enum.go
      Enum values for the Nicehash API
*/

package nicehash

var ALGORITHMS = map[int64]string{
	0:  "scrypt",
	1:  "sha256",
	2:  "scryptnf",
	3:  "x11",
	4:  "x13",
	5:  "keccak",
	6:  "x15",
	7:  "nist5",
	8:  "neoscrypt",
	9:  "lyra2re",
	10: "whirlpoolx",
	11: "qubit",
	12: "quark",
	13: "axiom",
	14: "lyra2rev2",
	15: "scryptjanenf16",
	16: "blake256r8",
	17: "blake256r14",
	18: "blake256r8vnl",
	19: "hodl",
	20: "daggerhashimoto",
	21: "decred",
	22: "cryptonight",
	23: "lbry",
	24: "equihash",
	25: "pascal",
	26: "x11gost",
	27: "sia",
}

var ORDER_TYPES = map[int64]string{
	0: "standard",
	1: "fixed",
}
